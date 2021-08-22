package controller

import (
	"fmt"

	"github.com/kyleu/projectforge/app/util"
	"github.com/kyleu/projectforge/views/vproject"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/projectforge/app/controller/cutil"

	"github.com/kyleu/projectforge/app"
)

func ProjectFileRoot(rc *fasthttp.RequestCtx) {
	act("project.file.root", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prj, err := getProject(rc, as)
		if err != nil {
			return "", err
		}

		ps.Title = fmt.Sprintf("%s (project %s)", prj.Title(), prj.Key)
		ps.Data = prj
		return render(rc, as, &vproject.Files{Project: prj}, ps, "projects", prj.Key)
	})
}

func ProjectFile(rc *fasthttp.RequestCtx) {
	act("project.file", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prj, err := getProject(rc, as)
		if err != nil {
			return "", err
		}

		pathS, err := rcRequiredString(rc, "path", false)
		if err != nil {
			return "", err
		}
		path := util.SplitAndTrim(pathS, "/")
		bcAppend := "||/p/" + prj.Key + "/fs"
		bc := []string{"projects", prj.Key, "Files" + bcAppend}
		for _, x := range path {
			bcAppend += "/" + x
			b := x + bcAppend
			bc = append(bc, b)
		}
		return render(rc, as, &vproject.Files{Project: prj, Path: path}, ps, bc...)
	})
}
