package checks

import (
	"strings"

	"go.uber.org/zap"
	"projectforge.dev/app/doctor"
	"projectforge.dev/app/util"
)

var repo = &doctor.Check{
	Key:     "repo",
	Section: "release",
	Title:   "Git Repo",
	Summary: "Verifies the git repository in the working directory",
	URL:     "https://git-scm.com",
	UsedBy:  "bin/build/release.sh",
	Fn:      checkRepo,
	Solve:   solveRepo,
}

func checkRepo(r *doctor.Result, logger *zap.SugaredLogger) *doctor.Result {
	exitCode, _, err := util.RunProcessSimple("git status", ".")
	if err != nil {
		return r.WithError(doctor.NewError("missing", "[git] is not present on your computer"))
	}
	if exitCode == 128 {
		return r.WithError(doctor.NewError("norepo", "no git repository in current directory"))
	}
	exitCode, out, err := util.RunProcessSimple("git status", ".")
	if err != nil {
		return r.WithError(doctor.NewError("error", "can't run [git status]: %+v", err))
	}
	if exitCode != 0 {
		return r.WithError(doctor.NewError("error", "can't run [git status]: %s", out))
	}
	if out = strings.TrimSpace(out); out == "" {
		return r.WithError(doctor.NewError("noremote", "no git remote configured", out))
	}
	exitCode, _, err = util.RunProcessSimple("git log -1", ".")
	if err != nil {
		return r.WithError(doctor.NewError("error", "can't run [git log]: %+v", err))
	}
	if exitCode == 128 {
		return r.WithError(doctor.NewError("nocommit", "git repo must have at least one commit"))
	}

	return r
}

func solveRepo(r *doctor.Result, logger *zap.SugaredLogger) *doctor.Result {
	if r.Errors.Find("norepo") != nil {
		r.AddSolution("run [git init] in this directory")
	}
	if r.Errors.Find("noremote") != nil {
		p, dr := loadRootProject(r, logger)
		dr.AddSolution("run [git remote add origin " + p.Info.Sourcecode + ".git]")
	}
	if r.Errors.Find("nocommit") != nil {
		r.AddSolution("run [git commit -am \"initial commit\"] in this directory")
	}
	return r
}
