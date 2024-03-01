Blast
==

This repo contains Blast's fork of optimism and op-geth. If you're interested in self hosting a Blast node for either mainnet or testnet, check out the docs [here](https://safe-violet-16b.notion.site/Blast-Deployment-Docs-b2f2b7b3c9a645fe8d6ce49fb963a467). 

## Running Blast locally (against a local L1)

### Prereqs
`docker`, `git`, `go1.20`, `node`, `pnpm`, `foundry`, `make`, `yarn`, `direnv`, `jq`

### Steps

1. Build the blast-geth docker image
```
docker build blast-geth -f blast-geth/Dockerfile -t blast-geth
```

2. Compile the smart contracts & start the devnet

```
export BLAST_ROOT=$(realpath blast-optimism)

# Compile the smart contracts & deploy them to the L1
cd blast-optimism/packages/contracts-bedrock
pnpm clean  
pnpm install
pnpm build

# If you're on a mac, you might need rosetta instructions, ONLY IF NEEDED FOR M1/M2:
# /usr/sbin/softwareupdate --install-rosetta --agree-to-license`

cd $BLAST_ROOT

# Install geth
make install-geth

# Make sure geth is in your PATH, if it's not, you can set it using
# export PATH="${PATH}:${GOPATH}/bin:${HOME}/go/bin"

make cannon-prestate

# Run the devnet in docker
# L1 RPC: http://localhost:8545
# L2 RPC: http://localhost:9545
make devnet-up

# To stop devnet run command
make devnet-down
```

### L1 faucet address

This is the private key for the devnet L1 faucet: `0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80`

This is the corresponding address: `0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266`



### Using the L2

1. Send some ETH from the L1 faucet to your wallet on the L1
2. Find the Blast bridge address using the command below:

```bash
cd $BLAST_ROOT
cat ./packages/contracts-bedrock/deployments/devnetL1/L1BlastBridgeProxy.json | grep -m 1 '"address": '
```

3. Send L1 ETH from your wallet to bridge address above
4. Wait for the tx to get confirmed
5. Your wallet should now have ETH on the L2

