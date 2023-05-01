package main

import (
	"C"
	"strings"

	"strconv"

	"github.com/xssnick/tonutils-go/ton/wallet"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/tvm/cell"
)

//export UseWallet
func UseWallet(seed *C.char) {
	words := strings.Split(C.GoString(seed), " ")
	w, err := wallet.FromSeed(api, words, wallet.V4R2)
	if err != nil {
		panic(err)

	}
	wallets[C.GoString(seed)] = w
}

//export FreeWallet
func FreeWallet(seed *C.char) {
	delete(wallets, C.GoString(seed))
}

//export GetWalletAddress
func GetWalletAddress(seed *C.char) *C.char {
	return C.CString(wallets[C.GoString(seed)].Address().String())
}

//export GetWalletBalance
func GetWalletBalance(seed *C.char) *C.char {
	block, err := api.CurrentMasterchainInfo(ctx)
	if err != nil {
		panic(err)
	}

	balance, err := wallets[C.GoString(seed)].GetBalance(ctx, block)
	if err != nil {
		panic(err)
	}

	return C.CString(balance.String())
}

//export CreateWallet
func CreateWallet() *C.char {
	return C.CString(strings.Join(wallet.NewSeed(), " "))
}

//export Transfer
func Transfer(seed *C.char, addr *C.char, value *C.char, msg *C.char, mode *C.char) {
	ton_addr := address.MustParseAddr(C.GoString(addr))
	comment := C.GoString(msg)

	w := wallets[C.GoString(seed)]
	var body *cell.Cell
	var err error
	if comment != "" {
		body, err = wallet.CreateCommentCell(comment)
		if err != nil {
			panic(err)
		}
	}
	parsed_mode, err := strconv.Atoi(C.GoString(mode))
	if err != nil {
		panic(err)
	}

	err = w.Send(ctx, &wallet.Message{
		Mode: uint8(parsed_mode), // pay fees from msg value
		InternalMessage: &tlb.InternalMessage{
			IHRDisabled: true,
			Bounce:      true,
			DstAddr:     ton_addr,
			Amount:      tlb.MustFromTON(C.GoString(value)),
			Body:        body,
		},
	}, true)

	if err != nil {
		panic(err)
	}
}
