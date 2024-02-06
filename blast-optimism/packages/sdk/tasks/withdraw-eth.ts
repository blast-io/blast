
import { promises as fs } from 'fs'

import { task, types } from 'hardhat/config'
import '@nomiclabs/hardhat-ethers'
import 'hardhat-deploy'
import { Deployment } from 'hardhat-deploy/types'
import { predeploys } from '@eth-optimism/core-utils'
import { providers, utils, ethers } from 'ethers'
import Artifact__L2ToL1MessagePasser from '@eth-optimism/contracts-bedrock/forge-artifacts/L2ToL1MessagePasser.sol/L2ToL1MessagePasser.json'
import Artifact__L2CrossDomainMessenger from '@eth-optimism/contracts-bedrock/forge-artifacts/L2CrossDomainMessenger.sol/L2CrossDomainMessenger.json'
import Artifact__L2StandardBridge from '@eth-optimism/contracts-bedrock/forge-artifacts/L2StandardBridge.sol/L2StandardBridge.json'
import Artifact__OptimismPortal from '@eth-optimism/contracts-bedrock/forge-artifacts/OptimismPortal.sol/OptimismPortal.json'
import Artifact__L1CrossDomainMessenger from '@eth-optimism/contracts-bedrock/forge-artifacts/L1CrossDomainMessenger.sol/L1CrossDomainMessenger.json'
import Artifact__L1StandardBridge from '@eth-optimism/contracts-bedrock/forge-artifacts/L1StandardBridge.sol/L1StandardBridge.json'
import Artifact__L2OutputOracle from '@eth-optimism/contracts-bedrock/forge-artifacts/L2OutputOracle.sol/L2OutputOracle.json'

import {
  CrossChainMessenger,
  MessageStatus,
  CONTRACT_ADDRESSES,
  OEContractsLike,
  DEFAULT_L2_CONTRACT_ADDRESSES,
} from '../src'
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers'

