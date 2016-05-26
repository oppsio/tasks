package main

import (
	"gopkg.in/urfave/cli.v2" // imports as package "cli"
	"os"
	"strings"
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
	app.Name = "tasks"
	app.UsageText = "tasks run [taskname]"
	app.Usage = "The CLI for running Oppsio tasks"
	app.Authors = []cli.Author{{Name: "Ricardo Rossi", Email: "ricardo@endata.com"}}
	app.Version = "0.9.1"
	app.Copyright = "© Oppsio 2016"

	app.Commands = []cli.Command{
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "the task to run",
			Action: func(c *cli.Context) error {
				taskName := strings.ToLower(c.Args().First())
				if t, ok := tasks[taskName]; ok {
					t.(task).Run()
				} else {
					return cli.NewExitError("task does not exist", 4)
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}
