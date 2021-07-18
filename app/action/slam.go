package action

import (
	"github.com/kyleu/projectforge/app/diff"
	"github.com/kyleu/projectforge/app/module"
	"github.com/kyleu/projectforge/app/project"
	"github.com/kyleu/projectforge/app/util"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func onSlam(prj *project.Project, mods module.Modules, cfg util.ValueMap, mSvc *module.Service, pSvc *project.Service, logger *zap.SugaredLogger) *Result {
	ret := newResult(cfg, logger)
	var res module.Results
	for _, mod := range mods {
		r, err := slam(prj, mod, true, mSvc, pSvc, logger)
		if err != nil {
			return ret.WithError(err)
		}
		res = append(res, r)
		logger.Info("applied module [" + mod.Key + "]")
	}
	ret.Modules = res
	return ret
}

func slam(prj *project.Project, mod *module.Module, addHeader bool, mSvc *module.Service, pSvc *project.Service, logger *zap.SugaredLogger) (*module.Result, error) {
	start := util.TimerStart()
	srcFiles, diffs, err := diffs(prj, mod, addHeader, mSvc, pSvc, logger)
	if err != nil {
		return nil, err
	}

	for _, f := range diffs {
		switch f.Status {
		case diff.StatusIdentical, diff.StatusMissing, diff.StatusSkipped:
			// noop
		case diff.StatusDifferent, diff.StatusNew:
			src := srcFiles.Get(f.Path)
			tgtFS := pSvc.GetFilesystem(prj)
			err := tgtFS.WriteFile(f.Path, []byte(src.Content), src.Mode, true)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to write updated content to [%s]", f.Path)
			}
		default:
			return nil, errors.Errorf("unhandled diff status [%s]", f.Status)
		}
	}

	return &module.Result{Key: mod.Key, Status: "OK", Diffs: diffs, Duration: util.TimerEnd(start)}, nil
}
