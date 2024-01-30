import http.client
import json
import sys
sys.set_int_max_str_digits(0)

from panoramix.decompiler import decompile_bytecode

def get_storage_slots(address):
    decoded = decompile_bytecode(get_code(address, "latest"))
    data = decoded.json
    slots = set()
    if "stor_defs" in data.keys():
      for storage_def in data["stor_defs"]:
          slots.add(hex(storage_def[2]))
    return slots

def call(method, params):
    packed_params = ",".join(map(lambda x: '"'+x+'"', params))

    conn = http.client.HTTPConnection("127.0.0.1:8545")
    headers = {'Content-type': 'application/json'}
    body = '{"id":3, "jsonrpc":"2.0", "method": ' \
            + '"' + method + '"' \
            + ', "params":[' + \
            packed_params + \
            ']}'
    conn.request('POST', '/', body, headers)
    response = conn.getresponse()
    data = response.read().decode()
    conn.close()
    jsonResult = json.loads(data)
    if "result" in jsonResult.keys():
      return jsonResult["result"]
    else:
      raise Exception(method + " failed: " + str(jsonResult))

def get_storage_proxy(proxy, impl, block_number):
    storage = {}
    for slot in [*get_storage_slots(proxy), *get_storage_slots(impl), *specific_slots]:
        slot = zero_pad(slot)
        value = zero_pad(call('eth_getStorageAt', [proxy, slot, block_number]))
        if (int(value, 16) > 0):
            storage[slot] = value
    return storage

def get_storage(address, block_number):
    storage = {}
    for slot in [*get_storage_slots(address), *specific_slots]:
        slot = zero_pad(slot)
        value = zero_pad(call('eth_getStorageAt', [address, slot, block_number]))
        if (int(value, 16) > 0):
            storage[slot] = value
    return storage

def get_code(address, block_number):
    return call("eth_getCode", [address, block_number])

def get_balance(address, block_number):
    return call("eth_getBalance", [address, block_number])

def zero_pad(x):
    return "0x" + x.lstrip("0x").rjust(64, "0")

contracts = [
    ["0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", "0x43506849d7c04f9138d1a2050bbf3a0c054402dd"], # USDC
    "0xdAC17F958D2ee523a2206206994597C13D831ec7", # USDT
    "0xbEbc44782C7dB0a1A60Cb6fe97d0b483032FF1C7", # Curve 3Pool
    "0x6c3F90f043a72FA612cbac8115EE7e52BDe6E490", # Curve 3Pool Token
    "0x6B175474E89094C44Da98b954EedeAC495271d0F", # DAI
    "0x373238337Bfe1146fb49989fc222523f83081dDb", # DSR_MANAGER
    "0x89B78CfA322F6C5dE0aBcEecab66Aee45393cC5A", # PSM
    "0x9759A6Ac90977b93B58547b4A71c78317f391A28", # DAI JOIN
    "0x0A59649758aa4d66E25f08Dd01271e891fe52199", # GEM JOIN
    "0xA191e578a6736167326d05c119CE0c90849E84B7", # GEM JOIN
    "0x35D1b3F3D7966A1DFe207aa4514C12a259A0492B", # VAT
    "0x197E90f9FAD81970bA7976f33CbD77088E5D7cf7", # POT
    ["0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84", "0x17144556fd3424EDC8Fc8A4C940B2D04936d17eb"], # LIDO
    ["0xb8ffc3cd6e7cf5a098a1c92f48009765b24088dc", "0x2b33CF282f867A7FF693A66e11B0FcC5552e4425"], # Kernel Proxy
    ["0x889edC2eDab5f40e902b864aD4d7AdE8E412F9B1", "0xE42C659Dc09109566720EA8b2De186c2Be7D94D9"], # Lido WithdrawalQueue
    ["0xB9D7934878B5FB9610B3fE8A5e441e8fad7E293f", "0xCC52f17756C04bBa7E377716d7062fC36D7f69Fd"], # Lido WithdrawalVault
    # "", ## Lido Locator
]

specific_slots = [
    "0x54b2b2de1ae6731a04bdbca30cee71852851cfcd3298aaf29f4ebff9452b27ad", # Kernel Proxy apps array
    "0x9fa7d1a90466effd74b6142b5e6b4b076451f75604d0878dc681cb2014f26d08", # DAI auth
    "0x9ef78dff90f100ea94042bd00ccb978430524befc391d3e510b5f55ff3166df7",
    "0xa3678de4a579be090bed1177e0a24f77cc29d181ac22fd7688aca344d8938015",
    "0xed310af23f61f96daefbcd140b306c0bdbf8c178398299741687b90e794772b0",
    "0xe6e35175eb53fc006520a2a9c3e9711a7c00de6ff2c32dd31df8c5a24cac1b5c",
    "0xa66d35f054e68143c18f32c990ed5cb972bb68a68f500cd2dd3a16bbf3686483",
    "0x9f70001d82b6ef54e9d3725b46581c3eb9ee3aa02b941b6aa54d678a9ca35b10",
    "0xafe016039542d12eec0183bb0b1ffc2ca45b027126a494672fba4154ee77facb",
    "0x644132c4ddd5bb6f0655d5fe2870dcec7870e6be4758890f366b83441f9fdece",
    "0x9fa7d1a90466effd74b6142b5e6b4b076451f75604d0878dc681cb2014f26d08",
    "0xed87b2ba768f47eacdb61813cec5387e5eb702b565075864e8450b2da9984eca",
]

def main(out_file, block_number):
    print("Dumping state at block:", block_number)
    dump = {}
    for contract in contracts:
        print("Dumping contract", contract)
        if isinstance(contract, list):
          [proxy, impl] = contract;
          dump[proxy] = {
              "code": get_code(proxy, block_number),
              "balance": get_balance(proxy, block_number),
              "storage": get_storage_proxy(proxy, impl, block_number),
              "nonce": 1
          }
          dump[impl] = {
              "code": get_code(impl, block_number),
              "balance": get_balance(impl, block_number),
              "storage": get_storage(impl, block_number),
              "nonce": 1
          }
        else:
          dump[contract] = {
              "code": get_code(contract, block_number),
              "balance": get_balance(contract, block_number),
              "storage": get_storage(contract, block_number),
              "nonce": 1
          }

    with open(out_file, "w") as f:
        json.dump({ "accounts": dump }, f, indent='  ')

block_number = "latest"
if (len(sys.argv) == 3):
  if sys.argv[2].startswith("0x"):
    block_number = sys.argv[2]
  else:
    block_number = hex(int(sys.argv[2]))

main(sys.argv[1], block_number)
