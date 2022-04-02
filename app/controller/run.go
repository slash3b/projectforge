package controller

import (
	"fmt"
	"strings"
	"sync"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/action"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/app/util"
	"projectforge.dev/projectforge/views/vaction"
)

func RunAction(rc *fasthttp.RequestCtx) {
	act("run.action", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		tgt, err := RCRequiredString(rc, "key", false)
		if err != nil {
			return "", err
		}
		actS, err := RCRequiredString(rc, "act", false)
		if err != nil {
			return "", err
		}

		cfg := util.ValueMap{}
		actT := action.TypeFromString(actS)
		prj, err := as.Services.Projects.Get(tgt)
		if err != nil {
			return "", err
		}
		cfg["path"] = prj.Path
		rc.QueryArgs().VisitAll(func(k []byte, v []byte) {
			cfg[string(k)] = string(v)
		})
		result := action.Apply(ps.Context, actionParams(tgt, actT, cfg, as, ps.Logger))
		if result.Project == nil {
			result.Project = prj
		}
		ps.Title = fmt.Sprintf("[%s] %s", actT.Title, prj.Title())
		ps.Data = result
		page := &vaction.Result{Ctx: &action.ResultContext{Prj: prj, Cfg: cfg, Res: result}}
		return render(rc, as, page, ps, "projects", prj.Key, actT.Title)
	})
}

func RunAllActions(rc *fasthttp.RequestCtx) {
	act("run.all.actions", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		actS, err := RCRequiredString(rc, "act", false)
		if err != nil {
			return "", err
		}
		cfg := util.ValueMap{}
		rc.QueryArgs().VisitAll(func(k []byte, v []byte) {
			cfg[string(k)] = string(v)
		})
		actT := action.TypeFromString(actS)
		prjs := as.Services.Projects.Projects()

		var results action.ResultContexts
		var mutex sync.Mutex
		var wg sync.WaitGroup
		wg.Add(len(prjs))
		for _, prj := range prjs {
			p := prj
			go func() {
				c := cfg.Clone()
				c["path"] = p.Path
				result := action.Apply(ps.Context, actionParams(p.Key, actT, c, as, ps.Logger))
				if result.Project == nil {
					result.Project = p
				}
				rc := &action.ResultContext{Prj: p, Cfg: c, Res: result}
				mutex.Lock()
				results = append(results, rc)
				wg.Done()
				mutex.Unlock()
			}()
		}
		wg.Wait()
		slices.SortFunc(results, func(l *action.ResultContext, r *action.ResultContext) bool {
			return strings.ToLower(l.Prj.Title()) < strings.ToLower(r.Prj.Title())
		})
		ps.Title = fmt.Sprintf("[%s] All Projects", actT.Title)
		ps.Data = results
		page := &vaction.Results{T: actT, Ctxs: results}
		return render(rc, as, page, ps, "projects", actT.Title)
	})
}

func actionParams(tgt string, t action.Type, cfg util.ValueMap, as *app.State, logger *zap.SugaredLogger) *action.Params {
	return &action.Params{
		ProjectKey: tgt, T: t, Cfg: cfg,
		MSvc: as.Services.Modules, PSvc: as.Services.Projects, ESvc: as.Services.Export, Logger: logger,
	}
}
