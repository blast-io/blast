import { promises as fs } from 'fs'

import { task, types } from 'hardhat/config'
import '@nomiclabs/hardhat-ethers'
import 'hardhat-deploy'
import { predeploys } from '@eth-optimism/core-utils'
import { providers, utils, ethers } from 'ethers'
import Artifact__Shares from '@eth-optimism/contracts-bedrock/forge-artifacts/Shares.sol/Shares.json'

type NestedJson = {name: string, dir?: NestedJson[], provider?: 'l1' | 'l2', balance?: ethers.BigNumber, hide?: boolean, hideIfZero?: boolean, isAddress?: boolean, negate?: boolean};
type NestedJsonWithBalance = {name: string, dir?: NestedJsonWithBalance[], provider?: 'l1' | 'l2', balance: ethers.BigNumber, hide?: boolean, hideIfZero?: boolean, isAddress?: boolean, negate?: boolean};

const { formatEther } = utils

task('audit', 'audits balances across L1/L2 networks')
  .addParam(
    'l2providerurl',
    'L2 provider URL.',
    'http://localhost:9545',
    types.string
  )
  .addParam(
    'l1providerurl',
    'L1 provider URL.',
    'http://localhost:8545',
    types.string
  )
  .addParam(
    'faucetkey',
    'Faucet Private Key',
    '',
    types.string
  )
   .addParam(
     'stats',
     'stats',
     '',
     types.string
   )
  .addParam(
    'book',
    'contract address book',
    '',
    types.string
  )
  .setAction(async (args, hre) => {
    const stats: {accounts: Map<string, string>, contracts: Map<string, string>} = JSON.parse(args.stats);
    if (Object.entries(stats).length == 0) {
      stats.contracts = new Map<string, string>();
      stats.accounts = new Map<string, string>();
    }
    const initialBook: Record<string, string> = JSON.parse(args.book);

    const l1Provider = new providers.StaticJsonRpcProvider(args.l1providerurl)
    const l2Provider = new providers.StaticJsonRpcProvider(args.l2providerurl)
    // Use the first configured signer for simplicity
    const l1Faucet = new hre.ethers.Wallet(args.faucetkey, l1Provider);
    const l2Faucet = new hre.ethers.Wallet(args.faucetkey, l2Provider);

    const book = {...predeploys, ...reverseObject(initialBook)}
    const l2book = predeploys
    const l1book = reverseObject(initialBook)
    const gasContracts = ['L1FeeVault', 'SequencerFeeVault', 'BaseFeeVault', 'Gas']
    const wethContracts = ['WETH9', 'WETHRebasing']
    const sharesContract = new ethers.Contract(l2book['Shares'], Artifact__Shares.abi, l2Provider)
    const price = await sharesContract.price();
    const count = await sharesContract.count();
    const pending = await sharesContract.pending();
    const predeployAddr = Object.values(predeploys)

    const l2: NestedJson = 
        {name: 'Blast eth', dir: [
          {name: 'Gas Costs', dir: [
                ...gasContracts.map((item) => {return {name: item, provider: 'l2' as const}})
            ]
          },
          {name: 'Wrapped ETH', dir: [
                ...wethContracts.map((item) => {return {name: item, provider: 'l2' as const}})
            ], hideIfZero: true
          },
          {name: 'Predeploys', dir: [
                ...Object.keys(l2book).filter((key) => !gasContracts.includes(key) && !wethContracts.includes(key)).map((item) => {return {name: item, provider: 'l2' as const, hideIfZero: true}})
            ]
          },
          {name: 'EOAs', dir: [
                ...Object.entries(stats.accounts).filter((item) => !predeployAddr.includes(item[0])).map((item) => {return {name: item[0], provider: 'l2' as const, balance: ethers.BigNumber.from(item[1])}})
            ]
          },
          {name: 'Contracts', dir: [
                ...Object.entries(stats.contracts).filter((item) => !predeployAddr.includes(item[0])).map((item) => {return {name: item[0], provider: 'l2' as const, balance: ethers.BigNumber.from(item[1])}})
            ]
          },{
            name: 'Burned', dir: [{name: 'L2ToL1MessagePasser', provider: 'l2' as const}], negate: true
          }
        ]}
    const l1BlastContracts: NestedJson = {name: 'Blast Contracts', dir: [
        ...Object.keys(l1book)
        .filter((key) => !key.includes("Proxy")) // account for proxy contracts together with implementation contracts
        .map((item) => {return {name: item, provider: 'l1' as const, hideIfZero: true}}),
    ]}

    const l1Misc: NestedJson = {name: 'L1 Misc', dir: Object.entries(stats.accounts).filter((item) => !predeployAddr.includes(item[0]) && l1Faucet.address.toLowerCase() != item[0].toLowerCase()).map((item) => {return {name: item[0], provider: 'l1' as const, isAddress: true}})};

    try {
        await Promise.all([populate(l1BlastContracts, {book, l1Provider, l2Provider}), populate(l2, {book, l1Provider, l2Provider}), populate(l1Misc, {book, l1Provider, l2Provider})]);
    } catch(err) {
        console.log(err)
        throw(err)
    }

    formatEthereumDistribution(transformToJsonString(l2 as any))
    formatEthereumDistribution(transformToJsonString(l1BlastContracts as any))
    formatEthereumDistribution(transformToJsonString(l1Misc as any))

    if (l1BlastContracts.balance === undefined || l2.balance === undefined) {
        throw new Error("balances are undefined")
    }
    if (l1BlastContracts.balance.eq(l2.balance)) {
        console.log("L1 balance == L2 balance")
    } else {
        console.log("Discrepancy detected in L1 balance, L2 balance")
        console.log(`L1 balance: ${l1BlastContracts.balance.toNumber()}`)
        console.log(`L2 balance: ${l2.balance.toNumber()}`)
    }
  })

