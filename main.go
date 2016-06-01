package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v2" // imports as package "cli"
	"os"
	"strings"
)

// The list of available runnable tasks
var tasks = map[string]task{
	"fetch.site": task{
		read:  "settings",
		write: "links",
		task:  fetchSite{},
	},
	"fetch.page": task{
		read:  "links",
		write: "pages",
		task:  fetchPage{},
	},
	"parse.page": task{
		read:  "pages",
		write: "jobs",
		task:  parsePage{},
	},
}

// all tasks must implement this interface
type Runner interface {
	Run()
}

// the task
type task struct {
	read  string
	write string
	task  Runner
}

func main() {
	app := cli.NewApp()
	app.Name = "tasks"
	app.UsageText = "tasks run [taskname]"
	app.Usage = "The CLI for running Oppsio tasks"
	app.Authors = []cli.Author{{Name: "Ricardo Rossi", Email: "ricardo@endata.com"}}
	app.Version = "0.9.1"
	app.Copyright = "© Oppsio - 2016"

	app.Commands = []cli.Command{
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "the task to run",
			Action: func(c *cli.Context) error {
				taskKey := strings.ToLower(c.Args().First())
				if t, ok := tasks[taskKey]; ok {
					t.task.(Runner).Run()
				} else {
					return cli.NewExitError("task does not exist", 4)
				}
				return nil
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "displays a list of available tasks",
			Action: func(c *cli.Context) error {
				fmt.Println("\nAVAILABLE TASKS:")
				for k := range tasks {
					fmt.Println("   " + k)
				}
				fmt.Println("")
				return nil
			},
		},
	}

	app.Run(os.Args)
}
