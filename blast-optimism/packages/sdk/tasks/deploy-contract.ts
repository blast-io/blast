import { promises as fs } from 'fs'

import { task, types } from 'hardhat/config'
import '@nomiclabs/hardhat-ethers'
import 'hardhat-deploy'
import { providers, utils, ethers } from 'ethers'

const { formatEther } = utils
task('deploy-contract', 'deploy contract on L2')
  .addParam(
    'l2providerurl',
    'L2 provider URL.',
    'http://localhost:9545',
    types.string
  )
  .addParam('userkey', 'user deploying contract', '', types.string)
  .addOptionalParam('filepath', 'path to file', '/Users/Downloads/HelloWorld.json', types.string)
  .addOptionalParam('args', 'array of args', '["0xa6b8cfd7588f6d52b3913c520894facd3fbc1a1e"]', types.string)
  .setAction(async (args, hre) => {
    const constructorArgs = JSON.parse(args.args)
    const l2Provider = new providers.StaticJsonRpcProvider(args.l2providerurl)
    const l2Signer = new hre.ethers.Wallet(
      args.userkey,
      l2Provider
    )
    const contractJsonRaw = await fs.readFile(args.filepath)
    const contractJson = JSON.parse(contractJsonRaw.toString())
    const abi = contractJson.abi;
    const bytecode = contractJson.data.bytecode.object;
    const tContractFactory = new ethers.ContractFactory(abi, bytecode, l2Signer);
    // Deploy the contract
    // Arguments for the contract deployment
    const tContract = await tContractFactory.deploy(...constructorArgs, {gasLimit: 10000000});
  
    // Wait for deployment to finish
    await tContract.deployed();
    console.log("txn hash", tContract.deployTransaction.hash)

    console.log(`Contract deployed at address: ${tContract.address}`);
  })