function reverseObject(obj: Record<string, string>) {
    const reversed = {};
    for (let key in obj) {
        if (obj.hasOwnProperty(key)) {
            reversed[obj[key]] = key;
        }
    }
    return reversed;
}

async function populate(input: NestedJson, args: {book: Record<string, string>, l1Provider: ethers.providers.StaticJsonRpcProvider, l2Provider: ethers.providers.StaticJsonRpcProvider}): Promise<{bal: ethers.BigNumber, isNegative: boolean}>{
  let balance = ethers.BigNumber.from(0);
  const isNeg = input.negate ?? false
  if (input.dir !== undefined) {
    for (const dir of input.dir) {
      const {bal, isNegative} = await populate(dir, args)
      if (isNegative) {
        balance = balance.sub(bal);
      } else {
        balance = balance.add(bal)
      }
    }
  } else {
    if (input.balance) {
      balance = input.balance;
    } else {
      let address: string = ""
      if (input.isAddress === true) {
        address = input.name
      } else  {
        address = args.book[input.name];
      }
      if (address.length === 0) {
        console.log(args.book)
        throw new Error(`${input.name} not defined`);
      }
      const providerType = input.provider ?? 'l1';
      const provider = providerType === 'l1' ? args.l1Provider : args.l2Provider;
      balance = await provider.getBalance(address);

      // check if proxy exists
      const proxyContractName = input.name + 'Proxy'
      const hasProxy = typeof args.book[proxyContractName] === 'string'
      if (hasProxy) {
        const proxyAddress = args.book[proxyContractName]
        const proxyBalance = await provider.getBalance(proxyAddress)
        balance = balance.add(proxyBalance);
      }
    }
  }
  input.balance = balance;
  return {bal: balance, isNegative: isNeg};
}

function roundETH(eth: string , shouldRound: boolean, isNegative: boolean) {
    if (!shouldRound) {
      if (isNegative){
        return '-' + eth
      }
      return eth;
    }

    const rounded = Math.round(parseFloat(eth) * 100) / 100; // Round to 2 decimal places
    if (rounded === 0 && parseFloat(eth) != 0) {
        return '~0'
    }
    if (isNegative) {
      return '-' + rounded
    }
    return rounded;
}


function transformToJsonString(obj: NestedJsonWithBalance, depth: number = 0): string {
    const indentation = '   '.repeat(depth);
    let result = `${indentation}- ${roundETH(formatEther(obj.balance), true, obj.negate ?? false)} eth in ${obj.name}\n`;

    // Recursively handle directories
    if (obj.dir && obj.dir.length > 0) {
        for (const subObj of obj.dir) {
            if (subObj.hide === true) {
                continue
            }
            if (subObj.hideIfZero === true && subObj.balance.eq(ethers.BigNumber.from(0))) {
                continue
            }
            // Adding indentation for nested entries
            result += `${transformToJsonString(subObj, depth+1)}`;
        }
    }

    return result;
}
function formatEthereumDistribution(data: string) {
    const colors = {
        cyan: '\x1b[36m',
        green: '\x1b[32m',
        yellow: '\x1b[33m',
        magenta: '\x1b[35m', // For numbers
        reset: '\x1b[0m'
    };

    const formatLine = (line: string, level: number) => {
        let color;
        if (level === 0) color = colors.cyan;
        else if (level === 1) color = colors.green;
        else color = colors.yellow;

        // Adding | for structure
        const indent = ' '.repeat(level);
        const formattedLine = line.replace(/-/g, '').replace(/(\d+\.?\d*\s*eth)/gi, `${colors.magenta}$1${color}`);

        return `${indent}${color}${formattedLine}${colors.reset}`;
    };

    const lines = data.split('\n');
    for (const line of lines) {
      if (line.length === 0) {
       continue 
      }
      const dashIndex = line.indexOf('-');
      const level = dashIndex / 3
      const formattedLine = formatLine(line, level)
      if (formattedLine.trim().length > 0 ) {
        console.log(formattedLine);
      }
    }
}
