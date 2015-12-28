package plist

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"os/user"
	"path"
	"reflect"
	"text/template"

	"github.com/chr0n1x/yodelr/templates"
)

func generatePropertyMD5(args ...string) string {
	bytes := []byte{}
	for _, str := range args {
		bytes = append(bytes, []byte(str)...)
	}

	return fmt.Sprintf("%x", md5.Sum(bytes))
}

func createFromTemplate(path string, data interface{}) ([]byte, error) {
	t := template.New("yodelr-template")
	buffer := bytes.NewBufferString("")

	content, err := templates.Asset(path)
	if err != nil {
		return buffer.Bytes(), err
	}

	t, err = t.Parse(string(content))
	if err != nil {
		return buffer.Bytes(), err
	}

	err = t.Execute(buffer, data)
	return buffer.Bytes(), err
}

// GeneratePath returns the full path for a yodelr plist file, combining
// the type & field names:value pairs into a hash
func GeneratePath(typ string, data interface{}) string {
	params := []string{}
	dataIter := reflect.ValueOf(data).Elem()
	for i := 0; i < dataIter.NumField(); i++ {
		str := fmt.Sprintf("%s-%s", dataIter.Field(i), dataIter.Interface())
		params = append(params, str)
	}

	user, _ := user.Current()

	return path.Join(
		user.HomeDir,
		fmt.Sprintf(
			"/Library/LaunchAgents/com.yodelr.%s.%s.plist",
			typ,
			generatePropertyMD5(params...),
		),
	)
}
