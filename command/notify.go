package command

import (
	"github.com/chr0n1x/yodelr/construct"
	"github.com/codegangsta/cli"

	log "github.com/Sirupsen/logrus"
)

func build(ping bool, args ...string) *construct.Notification {
	not := buildNotification(ping, args...)
	log.Debugf("%+v", not)
	return not
}

func pushNotification(context *cli.Context) {
	if len(context.Args()) < 1 {
		log.Fatal("message required")
	}
	build(context.Bool("ping"), context.Args()...).Show()
}

var Notify = cli.Command{
	Name:   "notify",
	Usage:  "push a notification, given a message and, optionally, title & subtitle",
	Action: pushNotification,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ping, p",
			Usage: "notify with a sound",
		},
	},
	Description: "Display a notification NOW.",
}
