package controller

import (
	"github.com/valyala/fasthttp"
	"projectforge.dev/app/module"
	"projectforge.dev/app/project"
	"projectforge.dev/views/vmodule"

	"projectforge.dev/app/controller/cutil"

	"projectforge.dev/app"
)

func ModuleList(rc *fasthttp.RequestCtx) {
	act("module.root", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		mods := as.Services.Modules.Modules()
		ps.Title = "Module Listing"
		ps.Data = mods
		return render(rc, as, &vmodule.List{Modules: mods}, ps, "modules")
	})
}

func ModuleDetail(rc *fasthttp.RequestCtx) {
	act("module.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		mod, err := getModule(rc, as, ps)
		if err != nil {
			return "", err
		}
		var usages project.Projects
		for _, p := range as.Services.Projects.Projects() {
			if p.HasModule(mod.Key) {
				usages = append(usages, p)
			}
		}
		return render(rc, as, &vmodule.Detail{Module: mod, Usages: usages}, ps, "modules", mod.Key)
	})
}

func getModule(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*module.Module, error) {
	key, err := RCRequiredString(rc, "key", true)
	if err != nil {
		return nil, err
	}

	mod, err := as.Services.Modules.Get(key)
	if err != nil {
		return nil, err
	}

	ps.Title = mod.Title()
	ps.Data = mod

	return mod, nil
}
