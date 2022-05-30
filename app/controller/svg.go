package controller

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/app/lib/filesystem"
	"projectforge.dev/projectforge/app/project"
	"projectforge.dev/projectforge/app/svg"
	"projectforge.dev/projectforge/views/vsvg"
)

func SVGList(rc *fasthttp.RequestCtx) {
	act("svg.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prj, err := getProject(rc, as)
		if err != nil {
			return "", err
		}
		fs := as.Services.Projects.GetFilesystem(prj)

		icons, contents, err := svg.Contents(fs, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to list project SVGs")
		}

		ps.Title = "SVG Tools"
		ps.Data = icons
		return render(rc, as, &vsvg.List{Project: prj, Keys: icons, Contents: contents}, ps, "projects", prj.Key, "SVG")
	})
}

func SVGDetail(rc *fasthttp.RequestCtx) {
	act("svg.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prj, fs, key, err := prjAndIcon(rc, as)
		if err != nil {
			return "", err
		}
		content, err := svg.Content(fs, key)
		if err != nil {
			return "", errors.Wrap(err, "unable to read SVG ["+key+"]")
		}
		x := &svg.SVG{Key: key, Markup: content}
		ps.Title = "SVG [" + key + "]"
		ps.Data = x
		return render(rc, as, &vsvg.View{Project: prj, SVG: x}, ps, "projects", prj.Key, "SVG||/svg/"+prj.Key, key)
	})
}

func SVGBuild(rc *fasthttp.RequestCtx) {
	act("svg.build", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prj, err := getProject(rc, as)
		if err != nil {
			return "", err
		}
		count, err := svg.Build(as.Services.Projects.GetFilesystem(prj), ps.Logger)
		if err != nil {
			return "", err
		}
		msg := fmt.Sprintf("Parsed [%d] SVG files", count)
		return flashAndRedir(true, msg, "/svg/"+prj.Key, rc, ps)
	})
}

func SVGAdd(rc *fasthttp.RequestCtx) {
	act("svg.add", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		qa := rc.URI().QueryArgs()
		src := string(qa.Peek("src"))
		if src == "" {
			return ersp("must provide [src]")
		}
		tgt := string(qa.Peek("tgt"))
		if tgt == "" {
			tgt = strings.TrimSuffix(src, "-solid")
		}
		prj, err := getProject(rc, as)
		if err != nil {
			return "", err
		}
		fs := as.Services.Projects.GetFilesystem(prj)

		x, err := svg.AddToProject(fs, src, tgt)
		if err != nil {
			return "", err
		}
		_, err = svg.Build(as.Services.Projects.GetFilesystem(prj), ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "SVG [" + x.Key + "]"
		ps.Data = x
		return render(rc, as, &vsvg.View{Project: prj, SVG: x}, ps, "projects", prj.Key, "SVG||/svg/"+prj.Key, x.Key)
	})
}

func SVGSetApp(rc *fasthttp.RequestCtx) {
	act("svg.setapp", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prj, fs, key, err := prjAndIcon(rc, as)
		if err != nil {
			return "", err
		}
		content, err := svg.Content(fs, key)
		if err != nil {
			return "", errors.Wrap(err, "unable to read SVG ["+key+"]")
		}
		prj.Icon = key
		err = as.Services.Projects.Save(prj)
		if err != nil {
			return "", errors.Wrap(err, "unable to set project icon ["+key+"]")
		}
		err = svg.SetAppIcon(ps.Context, prj, fs, &svg.SVG{Key: key, Markup: content}, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to set app icon to ["+key+"]")
		}
		_, err = svg.Build(fs, ps.Logger)
		if err != nil {
			return "", err
		}
		return flashAndRedir(true, "set SVG ["+key+"] as app icon", "/svg/"+prj.Key, rc, ps)
	})
}

func SVGRemove(rc *fasthttp.RequestCtx) {
	act("svg.remove", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prj, fs, key, err := prjAndIcon(rc, as)
		if err != nil {
			return "", err
		}
		if key == "app" {
			return "", errors.New("you can't remove the app icon")
		}
		err = svg.Remove(fs, key, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to remove SVG ["+key+"]")
		}
		_, err = svg.Build(fs, ps.Logger)
		if err != nil {
			return "", err
		}
		return flashAndRedir(true, "removed SVG ["+key+"]", "/svg/"+prj.Key, rc, ps)
	})
}

func prjAndIcon(rc *fasthttp.RequestCtx, as *app.State) (*project.Project, filesystem.FileLoader, string, error) {
	prj, err := getProject(rc, as)
	if err != nil {
		return nil, nil, "", err
	}
	fs := as.Services.Projects.GetFilesystem(prj)

	key, err := cutil.RCRequiredString(rc, "icon", false)
	if err != nil {
		return nil, nil, "", err
	}
	return prj, fs, key, nil
}
