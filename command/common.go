package command

import (
	"github.com/chr0n1x/yodelr/construct"
	"github.com/codegangsta/cli"
)

func defaultFlags() []cli.Flag {
	return []cli.Flag{}
}

func buildNotification(ping bool, args ...string) *construct.Notification {
	msgArgs := append([]string{}, args...)
	for len(msgArgs) < 3 {
		msgArgs = append(msgArgs, "")
	}

	return &construct.Notification{
		Message:  msgArgs[0],
		Title:    msgArgs[1],
		Subtitle: msgArgs[2],
		Ping:     ping,
	}
}
