package main

import (
	"fmt"
	"github.com/beyondbitcoinshow/glogchain/web"
	. "github.com/tendermint/go-common"
	cfg "github.com/tendermint/go-config"
	tmlog "github.com/tendermint/go-logger"
	tmcfg "github.com/tendermint/tendermint/config/tendermint"
	tendermintNode "github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/version"
	"github.com/tendermint/tmsp/server"
	"github.com/urfave/cli"
	"os"
)

var (
	config      cfg.Config
	DataDirFlag = utils.DirectoryFlag{
		Name:  "datadir",
		Usage: "Data directory for the databases and keystore",
		Value: utils.DirectoryString{DefaultDataDir()},
	}
)

const (
	clientIdentifier = "Glogchain"
	versionMajor     = 0
	versionMinor     = 3
	versionPatch     = 0
	versionMeta      = "unstable"
)

var (
	verString  string // Combined textual representation of all the version components
	cliApp     *cli.App
	mainLogger = logger.NewLogger("main")
)

func init() {
	verString = fmt.Sprintf("%d.%d.%d", versionMajor, versionMinor, versionPatch)
	if versionMeta != "" {
		verString += "-" + versionMeta
	}
	cliApp = newCliApp(verString, "the glogchain command line interface")
	cliApp.Action = GlogChainApp
	cliApp.Commands = []cli.Command{
		Action:      initCommand,
		Name:        "init",
		Usage:       "init genesis.json",
		Description: "",
	}
	cliApp.After = func(ctx *cli.Context) error {
		logger.Flush()
		return nil
	}
	logger.AddLogSystem(logger.NewStdLogSystem(os.Stdout, log.LstdFlags, logger.DebugLevel))
	glog.SetToStderr(true)

}

func main() {
	mainLogger.Infoln("Starting glogchain")
	if err := cliApp.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func initCommand(ctx *cli.Context) error {

	// Get configuration
	config := tmcfg.GetConfig(ctx)
	init_files()
}




	// set the log level
	logger.SetLogLevel(tmcfg.GetString("log_level"))

	genesisFile, err := os.Open(genesisPath)
	if err != nil {
		utils.Fatalf("failed to read genesis file: %v", err)
	}

	block, err := core.WriteGenesisBlock(chainDb, genesisFile)
	if err != nil {
		utils.Fatalf("failed to write genesis block: %v", err)
	}
	glog.V(logger.Info).Infof("successfully wrote genesis block and/or chain rule set: %x", block.Hash())
	return nil

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
	)

func getTendermintConfig(ctx *cli.Context) cfg.Config {
	datadir := ctx.GlobalString(DataDirFlag.Name)
	os.Setenv("TMROOT", datadir)
	config = tmcfg.GetConfig("")
	config.Set("node_laddr", ctx.GlobalString("node_laddr"))
	config.Set("seeds", ctx.GlobalString("seeds"))
	config.Set("fast_sync", ctx.GlobalBool("no_fast_sync"))
	config.Set("skip_upnp", ctx.GlobalBool("skip_upnp"))
	config.Set("rpc_laddr", ctx.GlobalString("rpc_laddr"))
	config.Set("proxy_app", ctx.GlobalString("addr"))
	config.Set("log_level", ctx.GlobalString("log_level"))

	tmlog.SetLogLevel(config.GetString("log_level"))

	return config
}


