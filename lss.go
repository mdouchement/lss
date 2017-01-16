package main

import (
	"fmt"
	"os"

	"github.com/mdouchement/lss/config"
	"github.com/mdouchement/lss/engines"
	"github.com/mdouchement/lss/web"
	"gopkg.in/urfave/cli.v2"
)

var app *cli.App

func init() {
	app = &cli.App{}
	app.Name = "Light Storage Service"
	app.Version = config.Cfg.Version
	app.Usage = ""

	app.Commands = []*cli.Command{
		web.ServerCommand,
	}

	config.Engine = &engines.OS{
		Workspace: config.Cfg.Workspace,
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