const { formatEther } = utils
task('withdraw-eth-2', 'withdraw ether from L2.')
  .addParam(
    'l2ProviderUrl',
    'L2 provider URL.',
    'http://localhost:9545',
    types.string
  )
  .addParam(
    'l1ProviderUrl',
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
  .addOptionalParam('userkey', 'user calling withdraw', '', types.string)
  .addOptionalParam('amount', 'Amount to withdraw', '', types.string)
  .addOptionalParam(
    'l1ContractsJsonPath',
    'Path to a JSON with L1 contract addresses in it',
    '',
    types.string
  )
  .setAction(async (args, hre) => {
    const l2Provider = new providers.StaticJsonRpcProvider(args.l2ProviderUrl)
    const l1Provider = new providers.StaticJsonRpcProvider(args.l1ProviderUrl)
    const l1Signer = new hre.ethers.Wallet(
      args.userkey,
     l1Provider
    )
    const l2Signer = new hre.ethers.Wallet(
      args.userkey,
      l2Provider
    )
    // Use the first configured signer for simplicity
    console.log('network', hre.network)
    const address = await l1Signer.getAddress()
    console.log(`Using signer ${address}`)

    // Ensure that the signer has a balance before trying to
    // do anything
    const balance = await l1Signer.getBalance()
    console.log(`L1 Signer balance: ${formatEther(balance.toString())}`)

    const withdrawAmount = utils.parseEther(args.amount)

    const l2ChainId = await l2Signer.getChainId()
    let contractAddrs = CONTRACT_ADDRESSES[l2ChainId]
    if (args.l1ContractsJsonPath) {
      const data = await fs.readFile(args.l1ContractsJsonPath)
      const json = JSON.parse(data.toString())
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
        },
        l2: DEFAULT_L2_CONTRACT_ADDRESSES,
      } as OEContractsLike
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

    console.log(`OptimismPortal: ${contractAddrs.l1.OptimismPortal}`)
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

    const L2ToL1MessagePasser = new hre.ethers.Contract(
      predeploys.L2ToL1MessagePasser,
      Artifact__L2ToL1MessagePasser.abi
    )

    console.log(predeploys.L2CrossDomainMessenger)
    const L2CrossDomainMessenger = new hre.ethers.Contract(
      predeploys.L2CrossDomainMessenger,
      Artifact__L2CrossDomainMessenger.abi
    )

    const L2StandardBridge = new hre.ethers.Contract(
      predeploys.L2StandardBridge,
      Artifact__L2StandardBridge.abi
    )

    const messenger = new CrossChainMessenger({
      l1SignerOrProvider: l1Signer,
      l2SignerOrProvider: l2Signer,
      l1ChainId: await l1Signer.getChainId(),
      l2ChainId,
      bedrock: true,
      contracts: contractAddrs,
    })

    const opBalance= await l1Signer!.provider!.getBalance(
      OptimismPortal.address
    )

    const l1BridgeBalance= await l1Signer!.provider!.getBalance(
      L1StandardBridge.address
    )

    console.log(`L1StandardBridge balance: ${formatEther(l1BridgeBalance)}`)

    console.log(`Optimism Portal balance: ${formatEther(opBalance)}`)

    console.log('Withdrawing ETH')
    const ethWithdraw = await messenger.withdrawETH(withdrawAmount)
    console.log(`Transaction hash: ${ethWithdraw.hash}`)
    const ethWithdrawReceipt = await ethWithdraw.wait()
    console.log(
      `ETH withdrawn on L2 - included in block ${ethWithdrawReceipt.blockNumber}`
    )

    {
      // check the logs
      for (const log of ethWithdrawReceipt.logs) {
        switch (log.address) {
          case L2ToL1MessagePasser.address: {
            const parsed = L2ToL1MessagePasser.interface.parseLog(log)
            console.log(parsed.name)
            console.log(parsed.args)
            console.log()
            break
          }
          case L2StandardBridge.address: {
            const parsed = L2StandardBridge.interface.parseLog(log)
            console.log(parsed.name)
            console.log(parsed.args)
            console.log()
            break
          }
          case L2CrossDomainMessenger.address: {
            const parsed = L2CrossDomainMessenger.interface.parseLog(log)
            console.log(parsed.name)
            console.log(parsed.args)
            console.log()
            break
          }
          default: {
            console.log(`Unknown log from ${log.address} - ${log.topics[0]}`)
          }
        }
      }
    }

    console.log('Waiting to be able to prove withdrawal')

    const proveInterval = setInterval(async () => {
      const currentStatus = await messenger.getMessageStatus(ethWithdrawReceipt)
      console.log(`Message status: ${MessageStatus[currentStatus]}`)
      const latest = await L2OutputOracle.latestBlockNumber()
      console.log(
        `Latest L2OutputOracle commitment number: ${latest.toString()}`
      )
      const tip = await l1Signer.provider!.getBlockNumber()
      console.log(`L1 chain tip: ${tip.toString()}`)
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
    if (ethProveReceipt.status !== 1) {
      throw new Error('Prove withdrawal transaction reverted')
    }
    console.log('Successfully proved withdrawal')

    console.log('Waiting to be able to finalize withdrawal')

    const finalizeInterval = setInterval(async () => {
      const currentStatus = await messenger.getMessageStatus(ethWithdrawReceipt)
      console.log(`Message status: ${MessageStatus[currentStatus]}`)
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
      switch (log.address) {
        case OptimismPortal.address: {
          const parsed = OptimismPortal.interface.parseLog(log)
          if (parsed.name === 'WithdrawalProven') {
            requestId = parsed.args.requestId
          }
        }
      }
    }

    console.log('Fulfilling withdrawal request...')
    const adminSigner = new hre.ethers.Wallet(args.adminkey, l1Provider)
    const ethFulfill = await OptimismPortal.connect(adminSigner).finalize(
      requestId
    )
    console.log(`Transaction hash: ${ethFulfill.hash}`)
    const ethFulfillReceipt = await ethFulfill.wait()
    if (ethFulfillReceipt.status !== 1) {
      throw new Error('Fulfill withdrawal reverted')
    }

    console.log('Finalizing eth withdrawal...')
    const ethFinalize = await messenger.finalizeMessage(ethWithdrawReceipt)
    console.log(`Transaction hash: ${ethFinalize.hash}`)
    const ethFinalizeReceipt = await ethFinalize.wait()
    if (ethFinalizeReceipt.status !== 1) {
      throw new Error('Finalize withdrawal reverted')
    }

    console.log(
      `ETH withdrawal complete - included in block ${ethFinalizeReceipt.blockNumber}`
    )
    {
      // Check that the logs are correct
      for (const log of ethFinalizeReceipt.logs) {
        switch (log.address) {
          case L1StandardBridge.address: {
            const parsed = L1StandardBridge.interface.parseLog(log)
            console.log(parsed.name)
            console.log(parsed.args)
            console.log()
            if (
              parsed.name !== 'ETHBridgeFinalized' &&
              parsed.name !== 'ETHWithdrawalFinalized'
            ) {
              throw new Error('Wrong event name from L1StandardBridge')
            }
            if (!parsed.args.amount.eq(withdrawAmount)) {
              throw new Error('Wrong amount in event')
            }
            if (parsed.args.from !== address) {
              throw new Error('Wrong to in event')
            }
            if (parsed.args.to !== address) {
              throw new Error('Wrong from in event')
            }
            break
          }
          case L1CrossDomainMessenger.address: {
            const parsed = L1CrossDomainMessenger.interface.parseLog(log)
            console.log(parsed.name)
            console.log(parsed.args)
            console.log()
            if (parsed.name !== 'RelayedMessage') {
              throw new Error('Wrong event from L1CrossDomainMessenger')
            }
            break
          }
          case OptimismPortal.address: {
            const parsed = OptimismPortal.interface.parseLog(log)
            console.log(parsed.name)
            console.log(parsed.args)
            console.log()
            // TODO: remove this if check
            if (parsed.name === 'WithdrawalFinalized') {
              if (parsed.args.success !== true) {
                throw new Error('Unsuccessful withdrawal call')
              }
            }
            break
          }
          default: {
            console.log(`Unknown log from ${log.address} - ${log.topics[0]}`)
          }
        }
      }
    }

    const opBalanceFinally = await l1Signer!.provider!.getBalance(
      OptimismPortal.address
    )

    if (!opBalanceFinally.eq(opBalance)) {
      throw new Error('OptimismPortal balance mismatch')
    }
    console.log('Withdraw success')
  })
