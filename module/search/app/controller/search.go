package controller

import (
	"fmt"

	"github.com/valyala/fasthttp"

	"{{{ .Package }}}/app"
	"{{{ .Package }}}/app/controller/cutil"
	"{{{ .Package }}}/app/lib/search"
	"{{{ .Package }}}/views/vsearch"
)

func Search(rc *fasthttp.RequestCtx) {
	act("search", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := string(rc.URI().QueryArgs().Peek("q"))
		paramSet := cutil.ParamSetFromRequest(rc)
		params := &search.Params{Q: q, PS: paramSet}
		results, errs := search.Search(ps.Context, as, params, ps.Logger)
		ps.Title = "Search Results"
		if q != "" {
			ps.Title = fmt.Sprintf("[%s] %s", q, ps.Title)
		}
		ps.Data = results
		return render(rc, as, &vsearch.Results{Params: params, Results: results, Errors: errs}, ps, "Search")
	})
}
