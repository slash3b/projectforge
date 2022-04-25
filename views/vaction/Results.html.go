// Code generated by qtc from "Results.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vaction/Results.html:1
package vaction

//line views/vaction/Results.html:1
import (
	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/action"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/views/components"
	"projectforge.dev/projectforge/views/layout"
	"projectforge.dev/projectforge/views/vproject"
	"projectforge.dev/projectforge/views/vsearch"
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
	StreamSharedActions(qw422016, p.T.Title, as, ps)
//line views/vaction/Results.html:19
	qw422016.N().S(`
    <div class="mt">
      <ul class="accordion">
`)
//line views/vaction/Results.html:22
	for _, x := range p.Ctxs {
//line views/vaction/Results.html:22
		qw422016.N().S(`        <li>
          <input id="accordion-`)
//line views/vaction/Results.html:24
		qw422016.E().S(x.Prj.Key)
//line views/vaction/Results.html:24
		qw422016.N().S(`" type="checkbox" hidden />
          <label for="accordion-`)
//line views/vaction/Results.html:25
		qw422016.E().S(x.Prj.Key)
//line views/vaction/Results.html:25
		qw422016.N().S(`">
            <div class="right">`)
//line views/vaction/Results.html:26
		qw422016.N().S(x.Status())
//line views/vaction/Results.html:26
		if x.Res != nil && len(x.Res.Errors) > 0 {
//line views/vaction/Results.html:26
			qw422016.N().S(` (`)
//line views/vaction/Results.html:26
			qw422016.N().D(len(x.Res.Errors))
//line views/vaction/Results.html:26
			qw422016.N().S(` errors)`)
//line views/vaction/Results.html:26
		}
//line views/vaction/Results.html:26
		qw422016.N().S(`</div>
            `)
//line views/vaction/Results.html:27
		components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vaction/Results.html:27
		qw422016.N().S(` `)
//line views/vaction/Results.html:27
		components.StreamSVGRef(qw422016, x.Prj.IconSafe(), 16, 16, "icon", ps)
//line views/vaction/Results.html:27
		qw422016.N().S(` `)
//line views/vaction/Results.html:27
		qw422016.E().S(x.Prj.Title())
//line views/vaction/Results.html:27
		qw422016.N().S(`
          </label>
          <div class="bd">
            `)
//line views/vaction/Results.html:30
		vproject.StreamSummary(qw422016, x.Prj, nil, nil, &x.Res.Action, ps)
//line views/vaction/Results.html:30
		qw422016.N().S(`
            `)
//line views/vaction/Results.html:31
		StreamDetail(qw422016, x.Cfg, x.Res, false, as, ps)
//line views/vaction/Results.html:31
		qw422016.N().S(`
          </div>
        </li>
`)
//line views/vaction/Results.html:34
	}
//line views/vaction/Results.html:34
	qw422016.N().S(`      </ul>
    </div>
  </div>
`)
//line views/vaction/Results.html:38
}

//line views/vaction/Results.html:38
func (p *Results) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vaction/Results.html:38
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vaction/Results.html:38
	p.StreamBody(qw422016, as, ps)
//line views/vaction/Results.html:38
	qt422016.ReleaseWriter(qw422016)
//line views/vaction/Results.html:38
}

//line views/vaction/Results.html:38
func (p *Results) Body(as *app.State, ps *cutil.PageState) string {
//line views/vaction/Results.html:38
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vaction/Results.html:38
	p.WriteBody(qb422016, as, ps)
//line views/vaction/Results.html:38
	qs422016 := string(qb422016.B)
//line views/vaction/Results.html:38
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vaction/Results.html:38
	return qs422016
//line views/vaction/Results.html:38
}

//line views/vaction/Results.html:40
func StreamSharedActions(qw422016 *qt422016.Writer, title string, as *app.State, ps *cutil.PageState) {
//line views/vaction/Results.html:40
	qw422016.N().S(`
  `)
//line views/vaction/Results.html:41
	vsearch.StreamForm(qw422016, "/p/search", "", "Search Files", ps)
//line views/vaction/Results.html:41
	qw422016.N().S(`
  <h3>All Projects: `)
//line views/vaction/Results.html:42
	qw422016.E().S(title)
//line views/vaction/Results.html:42
	qw422016.N().S(`</h3>
  <div class="mt">
`)
//line views/vaction/Results.html:44
	for _, t := range action.ProjectTypes {
//line views/vaction/Results.html:44
		qw422016.N().S(`    <a href="/run/`)
//line views/vaction/Results.html:45
		qw422016.E().S(t.Key)
//line views/vaction/Results.html:45
		qw422016.N().S(`" title="`)
//line views/vaction/Results.html:45
		qw422016.E().S(t.Description)
//line views/vaction/Results.html:45
		qw422016.N().S(`"><button>`)
//line views/vaction/Results.html:45
		qw422016.E().S(t.Title)
//line views/vaction/Results.html:45
		qw422016.N().S(`</button></a>
`)
//line views/vaction/Results.html:46
	}
//line views/vaction/Results.html:46
	qw422016.N().S(`    <a href="/run/deps" title="Manage the dependencies of your project"><button>Dependencies</button></a>
    <a href="/git" title="Git dashboard for all projects"><button>Git</button></a>
  </div>
`)
//line views/vaction/Results.html:50
}

//line views/vaction/Results.html:50
func WriteSharedActions(qq422016 qtio422016.Writer, title string, as *app.State, ps *cutil.PageState) {
//line views/vaction/Results.html:50
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vaction/Results.html:50
	StreamSharedActions(qw422016, title, as, ps)
//line views/vaction/Results.html:50
	qt422016.ReleaseWriter(qw422016)
//line views/vaction/Results.html:50
}

//line views/vaction/Results.html:50
func SharedActions(title string, as *app.State, ps *cutil.PageState) string {
//line views/vaction/Results.html:50
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vaction/Results.html:50
	WriteSharedActions(qb422016, title, as, ps)
//line views/vaction/Results.html:50
	qs422016 := string(qb422016.B)
//line views/vaction/Results.html:50
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vaction/Results.html:50
	return qs422016
//line views/vaction/Results.html:50
}
