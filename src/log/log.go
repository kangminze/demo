package log

import "github.com/sevenNt/wzap"

func Init()  {
	logconsole := map[string]interface{} {
		"color":true,
		"prefix":"[demo]",
		"level":"Debug",
	}
	logzap := map[string]interface{}{
		"path":"demo.log",
		"level":"Debug",
	}

	logger := wzap.New(
		wzap.WithOutputKV(logconsole),
		wzap.WithOutputKV(logzap),
	)
	wzap.SetDefaultLogger(logger)
}