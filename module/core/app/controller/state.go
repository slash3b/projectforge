package controller

import (
	"fmt"

	"github.com/valyala/fasthttp"

	"{{{ .Package }}}/app"
	"{{{ .Package }}}/app/controller/cutil"
	"{{{ .Package }}}/app/theme"
	"{{{ .Package }}}/app/util"
	"{{{ .Package }}}/views/verror"
)

var (
	_currentAppState  *app.State
	_currentSiteState *app.State
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
		ps.Breadcrumbs = cutil.Breadcrumbs{"Error"}
	}
	errDetail := util.GetErrorDetail(err)
	page := &verror.Error{Err: errDetail}

	err = clean(as, ps)
	if err != nil {
		as.Logger.Error(err)
		msg := fmt.Sprintf("error while cleaning request: %+v", err)
		as.Logger.Error(msg)
		_, _ = rc.Write([]byte(msg))
	}
	redir, err := render(rc, as, page, ps)
	if err != nil {
		msg := fmt.Sprintf("error while running error handler: %+v", err)
		as.Logger.Error(msg)
		_, _ = rc.Write([]byte(msg))
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
		ps.RootTitle = util.AppName
	}
	if ps.SearchPath == "" {
		ps.SearchPath = defaultSearchPath
	}
	if ps.ProfilePath == "" {
		ps.ProfilePath = defaultProfilePath
	}
	if len(ps.Menu) == 0 {
		m, err := MenuFor(ps.Context, as)
		if err != nil {
			return err
		}
		ps.Menu = m
	}
	return nil
}
