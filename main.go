package main

import (
	"fmt"
	"github.com/beyondbitcoinshow/glogchain/protocol"
	"github.com/beyondbitcoinshow/glogchain/web"
	. "github.com/tendermint/go-common"
	cfg "github.com/tendermint/go-config"
	tmlog "github.com/tendermint/go-logger"
	tmcfg "github.com/tendermint/tendermint/config/tendermint"
	tendermintNode "github.com/tendermint/tendermint/node"
	"github.com/tendermint/tmsp/server"
	"github.com/urfave/cli"
	"log"
	"os"
)

const (
	clientIdentifier = "Glogchain"
	versionMajor     = 0
	versionMinor     = 3
	versionPatch     = 0
	versionMeta      = "insane"
)

var (
	verString  string // Combined textual representation of all the version components
	cliApp     *cli.App
	mainLogger = tmlog.New("main")
)

//Init prints the version number and gets this show on the road by setting up cliApp
func init() {
	verString = fmt.Sprintf("%d.%d.%d", versionMajor, versionMinor, versionPatch)
	if versionMeta != "" {
		verString += "-" + versionMeta
	}

	cliApp = cli.NewApp(verString, "glogchain CLI, your blockchain speaking")
	cliApp.Commands = []cli.Command{
		{
			Action:      "initCommand",
			Name:        "init",
			Usage:       "init genesis.json",
			Description: "",
		},
	}
	cliApp.HideVersion = true
	logger.AddLogSystem(tmlog.NewStdLogSystem(os.Stdout, log.LstdFlags, tmlog.DebugLevel))

}

//main gets the chain up and running.
func main() {
	if err := cliApp.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	configz := getTendermintConfig()
	tendermintNode.RunNode(configz)

	// Start the listener
	_, err := server.NewServer("0.0.0.0", "grpc", app)
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

func getTendermintConfig() cfg.Config {
	os.Setenv("TMROOT", "/home/faddat/data")
	configz := tmcfg.GetConfig("")
	configz.Set("node_laddr", "0.0.0.0")
	configz.Set("seeds", "138.201.198.167,138.201.198.169,138.201.198.173,138.201.198.175")
	configz.Set("rpc_laddr", "0.0.0.0")
	configz.Set("proxy_app", "0.0.0.0")
	configz.Set("log_level", "info")
	tmlog.SetLogLevel(config.GetString("log_level"))
	return configz
}
