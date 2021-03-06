package plist

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/chr0n1x/yodelr/construct"
)

// CalendarProperties containers information required to generate
// a calendar plist file
type CalendarProperties struct {
	construct.Notification
	BinPath  string
	Interval map[string]int
}

// Generate returns contents for a reminder plist file with the data
// given in this CalendarProperties struct
func (f *CalendarProperties) Generate(tempType string) ([]byte, error) {
	if f.BinPath == "" {
		path, _ := filepath.Abs(os.Args[0])
		f.BinPath = path
	}

	return createFromTemplate(fmt.Sprintf("templates/%s.plist", tempType), *f)
}
