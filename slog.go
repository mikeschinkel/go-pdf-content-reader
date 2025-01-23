package pdf

import (
	slogPkg "log/slog"
	"os"
	"strings"
)

var slog *slogPkg.Logger

func SetSLog(sl *slogPkg.Logger) {
	slog = sl
}
func init() {
	SetSLog(NewSLog(LogFormat, &slogPkg.HandlerOptions{
		AddSource:   false,
		Level:       LogLevel,
		ReplaceAttr: nil,
	}))
}

var LogFormat = "text"
var LogLevel = slogPkg.LevelInfo
var LogWriter = os.Stdout

func NewSLog(format string, options *slogPkg.HandlerOptions) *slogPkg.Logger {
	var handler slogPkg.Handler
	switch strings.ToLower(format) {
	case "json":
		handler = slogPkg.NewJSONHandler(LogWriter, options)
	case "text":
		fallthrough
	default:
		handler = slogPkg.NewTextHandler(LogWriter, options)
	}
	return slogPkg.New(handler)
}
