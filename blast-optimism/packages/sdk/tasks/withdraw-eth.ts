
import { promises as fs } from 'fs'

import { task, types } from 'hardhat/config'
import '@nomiclabs/hardhat-ethers'
import 'hardhat-deploy'
import { Deployment } from 'hardhat-deploy/types'
import { predeploys } from '@eth-optimism/core-utils'
import { providers, utils, ethers, BigNumber, Wallet } from 'ethers'
import Artifact__L2ToL1MessagePasser from '@eth-optimism/contracts-bedrock/forge-artifacts/L2ToL1MessagePasser.sol/L2ToL1MessagePasser.json'
import Artifact__L2CrossDomainMessenger from '@eth-optimism/contracts-bedrock/forge-artifacts/L2CrossDomainMessenger.sol/L2CrossDomainMessenger.json'
import Artifact__L2StandardBridge from '@eth-optimism/contracts-bedrock/forge-artifacts/L2StandardBridge.sol/L2StandardBridge.json'
import Artifact__OptimismPortal from '@eth-optimism/contracts-bedrock/forge-artifacts/OptimismPortal.sol/OptimismPortal.json'
import Artifact__L1CrossDomainMessenger from '@eth-optimism/contracts-bedrock/forge-artifacts/L1CrossDomainMessenger.sol/L1CrossDomainMessenger.json'
import Artifact__L1StandardBridge from '@eth-optimism/contracts-bedrock/forge-artifacts/L1StandardBridge.sol/L1StandardBridge.json'
import Artifact__L2OutputOracle from '@eth-optimism/contracts-bedrock/forge-artifacts/L2OutputOracle.sol/L2OutputOracle.json'
import Artifact__ETHYieldManager from '@eth-optimism/contracts-bedrock/forge-artifacts/ETHYieldManager.sol/ETHYieldManager.json'
import Artifact__L2BlastBridge from '@eth-optimism/contracts-bedrock/forge-artifacts/L2BlastBridge.sol/L2BlastBridge.json'


import {
  CrossChainMessenger,
  MessageStatus,
  CONTRACT_ADDRESSES,
  OEContractsLike,
  DEFAULT_L2_CONTRACT_ADDRESSES,
} from '../src'
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers'
import { get, request } from 'http'

