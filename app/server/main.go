/*
Copyright © 2025 phenix3443
*/
package main

import (
	"os"

	"github.com/ethereum/go-ethereum/log"

	"github.com/phenix3443/dex-trade-bot/app/server/cmd"
)

func main() {
	log.SetDefault(log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stderr, log.LevelDebug, true)))
	cmd.Execute()
}
