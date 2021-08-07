package controller

import (
	"github.com/kyleu/projectforge/app/module"
	"github.com/kyleu/projectforge/views/vmodule"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/projectforge/app/controller/cutil"

	"github.com/kyleu/projectforge/app"
)

func ModuleList(ctx *fasthttp.RequestCtx) {
	act("module.root", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		mods := as.Services.Modules.Modules()
		ps.Title = "Module Listing"
		ps.Data = mods
		return render(ctx, as, &vmodule.List{Modules: mods}, ps, "modules")
	})
}

func ModuleDetail(ctx *fasthttp.RequestCtx) {
	act("module.detail", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		mod, err := getModule(ctx, as, ps)
		if err != nil {
			return "", err
		}
		return render(ctx, as, &vmodule.Detail{Module: mod}, ps, "modules", mod.Key)
	})
}

func getModule(ctx *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*module.Module, error) {
	key, err := ctxRequiredString(ctx, "key", true)
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
