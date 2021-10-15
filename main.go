package main

import (
	"os"

	"github.com/kyleu/projectforge/app"
	"github.com/kyleu/projectforge/app/cmd"
	"github.com/kyleu/projectforge/app/log"
)

var (
	version = "0.1.4" // updated by bin/tag.sh and ldflags
	commit  = ""
	date    = "unknown"
)

func main() {
	logger, err := cmd.Run(&app.BuildInfo{Version: version, Commit: commit, Date: date})
	if err != nil {
		msg := "exiting due to error"
		if logger == nil {
			println(log.Red.Add(err.Error())) //nolint
			println(log.Red.Add(msg))         //nolint
		} else {
			logger.Error(err)
			logger.Error(msg)
		}
		os.Exit(1)
	}
}
