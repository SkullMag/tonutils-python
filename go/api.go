package main

import (
	"C"
	"context"

	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
)

//export InitAPI
func InitAPI(configUrl *C.char) {
	client := liteclient.NewConnectionPool()
	ctx = client.StickyContext(context.Background())

	err := client.AddConnectionsFromConfigUrl(ctx, C.GoString(configUrl))
	if err != nil {
		panic(err)
	}
	api = ton.NewAPIClient(client)
}
