// Code generated by Project Forge, see https://projectforge.dev for details.
package controller

import (
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"

	"github.com/kyleu/projectforge/app"
	"github.com/kyleu/projectforge/app/controller/cutil"
	"github.com/kyleu/projectforge/app/site"
	"github.com/kyleu/projectforge/app/theme"
	"github.com/kyleu/projectforge/app/util"
	"github.com/kyleu/projectforge/views/verror"
)

const (
	defaultSearchPath  = "/search"
	defaultProfilePath = "/profile"
	defaultIcon        = "app"
)

func act(key string, ctx *fasthttp.RequestCtx, f func(as *app.State, ps *cutil.PageState) (string, error)) {
	as := _currentAppState
	ps := loadPageState(ctx, key, as)
	err := initAppRequest(as, ps)
	if err != nil {
		as.Logger.Warnf("%+v", err)
	}
	actComplete(key, as, ps, ctx, f)
}

func actSite(key string, ctx *fasthttp.RequestCtx, f func(as *app.State, ps *cutil.PageState) (string, error)) {
	as := _currentSiteState
	ps := loadPageState(ctx, key, as)
	ps.Menu = site.Menu(ps.Profile, ps.Accounts)
	err := initSiteRequest(as, ps)
	if err != nil {
		as.Logger.Warnf("%+v", err)
	}
	actComplete(key, as, ps, ctx, f)
}

func actComplete(key string, as *app.State, ps *cutil.PageState, ctx *fasthttp.RequestCtx, f func(as *app.State, ps *cutil.PageState) (string, error)) {
	err := clean(as, ps)
	if err != nil {
		as.Logger.Warnf("error while cleaning request, somehow: %+v", err)
	}
	status := fasthttp.StatusOK
	cutil.WriteCORS(ctx)
	startNanos := time.Now().UnixNano()
	redir, err := f(as, ps)
	if err != nil {
		redir, err = handleError(key, as, ps, ctx, err)
	}
	if redir != "" {
		ctx.Response.Header.Set("Location", redir)
		status = fasthttp.StatusFound
		ctx.SetStatusCode(status)
	}
	elapsedMillis := float64((time.Now().UnixNano()-startNanos)/int64(time.Microsecond)) / float64(1000)
	defer ps.Close()
	l := ps.Logger.With(zap.String("method", ps.Method), zap.Int("status", status), zap.Float64("elapsed", elapsedMillis))
	l.Debugf("processed request in [%.3fms] (render: %.3fms)", elapsedMillis, ps.RenderElapsed)
}

func handleError(key string, as *app.State, ps *cutil.PageState, ctx *fasthttp.RequestCtx, err error) (string, error) {
	ctx.SetStatusCode(fasthttp.StatusInternalServerError)

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
		_, _ = ctx.Write([]byte(msg))
	}
	redir, err := render(ctx, as, page, ps)
	if err != nil {
		msg := fmt.Sprintf("error while running error handler: %+v", err)
		as.Logger.Error(msg)
		_, _ = ctx.Write([]byte(msg))
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
		m, err := MenuFor(as)
		if err != nil {
			return err
		}
		ps.Menu = m
	}
	return nil
}
