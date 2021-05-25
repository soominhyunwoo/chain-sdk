package main

import (
	"os"

	"github.com/soominhyunwoo/chain-sdk/server"
	svrcmd "github.com/soominhyunwoo/chain-sdk/server/cmd"
	"github.com/soominhyunwoo/chain-sdk/simapp"
	"github.com/soominhyunwoo/chain-sdk/simapp/simd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, simapp.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
