from tonutils import init_api, Wallet

config_url = "https://ton-blockchain.github.io/testnet-global.config.json"

init_api(config_url)
address = ""
with Wallet(open("wallet.seed").read().strip()) as w:
    w.transfer(address, 0.1, "Shut up and take my money!", 128)
