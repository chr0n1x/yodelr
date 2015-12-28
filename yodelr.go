package main

import (
	"os"
	"os/user"
	"path"

	commands "github.com/chr0n1x/yodelr/command"

	log "github.com/Sirupsen/logrus"
	cli "github.com/codegangsta/cli"
)

var (
	appName = "yodelr"
	author  = "chr0n1x"
	email   = "heilong24@gmail.com"
	// replace using:
	// go build -ldflags "-X main.version $(git rev-parse HEAD)"
	version string
)

// set version of app depending on whether it is being compiled with ldflags or not
func setVersion(app *cli.App) {
	if len(version) > 0 {
		app.Version = version
	} else {
		app.Version = "DEV"
	}
}

// set global flags for the app
func setFlags(app *cli.App) {
	user, _ := user.Current()
	appHome := path.Join(user.HomeDir, "."+appName)
	appConfig := path.Join(appHome, "config.yaml")

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Config file to load",
			Value: appConfig,
		},
		cli.BoolFlag{
			Name:  "verbose, V",
			Usage: "Turn on debug/verbose logging.",
		},
	}
}

// set all commands
func setCommands(app *cli.App) {
	app.Commands = []cli.Command{
		commands.Notify,
	}
}

// set log level to debug if specified
func setLogLevel(c *cli.Context) error {
	if c.GlobalBool("verbose") {
		log.SetLevel(log.DebugLevel)
	}

	return nil
}

// create app with basic fields using vars above
func bootstrapApp() *cli.App {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = "Ping yoself."
	app.Author = author
	app.Email = email

	return app
}

// ruuuun program
func main() {
	app := bootstrapApp()

	setVersion(app)
	setFlags(app)
	setCommands(app)

	app.Before = setLogLevel

	app.Run(os.Args)
}
