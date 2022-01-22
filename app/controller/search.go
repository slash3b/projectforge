// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"github.com/valyala/fasthttp"

	"github.com/kyleu/projectforge/app"
	"github.com/kyleu/projectforge/app/controller/cutil"
	"github.com/kyleu/projectforge/app/lib/search"
	"github.com/kyleu/projectforge/views/vsearch"
)

func Search(rc *fasthttp.RequestCtx) {
	act("search", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := string(rc.URI().QueryArgs().Peek("q"))
		paramSet := cutil.ParamSetFromRequest(rc)
		params := &search.Params{Q: q, PS: paramSet}
		results, errs := search.Search(ps.Context, as, params)
		ps.Title = "Search Results"
		ps.Data = results
		return render(rc, as, &vsearch.Results{Params: params, Results: results, Errors: errs}, ps, "Search")
	})
}
