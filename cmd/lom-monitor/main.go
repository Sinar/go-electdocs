package main

import (
	"fmt"
	"github.com/urfave/cli/v3"
)

// Use urfave/cli to have different profiles of service
// worker
// REST / proxy front?
// admin / operations - including a starter

func main() {
	fmt.Println("Welcome to LOM Monitoring!!! ..")
	Run()
}

func Run() {
	app := &cli.App{
		Name: "lom",
		Commands: []*cli.Command{
			{
				Name:        "admin",
				Aliases:     []string{"d"},
				Usage:       "use it to see a description",
				Description: "This is how we describe describeit the function",
				Subcommands: []*cli.Command{
					{
						Name:    "daily",
						Aliases: []string{"d"},
						Action: func(context *cli.Context) error {
							fmt.Println("Going daily .. daily ..")
							return nil
						},
					},
					{
						Name:    "repli",
						Aliases: []string{"r"},
						Action: func(context *cli.Context) error {
							fmt.Println("replu .. repli ..")
							return nil
						},
					},
				},
			},
			{
				Name:        "worker",
				Aliases:     []string{"c"},
				Usage:       "use it to see a Cescription",
				Description: "This is how we describe cescribeit the function",
			},
		},
	}
	_ = app.Run([]string{"lom", "d", "r"})
}
