from tonutils.ton import free_wallet, get_wallet_address, get_wallet_balance, transfer, use_wallet


class Wallet:
    def __init__(self, seed: str):
        self.__seed = seed

    @property
    def address(self):
        return get_wallet_address(self.__seed)

    @property
    def balance(self):
        return get_wallet_balance(self.__seed)

    def transfer(self, address: str, value: float, msg: str = "", mode: int = 0):
        transfer(self.__seed, address, value, msg, mode)

    def __enter__(self):
        use_wallet(self.__seed)
        return self

    def __exit__(self, *_):
        free_wallet(self.__seed)
