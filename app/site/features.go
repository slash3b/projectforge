package site

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/app/lib/menu"
	"projectforge.dev/projectforge/app/module"
	"projectforge.dev/projectforge/doc"
	"projectforge.dev/projectforge/views/layout"
	"projectforge.dev/projectforge/views/vsite"
)

func featuresMenu(ctx context.Context, mSvc *module.Service) menu.Items {
	ms := mSvc.Modules()
	ret := make(menu.Items, 0, len(ms))
	for _, m := range ms {
		ret = append(ret, &menu.Item{Key: m.Key, Title: m.Title(), Description: m.Description, Icon: m.Icon, Route: "/features/" + m.Key})
	}
	return ret
}

func featureList(as *app.State, ps *cutil.PageState) (layout.Page, error) {
	mods := as.Services.Modules.Modules()
	ps.Title = "Available Modules"
	ps.Data = mods
	return &vsite.FeatureList{Modules: mods}, nil
}

func featureDetail(key string, as *app.State, ps *cutil.PageState) (layout.Page, error) {
	mod, err := as.Services.Modules.Get(key)
	if err != nil {
		return nil, err
	}
	_, html, err := doc.HTMLString("feature:"+mod.Key, []byte(mod.UsageMD), func(s string) (string, string, error) {
		ret, err := cutil.FormatMarkdown(s)
		if err != nil {
			return "", "", err
		}
		if h1Idx := strings.Index(ret, "<h1>"); h1Idx > -1 {
			if h1EndIdx := strings.Index(ret, "</h1>"); h1EndIdx > -1 {
				ret = ret[:h1Idx] + ret[h1EndIdx+5:]
			}
		}
		return "", ret, nil
	})
	if err != nil {
		return nil, err
	}
	ps.Title = mod.Title()
	ps.Data = mod
	return &vsite.FeatureDetail{Module: mod, HTML: html}, nil
}

func featureFiles(path []string, as *app.State, ps *cutil.PageState) ([]string, layout.Page, error) {
	if len(path) < 3 {
		return path, nil, errors.New("invalid path")
	}
	if path[2] != "files" {
		return path, nil, errors.New("invalid file path")
	}
	u := "/features/" + path[1] + "/files"
	bc := append([]string{}, path[:2]...)
	bc = append(bc, "Files||"+u)
	for _, x := range path[3:] {
		u += "/" + x
		bc = append(bc, x+"||"+u)
	}
	mod, err := as.Services.Modules.Get(path[1])
	if err != nil {
		return path, nil, err
	}
	ps.Data = mod
	return bc, &vsite.FeatureFiles{Module: mod, Path: path[3:]}, nil
}
