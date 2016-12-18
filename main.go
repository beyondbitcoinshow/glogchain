package main

import (
	"flag"

	"github.com/beyondbitcoinshow/glogchain/config"
	"github.com/beyondbitcoinshow/glogchain/web"
	. "github.com/tendermint/go-common"
	"github.com/tendermint/tmsp/server"
)

func main() {
	addrPtr := flag.String("addr", config.GlogchainConfigGlobal.TmspAddr, "Listen address")

	flag.Parse()
	app := NewGlogChainApp()

	// Start the listener
	_, err := server.NewServer(*addrPtr, "grpc", app)
	if err != nil {
		Exit(err.Error())
	}

	// start web server on port 8000
	go web.StartWebServer()

	// Wait forever
	TrapSignal(func() {
		// Cleanup
	})
}
