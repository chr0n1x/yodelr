package construct

import "github.com/everdev/mack"

// Notification houses all information to pass to mack.Notify
type Notification struct {
	Title    string
	Message  string
	Subtitle string
	Ping     bool
}

// Show pushes a notification to the desktop
func (n *Notification) Show() error {
	ping := ""
	if n.Ping {
		ping = "Ping"
	}

	return mack.Notify(n.Message, n.Title, n.Subtitle, ping)
}
