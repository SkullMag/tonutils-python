# tonutils-python
Python bindings for [tonutils-go](https://github.com/xssnick/tonutils-go.git)

## Prerequisites
1. Go
2. Python

## Installation
Build tonutils-go static library for your system
```sh
make
```
This command will create file in directory `lib` named `libton.so`

## Usage
```python
from tonutils import init_api, Wallet

config_url = "https://ton-blockchain.github.io/testnet-global.config.json"

init_api(config_url)
address = ""
with Wallet(open("wallet.seed").read().strip()) as w:
    w.transfer(address, 0.1, "Shut up and take my money!", 128) # this will transfer all money
```
