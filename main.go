package main

import (
	"./cmd"
	"github.com/codegangsta/cli"
	"os"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	app := cli.NewApp()
	/*
		app.Name = setting.AppName
		app.Usage = setting.Usage
		app.Version = setting.Version
		app.Author = setting.Author
		app.Email = setting.Email
	*/
	app.Commands = []cli.Command{
		cmd.CmdWeb,
	}

	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
