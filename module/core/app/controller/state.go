package controller

import (
	"fmt"

	"github.com/valyala/fasthttp"

	"{{{ .Package }}}/app"
	"{{{ .Package }}}/app/controller/cutil"
	"{{{ .Package }}}/app/lib/theme"
	"{{{ .Package }}}/app/util"
	"{{{ .Package }}}/views/verror"
)

var (
	_currentAppState       *app.State{{{ if.HasModule "marketing" }}}
	_currentSiteState      *app.State{{{ end }}}
	defaultRootTitleAppend = util.GetEnv("app_display_name_append")
	defaultRootTitle       = func() string {
		if tmp := util.GetEnv("app_display_name"); tmp != "" {
			return tmp
		}
		return util.AppName
	}()
)

func SetAppState(a *app.State) {
	_currentAppState = a
	initApp(a)
}{{{ if.HasModule "marketing" }}}

func SetSiteState(a *app.State) {
	_currentSiteState = a
	initSite(a)
}{{{ end }}}

func handleError(key string, as *app.State, ps *cutil.PageState, rc *fasthttp.RequestCtx, err error) (string, error) {
	rc.SetStatusCode(fasthttp.StatusInternalServerError)

	ps.Logger.Errorf("error running action [%s]: %+v", key, err)

	if len(ps.Breadcrumbs) == 0 {
		bc := util.StringSplitAndTrim(string(rc.URI().Path()), "/")
		bc = append(bc, "Error")
		ps.Breadcrumbs = bc
	}

	err = clean(as, ps)
	if err != nil {
		ps.Logger.Error(err)
		msg := fmt.Sprintf("error while cleaning request: %+v", err)
		ps.Logger.Error(msg)
		_, _ = rc.WriteString(msg)
	}

	redir, err := render(rc, as, &verror.Error{Err: util.GetErrorDetail(err)}, ps)
	if err != nil {
		msg := fmt.Sprintf("error while running error handler: %+v", err)
		ps.Logger.Error(msg)
		_, _ = rc.WriteString(msg)
	}
	return redir, err
}

func clean(as *app.State, ps *cutil.PageState) error {
	if ps.Profile != nil && ps.Profile.Theme == "" {
		ps.Profile.Theme = theme.ThemeDefault.Key
	}
	if ps.RootIcon == "" {
		ps.RootIcon = defaultIcon
	}
	if ps.RootPath == "" {
		ps.RootPath = "/"
	}
	if ps.RootTitle == "" {
		ps.RootTitle = defaultRootTitle
	}
	if defaultRootTitleAppend != "" {
		ps.RootTitle += " " + defaultRootTitleAppend
	}{{{ if .HasModule "search" }}}
	if ps.SearchPath == "" {
		ps.SearchPath = defaultSearchPath
	}{{{ end }}}
	if ps.ProfilePath == "" {
		ps.ProfilePath = defaultProfilePath
	}
	if len(ps.Menu) == 0 {
		m, err := MenuFor(ps.Context, ps.Authed, ps.Admin, as)
		if err != nil {
			return err
		}
		ps.Menu = m
	}
	return nil
}
