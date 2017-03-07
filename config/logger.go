package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	prefixer "github.com/x-cray/logrus-prefixed-formatter"
)

// Log is an instance of Logrus logger
var Log = logrus.New()

func init() {
	if os.Getenv(EnvironmentVariableName) != Production {
		Log.Formatter = new(prefixer.TextFormatter)
	} else {
		Log.Formatter = new(ProductionFormatter)
	}
	Log.Level = logrus.InfoLevel
}

// ProductionFormatter.
type ProductionFormatter struct{}

// Format implements Logrus formatter.
func (f *ProductionFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	fields := ""
	if len(entry.Data) > 0 {
		fs := []string{}
		for k, v := range entry.Data {
			fs = append(fs, fmt.Sprintf("%s=%s", k, v))
		}
		fields = fmt.Sprintf(" (%s)", strings.Join(fs, ", "))
	}

	data := fmt.Sprintf("[%s] %+5s: %s%s\n",
		time.Now().Format(time.RFC1123),
		strings.ToUpper(entry.Level.String()),
		entry.Message,
		fields,
	)
	return []byte(data), nil
}
