// Content managed by Project Forge, see [projectforge.md] for details.
package app

import (
	"fmt"
	"time"

	"go.uber.org/zap"

	"projectforge.dev/app/lib/filesystem"
	"projectforge.dev/app/lib/telemetry"
	"projectforge.dev/app/lib/theme"
	"projectforge.dev/app/util"
)

type BuildInfo struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Date    string `json:"date"`
}

func (b *BuildInfo) String() string {
	if b.Date == "unknown" {
	} else {
		d, _ := util.TimeFromJS(b.Date)
		return fmt.Sprintf("%s (%s)", b.Version, util.TimeToYMD(d))
	}
	return b.Version
}

type State struct {
	Debug     bool
	BuildInfo *BuildInfo
	Files     filesystem.FileLoader
	Themes    *theme.Service
	Logger    *zap.SugaredLogger
	Services  *Services
	Started   time.Time
}

func NewState(debug bool, bi *BuildInfo, f filesystem.FileLoader, logger *zap.SugaredLogger) (*State, error) {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		return nil, err
	}
	time.Local = loc

	_ = telemetry.InitializeIfNeeded(true, logger)
	ts := theme.NewService(f, logger)

	return &State{
		Debug:     debug,
		BuildInfo: bi,
		Files:     f,
		Themes:    ts,
		Logger:    logger,
		Started:   time.Now(),
	}, nil
}
