package main

import (
	"github.com/urfave/cli"
)

const Version string = "0.1.0"

func main() {
	var env = GetEnv()
	newApp().Run(os.Args)
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "outenv"
	app.Usage = "manage the environment without polluting the directory"
	app.Version = Version
	app.Author = "nakabonne"
	app.Email = "rodriguez.nak@gmail.com"

	return app
}
