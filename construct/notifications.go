package construct

import "github.com/everdev/mack"

type Notification struct {
	Title    string
	Message  string
	Subtitle string
	Ping     bool
}

func (n *Notification) Show() error {
	ping := ""
	if n.Ping {
		ping = "Ping"
	}

	return mack.Notify(n.Message, n.Title, n.Subtitle, ping)
}
