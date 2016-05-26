package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v2" // imports as package "cli"
	"os"
)

// all tasks must implement this interface
type task interface {
	Run()
}

// a map of tasks to run
var tasks = make(map[string]task)

// register the available tasks
func register() {
	tasks["fetch.site"] = fetchSite{}
}

func main() {
	register()
	app := cli.NewApp()
	app.Name = "oppsio"
	app.UsageText = "oppsio --task [taskname]"
	app.Usage = "The CLI for running Oppsio tasks"
	app.Version = "0.9.1"
	app.Copyright = "© Oppsio 2016"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Aliases: []string{"t"},
			Name:    "task",
			Usage:   "which task to run",
		},
	}
	app.Action = func(c *cli.Context) error {
		fmt.Println("boom! I say!")
		return nil
	}

	app.Run(os.Args)
	// tasks["fetch.site"].(task).Run()
	// cli.NewApp().Run(os.Args)
	// fmt.Println(tasks["hhh"])
}
