package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "golapisle"
	app.Usage = "API to store hiera variables"
	app.Version = "2.0.0"
	app.Run(os.Args)
}