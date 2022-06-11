package controller

import (
	"github.com/valyala/fasthttp"

	"{{{ .Package }}}/app/controller/csession"
	"{{{ .Package }}}/app/controller/cutil"
	"{{{ .Package }}}/app/lib/telemetry"
	"{{{ .Package }}}/app/lib/telemetry/httpmetrics"
	"{{{ .Package }}}/app/lib/user"
	"{{{ .Package }}}/app/util"
)

var initialIcons = {{{ if .HasModule "search" }}}[]string{"searchbox"}{{{ else }}}[]string{}{{{ end }}}

func loadPageState(rc *fasthttp.RequestCtx, key string, logger util.Logger) *cutil.PageState {
	ctx, logger := httpmetrics.ExtractHeaders(rc, logger)
	traceCtx, span, logger := telemetry.StartSpan(ctx, "http:"+key, logger)
	span.Attribute("path", string(rc.Request.URI().Path()))
	httpmetrics.InjectHTTP(rc, span)

	session, flashes, prof, accts := loadSession(rc, logger)

	isAuthed, _ := user.Check("/", accts)
	isAdmin, _ := user.Check("/admin", accts)

	return &cutil.PageState{
		Method: string(rc.Method()), URI: rc.Request.URI(), Flashes: flashes, Session: session,
		Profile: prof, Accounts: accts, Authed: isAuthed, Admin: isAdmin,
		Icons: initialIcons, Context: traceCtx, Span: span, Logger: logger,
	}
}

func loadSession(rc *fasthttp.RequestCtx, logger util.Logger) (util.ValueMap, []string, *user.Profile, user.Accounts) {
	sessionBytes := rc.Request.Header.Cookie(util.AppKey)
	session := util.ValueMap{}
	if len(sessionBytes) > 0 {
		dec, err := util.DecryptMessage(nil, string(sessionBytes), logger)
		if err != nil {
			logger.Warnf("error decrypting session: %+v", err)
		}
		err = util.FromJSON([]byte(dec), &session)
		if err != nil {
			logger.Warnf("error parsing session: %+v", err)
		}
	}

	flashes := util.StringSplitAndTrim(session.GetStringOpt(csession.WebFlashKey), ";")
	if len(flashes) > 0 {
		delete(session, csession.WebFlashKey)
		err := csession.SaveSession(rc, session, logger)
		if err != nil {
			logger.Warnf("can't save session: %+v", err)
		}
	}

	prof, err := LoadProfile(session)
	if err != nil {
		logger.Warnf("can't load profile: %+v", err)
	}

	var accts user.Accounts
	authX, ok := session[csession.WebAuthKey]
	if ok {
		authS, ok := authX.(string)
		if ok {
			accts = user.AccountsFromString(authS)
		}
	}

	return session, flashes, prof, accts
}

func LoadProfile(session util.ValueMap) (*user.Profile, error) {
	x, ok := session["profile"]
	if !ok {
		return user.DefaultProfile.Clone(), nil
	}
	s, ok := x.(string)
	if !ok {
		return user.DefaultProfile.Clone(), nil
	}
	p := &user.Profile{}
	err := util.FromJSON([]byte(s), p)
	if err != nil {
		return nil, err
	}
	return p, nil
}
