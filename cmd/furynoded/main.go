package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/Furynet/furynode/app"
	"github.com/Furynet/furynode/cmd/furynoded/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	app.SetConfig(true)

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
