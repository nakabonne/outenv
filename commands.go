package main

import (
	"fmt"

	"github.com/urfave/cli"
)

var Commands = []cli.Command{
	commandSet,
}

var commandSet = cli.Command{
	Name:  "set",
	Usage: "Set environment variables of directory",
	Description: `
    Set environment variables of directory
`,
	Action: setSpecificEnv,
	Flags: []cli.Flag{
		cli.BoolFlag{Name: "update, u", Usage: "Update local repository if cloned already"},
	},
}

func setSpecificEnv(c *cli.Context) error {
	_ = c.Args().Get(0)
	_ = c.Bool("update")
	fmt.Println("hey")
	return nil
}
