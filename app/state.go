package app

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/kyleu/projectforge/app/auth"
	"github.com/kyleu/projectforge/app/filesystem"
	"github.com/kyleu/projectforge/app/telemetry"
	"github.com/kyleu/projectforge/app/theme"
	"github.com/kyleu/projectforge/app/util"
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
	Auth      *auth.Service
	Themes    *theme.Service
	Logger    *zap.SugaredLogger
	Services  *Services
}

func NewState(debug bool, bi *BuildInfo, f filesystem.FileLoader, logger *zap.SugaredLogger) (*State, error) {
	_ = telemetry.InitializeIfNeeded(true, logger)
	as := auth.NewService("", logger)
	ts := theme.NewService(f, logger)

	return &State{
		Debug:     debug,
		BuildInfo: bi,
		Files:     f,
		Auth:      as,
		Themes:    ts,
		Logger:    logger,
	}, nil
}
