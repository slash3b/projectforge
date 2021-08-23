package controller

import (
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"

	"{{{ .Package }}}/app"
	"{{{ .Package }}}/app/controller/cutil"
	{{{ if.HasModule "marketing" }}}"{{{ .Package }}}/app/site"
	{{{ end }}}"{{{ .Package }}}/app/theme"
	"{{{ .Package }}}/app/util"
	"{{{ .Package }}}/views/verror"
)

const (
	defaultSearchPath  = "/search"
	defaultProfilePath = "/profile"
	defaultIcon        = "app"
)

func act(key string, rc *fasthttp.RequestCtx, f func(as *app.State, ps *cutil.PageState) (string, error)) {
	as := _currentAppState
	ps := loadPageState(rc, key, as)
	err := initAppRequest(as, ps)
	if err != nil {
		as.Logger.Warnf("%+v", err)
	}
	actComplete(key, as, ps, rc, f)
}
{{{ if.HasModule "marketing" }}}
func actSite(key string, rc *fasthttp.RequestCtx, f func(as *app.State, ps *cutil.PageState) (string, error)) {
	as := _currentSiteState
	ps := loadPageState(rc, key, as)
	ps.Menu = site.Menu(ps.Profile, ps.Accounts)
	err := initSiteRequest(as, ps)
	if err != nil {
		as.Logger.Warnf("%+v", err)
	}
	actComplete(key, as, ps, rc, f)
}
{{{ end }}}
func actComplete(key string, as *app.State, ps *cutil.PageState, rc *fasthttp.RequestCtx, f func(as *app.State, ps *cutil.PageState) (string, error)) {
	err := clean(as, ps)
	if err != nil {
		as.Logger.Warnf("error while cleaning request, somehow: %+v", err)
	}
	status := fasthttp.StatusOK
	cutil.WriteCORS(rc)
	startNanos := time.Now().UnixNano()
	redir, err := f(as, ps)
	if err != nil {
		redir, err = handleError(key, as, ps, rc, err)
	}
	if redir != "" {
		rc.Response.Header.Set("Location", redir)
		status = fasthttp.StatusFound
		rc.SetStatusCode(status)
	}
	elapsedMillis := float64((time.Now().UnixNano()-startNanos)/int64(time.Microsecond)) / float64(1000)
	defer ps.Close()
	l := ps.Logger.With(zap.String("method", ps.Method), zap.Int("status", status), zap.Float64("elapsed", elapsedMillis))
	l.Debugf("processed request in [%.3fms] (render: %.3fms)", elapsedMillis, ps.RenderElapsed)
}

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
