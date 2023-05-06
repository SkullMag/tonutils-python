from ctypes import *
from os.path import join as path_join

__lib = CDLL(path_join("lib", "libton.so"))

__init_api = __lib.InitAPI
__init_api.argtypes = [c_char_p]

__use_wallet = __lib.UseWallet
__use_wallet.argtypes = [c_char_p]

__get_wallet_address = __lib.GetWalletAddress
__get_wallet_address.argtypes = [c_char_p]
__get_wallet_address.restype = c_char_p

__get_wallet_balance = __lib.GetWalletBalance
__get_wallet_balance.argtypes = [c_char_p]
__get_wallet_balance.restype = c_char_p

__create_wallet = __lib.CreateWallet
__create_wallet.restype = c_char_p

__free_wallet = __lib.FreeWallet
__free_wallet.argtypes = [c_char_p]

__transfer = __lib.Transfer
__transfer.argtypes = [c_char_p, c_char_p, c_char_p, c_char_p, c_char_p]


def init_api(config_url: str):
    __init_api(config_url.encode("utf8"))


def use_wallet(seed: str):
    __use_wallet(seed.encode("utf8"))


def create_wallet() -> str:
    return __create_wallet().decode("utf8")


def get_wallet_address(seed: str) -> str:
    return __get_wallet_address(seed.encode("utf8")).decode("utf8")


def get_wallet_balance(seed: str) -> float:
    return __get_wallet_balance(seed.encode("utf8")).decode("utf8")


def transfer(seed: str, address: str, value: float, msg: str, mode: int = 0):
    __transfer(
        seed.encode("utf8"), address.encode("utf8"), str(value).encode("utf8"), msg.encode("utf8"), str(mode).encode("utf8")
    )


def free_wallet(seed: str):
    __free_wallet(seed.encode("utf8"))
