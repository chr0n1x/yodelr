package command

import (
	"io/ioutil"

	"github.com/chr0n1x/yodelr/launchctl"
	"github.com/chr0n1x/yodelr/plist"
	"github.com/chr0n1x/yodelr/time"

	"github.com/codegangsta/cli"

	log "github.com/Sirupsen/logrus"
)

func startReminder(context *cli.Context) {
	if context.String("every") != "" && context.String("in") != "" {
		log.Fatal("Can only specify '--every' or '--in', not both")
	}

	reminderType := "every"
	reminderVal := "minute"
	if context.String("every") != "" || context.String("in") != "" {
		if context.String("in") != "" {
			reminderType = "in"
		}
		reminderVal = context.String(reminderType)
	}

	note := buildNotification(context.Bool("ping"), context.Args()...)
	props := plist.CalendarProperties{
		Interval: time.ToInterval(reminderVal),
	}
	props.Notification = *note

	template, err := props.Generate(reminderType)
	if err != nil {
		log.Fatalf("Could not generate reminder: %+v", err)
	}

	filepath := plist.GeneratePath(reminderType, &props)

	log.Infof("Writing %s", filepath)
	log.Debugf("Contents\n%s", template)

	err = ioutil.WriteFile(filepath, template, 0644)
	if err != nil {
		log.Fatalf("Could not write %s: %+v", filepath, err)
	}

	err = launchctl.Reload(filepath)
	if err != nil {
		log.Fatalf("Failed to reload launchctl: %+v", err)
	}
}

// Remind - generate a plist file that invokes the notify command
var Remind = cli.Command{
	Name:   "remind",
	Usage:  "dynamically spin off a reminder",
	Action: startReminder,
	Flags: append(
		defaultFlags(),
		cli.BoolFlag{
			Name:  "ping, p",
			Usage: "notify with a sound",
		},
		cli.StringFlag{
			Name:  "every, e",
			Usage: "Repeating reminder.",
		},
		cli.StringFlag{
			Name:  "in, i",
			Usage: "Remind once.",
		},
	),
	Description: "Spin off a reminder.",
}
