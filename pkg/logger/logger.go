package logger

import (
	"fmt"
	"path"
	"runtime"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"github.com/zput/zxcTool/ztLog/zt_formatter"
)

// Logger is the exported logger, use it for logs
var Logger *logrus.Logger

func init() {
	var defaultFormatter = &zt_formatter.ZtFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
		Formatter: nested.Formatter{
			NoColors: true,
		},
	}

	Logger = logrus.New()
	Logger.SetFormatter(defaultFormatter)
	SetupHooks(Logger)
	Logger.SetLevel(logrus.DebugLevel)
	Logger.SetReportCaller(true)
	return
}
