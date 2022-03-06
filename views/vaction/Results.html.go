// Code generated by qtc from "Results.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vaction/Results.html:1
package vaction

//line views/vaction/Results.html:1
import (
	"projectforge.dev/app"
	"projectforge.dev/app/action"
	"projectforge.dev/app/controller/cutil"
	"projectforge.dev/views/components"
	"projectforge.dev/views/layout"
	"projectforge.dev/views/vproject"
	"projectforge.dev/views/vsearch"
)

//line views/vaction/Results.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vaction/Results.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vaction/Results.html:11
type Results struct {
	layout.Basic
	T    action.Type
	Ctxs action.ResultContexts
}

//line views/vaction/Results.html:17
func (p *Results) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vaction/Results.html:17
	qw422016.N().S(`
  <div class="card">
    `)
//line views/vaction/Results.html:19
	vsearch.StreamForm(qw422016, "/p/search", "", "Search Files", ps)
//line views/vaction/Results.html:19
	qw422016.N().S(`
    <h3>All Projects: `)
//line views/vaction/Results.html:20
	qw422016.E().S(p.T.Title)
//line views/vaction/Results.html:20
	qw422016.N().S(`</h3>
    <div class="mt">
`)
//line views/vaction/Results.html:22
	for _, t := range action.ProjectTypes {
//line views/vaction/Results.html:22
		qw422016.N().S(`      <a href="/run/`)
//line views/vaction/Results.html:23
		qw422016.E().S(t.Key)
//line views/vaction/Results.html:23
		qw422016.N().S(`" title="`)
//line views/vaction/Results.html:23
		qw422016.E().S(t.Description)
//line views/vaction/Results.html:23
		qw422016.N().S(`"><button>`)
//line views/vaction/Results.html:23
		qw422016.E().S(t.Title)
//line views/vaction/Results.html:23
		qw422016.N().S(`</button></a>
`)
//line views/vaction/Results.html:24
	}
//line views/vaction/Results.html:24
	qw422016.N().S(`      <a href="/git" title="Git dashboard for all projects"><button>Git</button></a>
    </div>
    <div class="mt">
      <ul class="accordion">
`)
//line views/vaction/Results.html:29
	for _, x := range p.Ctxs {
//line views/vaction/Results.html:29
		qw422016.N().S(`        <li>
          <input id="accordion-`)
//line views/vaction/Results.html:31
		qw422016.E().S(x.Prj.Key)
//line views/vaction/Results.html:31
		qw422016.N().S(`" type="checkbox" hidden />
          <label for="accordion-`)
//line views/vaction/Results.html:32
		qw422016.E().S(x.Prj.Key)
//line views/vaction/Results.html:32
		qw422016.N().S(`">
            <div class="right">`)
//line views/vaction/Results.html:33
		qw422016.N().S(x.Status())
//line views/vaction/Results.html:33
		if x.Res != nil && len(x.Res.Errors) > 0 {
//line views/vaction/Results.html:33
			qw422016.N().S(` (`)
//line views/vaction/Results.html:33
			qw422016.N().D(len(x.Res.Errors))
//line views/vaction/Results.html:33
			qw422016.N().S(` errors)`)
//line views/vaction/Results.html:33
		}
//line views/vaction/Results.html:33
		qw422016.N().S(`</div>
            `)
//line views/vaction/Results.html:34
		components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vaction/Results.html:34
		qw422016.N().S(` `)
//line views/vaction/Results.html:34
		components.StreamSVGRef(qw422016, x.Prj.IconSafe(), 16, 16, "icon", ps)
//line views/vaction/Results.html:34
		qw422016.N().S(` `)
//line views/vaction/Results.html:34
		qw422016.E().S(x.Prj.Title())
//line views/vaction/Results.html:34
		qw422016.N().S(`
          </label>
          <div class="bd">
            `)
//line views/vaction/Results.html:37
		vproject.StreamSummary(qw422016, x.Prj, nil, nil, ps)
//line views/vaction/Results.html:37
		qw422016.N().S(`
            `)
//line views/vaction/Results.html:38
		StreamDetail(qw422016, x.Cfg, x.Res, false, as, ps)
//line views/vaction/Results.html:38
		qw422016.N().S(`
          </div>
        </li>
`)
//line views/vaction/Results.html:41
	}
//line views/vaction/Results.html:41
	qw422016.N().S(`      </ul>
    </div>
  </div>
`)
//line views/vaction/Results.html:45
}

//line views/vaction/Results.html:45
func (p *Results) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vaction/Results.html:45
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vaction/Results.html:45
	p.StreamBody(qw422016, as, ps)
//line views/vaction/Results.html:45
	qt422016.ReleaseWriter(qw422016)
//line views/vaction/Results.html:45
}

//line views/vaction/Results.html:45
func (p *Results) Body(as *app.State, ps *cutil.PageState) string {
//line views/vaction/Results.html:45
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vaction/Results.html:45
	p.WriteBody(qb422016, as, ps)
//line views/vaction/Results.html:45
	qs422016 := string(qb422016.B)
//line views/vaction/Results.html:45
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vaction/Results.html:45
	return qs422016
//line views/vaction/Results.html:45
}
