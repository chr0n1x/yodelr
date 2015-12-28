package command

import (
	"github.com/chr0n1x/yodelr/construct"
	"github.com/codegangsta/cli"

	log "github.com/Sirupsen/logrus"
)

func build(ping bool, args ...string) *construct.Notification {
	msgArgs := append([]string{}, args...)
	for len(msgArgs) < 3 {
		msgArgs = append(msgArgs, "")
	}

	log.Debugf("[Message Title Subtitle]: %+v", msgArgs)

	return &construct.Notification{
		Message:  msgArgs[0],
		Title:    msgArgs[1],
		Subtitle: msgArgs[2],
		Ping:     ping,
	}
}

func pushNotification(context *cli.Context) {
	if len(context.Args()) < 1 {
		log.Fatal("message required")
	}
	build(context.Bool("ping"), context.Args()...).Show()
}

var Notify = cli.Command{
	Name:   "notify",
	Usage:  "notify MESSAGE [TITLE] [SUBTITLE]",
	Action: pushNotification,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ping, p",
			Usage: "notify with a sound",
		},
	},
	Description: "Display a notification NOW.",
}
