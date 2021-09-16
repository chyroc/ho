package main

import (
	"log"
	"os"

	"github.com/chyroc/ho/command"
	"github.com/urfave/cli/v2"
)


func main() {
	app := &cli.App{
		Name:  "ho",
		Action: func(c *cli.Context) error {
			return cli.ShowAppHelp(c)
		},
		Commands: []*cli.Command{
			command.Timestamp(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}



