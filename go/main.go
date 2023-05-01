package main

import (
	"C"
	"context"

	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
)

var api *ton.APIClient
var ctx context.Context
var wallets map[string]*wallet.Wallet = make(map[string]*wallet.Wallet)

func main() {}