const { formatEther } = utils
task('withdraw-eth-2', 'withdraw ether from L2.')
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
    'adminkey',
    'Yield provider admin key.',
    'http://localhost:8545',
    types.string
  )
  .addParam(
    'book',
    'contract address book',
    '',
    types.string
  )
  .addParam(
    'bridge',
    'bridge',
    'blast',
    types.string
  )
  .addOptionalParam('userkey', 'user calling withdraw', '', types.string)
  .addOptionalParam('amount', 'Amount to withdraw', '', types.string)
  .addOptionalParam(
    'l1ContractsJsonPath',
    'Path to a JSON with L1 contract addresses in it',
    'TODO_FILL_THIS_IN',
    types.string
  )
  .addOptionalParam('to', 'recipient of ETH', '', types.string)
  .setAction(async (args, hre) => {
    const book: Record<string, string> = reverseObject(JSON.parse(args.book));
    const l2Provider = new providers.StaticJsonRpcProvider(args.l2providerurl)
    const l1Provider = new providers.StaticJsonRpcProvider(args.l1providerurl)
    const adminSigner = new hre.ethers.Wallet(
      args.adminkey,
     l1Provider
    )
    const l1Signer = new hre.ethers.Wallet(
      args.userkey,
     l1Provider
    )
    const l2Signer = new hre.ethers.Wallet(
      args.userkey,
      l2Provider
    )
    // Use the first configured signer for simplicity
    const address = await l1Signer.getAddress()

    // Ensure that the signer has a balance before trying to
    // do anything
    const balance = await l1Signer.getBalance()
    console.log(`L1 Signer balance: ${formatEther(balance.toString())}`)

    const withdrawAmount = utils.parseEther(args.amount)

    const l2ChainId = await l2Signer.getChainId()
    let contractAddrs = CONTRACT_ADDRESSES[l2ChainId]
    if (book) {
      // const data = await fs.readFile(args.l1ContractsJsonPath)
      const json  = book as any
      // const json = JSON.parse(data.toString())
      contractAddrs = {
        l1: {
          AddressManager: json.AddressManager,
          L1CrossDomainMessenger: json.L1CrossDomainMessengerProxy,
          L1StandardBridge: json.L1StandardBridgeProxy,
          StateCommitmentChain: ethers.constants.AddressZero,
          CanonicalTransactionChain: ethers.constants.AddressZero,
          BondManager: ethers.constants.AddressZero,
          OptimismPortal: json.OptimismPortalProxy,
          L2OutputOracle: json.L2OutputOracleProxy,
          ETHYieldManager: json.ETHYieldManagerProxy
        },
        l2: DEFAULT_L2_CONTRACT_ADDRESSES,
      } as OEContractsLike
      console.log(contractAddrs)
    } else if (!contractAddrs) {
      // If the contract addresses have not been hardcoded,
      // attempt to read them from deployment artifacts
      let Deployment__AddressManager: Deployment
      try {
        Deployment__AddressManager = await hre.deployments.get('AddressManager')
      } catch (e) {
        Deployment__AddressManager = await hre.deployments.get(
          'Lib_AddressManager'
        )
      }
      let Deployment__L1CrossDomainMessenger: Deployment
      try {
        Deployment__L1CrossDomainMessenger = await hre.deployments.get(
          'L1CrossDomainMessengerProxy'
        )
      } catch (e) {
        Deployment__L1CrossDomainMessenger = await hre.deployments.get(
          'Proxy__OVM_L1CrossDomainMessenger'
        )
      }
      let Deployment__L1StandardBridge: Deployment
      try {
        Deployment__L1StandardBridge = await hre.deployments.get(
          'L1StandardBridgeProxy'
        )
      } catch (e) {
        Deployment__L1StandardBridge = await hre.deployments.get(
          'Proxy__OVM_L1StandardBridge'
        )
      }

      const Deployment__OptimismPortal = await hre.deployments.get(
        'OptimismPortalProxy'
      )
      const Deployment__L2OutputOracle = await hre.deployments.get(
        'L2OutputOracleProxy'
      )
      contractAddrs = {
        l1: {
          AddressManager: Deployment__AddressManager.address,
          L1CrossDomainMessenger: Deployment__L1CrossDomainMessenger.address,
          L1StandardBridge: Deployment__L1StandardBridge.address,
          StateCommitmentChain: ethers.constants.AddressZero,
          CanonicalTransactionChain: ethers.constants.AddressZero,
          BondManager: ethers.constants.AddressZero,
          OptimismPortal: Deployment__OptimismPortal.address,
          L2OutputOracle: Deployment__L2OutputOracle.address,
        },
        l2: DEFAULT_L2_CONTRACT_ADDRESSES,
      }
    }

    const OptimismPortal = new hre.ethers.Contract(
      contractAddrs.l1.OptimismPortal,
      Artifact__OptimismPortal.abi,
      l1Signer
    )

    const L1CrossDomainMessenger = new hre.ethers.Contract(
      contractAddrs.l1.L1CrossDomainMessenger,
      Artifact__L1CrossDomainMessenger.abi,
     l1Signer
    )

    const L1StandardBridge = new hre.ethers.Contract(
      contractAddrs.l1.L1StandardBridge,
      Artifact__L1StandardBridge.abi,
     l1Signer
    )

    const L2OutputOracle = new hre.ethers.Contract(
      contractAddrs.l1.L2OutputOracle,
      Artifact__L2OutputOracle.abi,
     l1Signer
    )

    const ETHYieldManager = new hre.ethers.Contract(
      contractAddrs.l1.ETHYieldManager,
      Artifact__ETHYieldManager.abi,
     adminSigner
    )

    const L2ToL1MessagePasser = new hre.ethers.Contract(
      predeploys.L2ToL1MessagePasser,
      Artifact__L2ToL1MessagePasser.abi
    )

    const L2CrossDomainMessenger = new hre.ethers.Contract(
      predeploys.L2CrossDomainMessenger,
      Artifact__L2CrossDomainMessenger.abi
    )

    const L2StandardBridge = new hre.ethers.Contract(
      predeploys.L2StandardBridge,
      Artifact__L2StandardBridge.abi
    )

    const L2BlastBridge = new hre.ethers.Contract(
      predeploys.L2BlastBridge,
      Artifact__L2BlastBridge.abi
    )

    const messenger = new CrossChainMessenger({
      l1SignerOrProvider: l1Signer,
      l2SignerOrProvider: l2Signer,
      l1ChainId: await l1Signer.getChainId(),
      l2ChainId,
      bedrock: true,
      contracts: contractAddrs,
    })

    const opBalance = await l1Signer!.provider!.getBalance(
      ETHYieldManager.address
    )


    let ethWithdrawReceipt: providers.TransactionReceipt;
    console.log('Withdrawing ETH')
    if (args.bridge === 'blast') {
      const tx = await l2Signer.sendTransaction({
        to: L2BlastBridge.address,
        value: withdrawAmount 
      });
      console.log(`Transaction hash: ${tx.hash}`)
      ethWithdrawReceipt = await tx.wait();

    } else if (args.bridge === 'standard'){
      const tx = await l2Signer.sendTransaction({
        to: L2StandardBridge.address,
        value: withdrawAmount 
      });
      console.log(`Transaction hash: ${tx.hash}`)
      ethWithdrawReceipt = await tx.wait();
    } else {
     const ethWithdraw = await messenger.withdrawETH(withdrawAmount)
     console.log(`Transaction hash: ${ethWithdraw.hash}`)
     ethWithdrawReceipt = await ethWithdraw.wait()
    }



     console.log(
       `ETH withdrawn on L2 - included in block ${ethWithdrawReceipt.blockNumber}`
     )

       type WithdrawalTransaction =  {nonce: BigNumber, sender: string, target: string, value: BigNumber, gasLimit: BigNumber, data: string }
       let wt: null | WithdrawalTransaction = null
     {
       // check the logs

       for (const log of ethWithdrawReceipt.logs) {
         switch (log.address) {
           case L2ToL1MessagePasser.address: {
             const parsed = L2ToL1MessagePasser.interface.parseLog(log)
             // console.log(parsed.name)
             // console.log(parsed.args)
             // console.log()
             if (parsed.name == 'MessagePassed') {
               wt = parsed.args as any
             }
             break
           }
           case L2StandardBridge.address: {
             const parsed = L2StandardBridge.interface.parseLog(log)
             // console.log(parsed.name)
             // console.log(parsed.args)
             // console.log()
             break
           }
           case L2CrossDomainMessenger.address: {
             const parsed = L2CrossDomainMessenger.interface.parseLog(log)
             // console.log(parsed.name)
             // console.log(parsed.args)
             // console.log()
             break
           }
           default: {
             console.log(`Unknown log from ${log.address} - ${log.topics[0]}`)
           }
         }
       }
     }
     if (wt === null) {
       console.log('couldnt find log')
       throw new Error()
     }
     console.log("withdrawal transaction parameters", JSON.stringify(wt))


     console.log('Waiting to be able to prove withdrawal')
     let latestNumber = 0

     const proveInterval = setInterval(async () => {
       const currentStatus = await messenger.getMessageStatus(ethWithdrawReceipt)
       const latest = await L2OutputOracle.latestBlockNumber()
       if (latestNumber != latest.toNumber()) {
       console.log(`Message status: ${MessageStatus[currentStatus]}`)
        console.log(
          `Latest L2OutputOracle commitment number: ${latest.toString()}`
        )
        latestNumber = latest.toNumber()
       }

       // const tip = await l1Signer.provider!.getBlockNumber()
       // console.log(`L1 chain tip: ${tip.toString()}`)
     }, 3000)

     try {
       await messenger.waitForMessageStatus(
         ethWithdrawReceipt,
         MessageStatus.READY_TO_PROVE
       )
     } finally {
       clearInterval(proveInterval)
     }

     console.log('Proving eth withdrawal...')
     const ethProve = await messenger.proveMessage(ethWithdrawReceipt)
     console.log(`Transaction hash: ${ethProve.hash}`)
     const ethProveReceipt = await ethProve.wait()
     // console.log({ethProve})
     // console.log({ethProveReceipt})

     if (ethProveReceipt.status !== 1) {
       throw new Error('Prove withdrawal transaction reverted')
     }
     console.log('Successfully proved withdrawal')

     console.log('Waiting to be able to finalize withdrawal')

     const finalizeInterval = setInterval(async () => {
       const currentStatus = await messenger.getMessageStatus(ethWithdrawReceipt)
       // console.log(`Message status: ${MessageStatus[currentStatus]}`)
     }, 3000)

     try {
       await messenger.waitForMessageStatus(
         ethWithdrawReceipt,
         MessageStatus.READY_FOR_RELAY
       )
     } finally {
       clearInterval(finalizeInterval)
     }

     let requestId
     for (const log of ethProveReceipt.logs) {
      console.log(log)
       switch (log.address.toLowerCase()) {
         case OptimismPortal.address.toLowerCase(): {
           const parsed = OptimismPortal.interface.parseLog(log)
           console.log(parsed)
           if (parsed.name === 'WithdrawalProven') {
             requestId = parsed.args.requestId
           }
         }
       }
     }

     console.log('eth yield manager finalizing requests')
     console.log(requestId)
     const gasLimit = 600000 // Set your desired gas limit here
     const response = await ETHYieldManager.finalize(BigNumber.from(requestId), {gasLimit})
     const finalizeResonseLogs = await response.wait()
     console.log('get hint from eth yield manager')
    const hint = await getHintRecursively(ETHYieldManager, requestId, adminSigner, 1)
    console.log({hint, requestId})

    console.log('Fulfilling withdrawal request...')
     const ethFulfill = await OptimismPortal.connect(adminSigner).finalizeWithdrawalTransaction(
       hint, 
       {nonce: wt.nonce, sender: wt.sender, target: wt.target, value: wt.value, gasLimit: wt.gasLimit, data: wt.data}
    )
     console.log(`Transaction hash: ${ethFulfill.hash}`)
     const ethFulfillReceipt = await ethFulfill.wait()
     if (ethFulfillReceipt.status !== 1) {
       throw new Error('Fulfill withdrawal reverted')
     }
    console.log('Withdraw success')
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
async function getHint(yieldManager: ethers.Contract, requestId: number, signer: Wallet) {
    const lastCheckpointId: BigNumber = await yieldManager.connect(signer).getLastCheckpointId()
    console.log({lastCheckpointId})
    const hint: BigNumber = await yieldManager.connect(signer).findCheckpointHint(requestId, 1, lastCheckpointId)
    if (hint.eq(BigNumber.from(0))) {
      throw new Error("cound not find hint")
    }
    return hint
  }

async function getHintRecursively(yieldManager: ethers.Contract, requestId: number, signer: Wallet, lastCheckpointId: number) {
  try {
    const hint: BigNumber = await yieldManager.connect(signer).findCheckpointHint(requestId, 1, lastCheckpointId)
    if (hint.eq(BigNumber.from(0))) {
      return await getHintRecursively(yieldManager, requestId, signer, lastCheckpointId+1)
    }
    return hint
  } catch (err) {
    throw new Error(err)
  }
}
