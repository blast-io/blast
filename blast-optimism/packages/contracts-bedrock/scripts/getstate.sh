#!/usr/bin/env bash

pushd .
if [ $ETH_RPC_URL == "" ]; then
    echo "Set ETH_RPC_URL"
    exit 1
fi
if [ $1 == "" ]; then
    echo "Usage: getstate.sh BLOCK_NUMBER"
    exit 1
fi
nohup anvil --fork-url $ETH_RPC_URL --fork-block-number $1 --chain-id 1337 &

sleep 10

cd $(dirname $0) && python3 getstate.py ../dump.json

kill $!
popd

rm nohup.out
