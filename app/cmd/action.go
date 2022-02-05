package cmd

import (
	"context"
	"strings"

	"github.com/muesli/coral"
	"github.com/pkg/errors"

	"github.com/kyleu/projectforge/app/action"
	"github.com/kyleu/projectforge/app/util"
)

func actionF(ctx context.Context, t action.Type, args []string) error {
	if err := initIfNeeded(); err != nil {
		return errors.Wrap(err, "error initializing application")
	}

	_, cfg := extractConfig(args)
	logResult(t, runToCompletion(ctx, "", t, cfg))

	return nil
}

func actionCmd(ctx context.Context, t action.Type) *coral.Command {
	f := func(cmd *coral.Command, args []string) error { return actionF(ctx, t, args) }
	return &coral.Command{Use: t.Key, Short: t.Description, RunE: f}
}

func actionCommands() []*coral.Command {
	ctx := context.Background()
	ret := make([]*coral.Command, 0, len(action.AllTypes))
	for _, a := range action.AllTypes {
		if !a.Hidden {
			ret = append(ret, actionCmd(ctx, a))
		}
	}
	return ret
}

func logResult(t action.Type, r *action.Result) {
	_logger.Infof("%s [%s]: %s in [%s]", util.AppName, t.String(), r.Status, util.MicrosToMillis(r.Duration))
	if len(r.Errors) > 0 {
		_logger.Warnf("Errors:")
		for _, e := range r.Errors {
			_logger.Warn(" - " + e)
		}
	}
	if r.Modules.DiffCount(false) > 0 {
		for _, m := range r.Modules {
			for _, d := range m.DiffsFiltered(false) {
				_logger.Infof("%s [%s]:", d.Path, d.Status)
				for _, c := range d.Changes {
					for _, l := range c.Lines {
						_logger.Info(strings.TrimSuffix(l.String(), "\n"))
					}
				}
			}
		}
	}
}
