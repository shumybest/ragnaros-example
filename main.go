package main

import (
	"fmt"
	"github.com/shumybest/ragnaros2"
	"github.com/urfave/cli/v2"
	"os"
	"ragnaros2-example/app"
	"sort"
)

func main() {
	// runOpts()
	runNoOpts()
}

func runNoOpts() {
	ragnaros.InjectApps(app.DemoController)
	ragnaros.Start()
}

func runOpts() {
	cliApp := cli.NewApp()
	cliApp.Name = "ragnaros spring cloud framework"
	cliApp.Usage = "[-p]"
	cliApp.Commands = Commands
	cliApp.Flags = Flags
	sort.Sort(cli.CommandsByName(cliApp.Commands))

	ragnaros.InjectApps(app.DemoController, func(r *ragnaros.Context) {
		fmt.Println("welcome to ragnaros")
	})

	err := cliApp.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
