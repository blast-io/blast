#!/usr/bin/env bash

ADDRESS=0x4200000000000000000000000000000000000016
SLOT=$1
BLOCK_NUMBER=0x$(echo "obase=16; $2" | bc)

getBlockByNumber() {
    cast rpc --rpc-url http://127.0.0.1:9545 eth_getBlockByNumber $1 true
}

getProof() {
    addr="[\"$2\"]"
    cast rpc --rpc-url http://127.0.0.1:9545 eth_getProof $1 $addr $3
}

BLOCK=$(getBlockByNumber $BLOCK_NUMBER)
PROOF=$(getProof $ADDRESS $SLOT $BLOCK_NUMBER)

proof=$(echo $PROOF | jq '.storageProof[0].proof')
storageHash=$(echo $PROOF | jq '.storageHash')
stateRoot=$(echo $BLOCK | jq '.stateRoot')
hash=$(echo $BLOCK | jq '.hash')

echo "{ \"proof\": $proof, \"storageHash\": $storageHash, \"stateRoot\": $stateRoot, \"hash\": $hash }"
