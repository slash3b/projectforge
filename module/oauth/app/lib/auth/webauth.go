package auth

import (
	"github.com/valyala/fasthttp"

	"{{{ .Package }}}/app/controller/csession"
	"{{{ .Package }}}/app/lib/user"
	"{{{ .Package }}}/app/util"
)

func getAuthURL(prv *Provider, rc *fasthttp.RequestCtx, websess util.ValueMap, logger util.Logger) (string, error) {
	g, err := gothFor(rc, prv)
	if err != nil {
		return "", err
	}

	sess, err := g.BeginAuth(setState(rc))
	if err != nil {
		return "", err
	}

	u, err := sess.GetAuthURL()
	if err != nil {
		return "", err
	}

	err = csession.StoreInSession(prv.ID, sess.Marshal(), rc, websess, logger)
	if err != nil {
		return "", err
	}

	return u, err
}

func getCurrentAuths(websess util.ValueMap) user.Accounts {
	authS, err := csession.GetFromSession(WebAuthKey, websess)
	var ret user.Accounts
	if err == nil && authS != "" {
		ret = user.AccountsFromString(authS)
	}
	return ret
}

func setCurrentAuths(s user.Accounts, rc *fasthttp.RequestCtx, websess util.ValueMap, logger util.Logger) error {
	s.Sort()
	if len(s) == 0 {
		return csession.RemoveFromSession(WebAuthKey, rc, websess, logger)
	}
	return csession.StoreInSession(WebAuthKey, s.String(), rc, websess, logger)
}
