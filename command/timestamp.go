package command

import (
	"fmt"
	"time"

	"github.com/urfave/cli/v2"
)

func Timestamp() *cli.Command {
	return &cli.Command{
		Name:   "timestamp",
		Aliases: []string{"ts"},
		Action: actionTimestamp,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name: "f",
			},
		},
	}
}

func actionTimestamp(c *cli.Context) error {
	f := c.Bool("f")

	if f {
		for {
			fmt.Println(time.Now().Unix())
			time.Sleep(time.Second)
		}
	} else {
		fmt.Println(time.Now().Unix())
	}
	return nil
}
