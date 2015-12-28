package command

import (
	"io/ioutil"

	"github.com/chr0n1x/yodelr/launchctl"
	"github.com/chr0n1x/yodelr/plist"

	"github.com/codegangsta/cli"

	log "github.com/Sirupsen/logrus"
)

func startReminder(context *cli.Context) {
	note := buildNotification(context.Bool("ping"), context.Args()...)
	props := plist.CalendarProperties{}
	props.Notification = *note

	template, err := props.Generate()
	if err != nil {
		log.Fatalf("Could not generate reminder: %+v", err)
	}

	filepath := plist.GeneratePath("reminder", &props)

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
	Name:        "remind",
	Usage:       "dynamically spin off a reminder",
	Action:      startReminder,
	Flags:       defaultFlags(),
	Description: "Spin off a reminder.",
}
