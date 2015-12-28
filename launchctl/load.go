package launchctl

import "os/exec"

// Reload - launchctl load -w ...
func Reload(file string) error {
	return exec.Command("launchctl", "load", "-w", file).Run()
}
