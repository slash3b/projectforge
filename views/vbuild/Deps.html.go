// Code generated by qtc from "Deps.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vbuild/Deps.html:1
package vbuild

//line views/vbuild/Deps.html:1
import (
	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/action"
	"projectforge.dev/projectforge/app/build"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/app/project"
	"projectforge.dev/projectforge/views/components"
	"projectforge.dev/projectforge/views/layout"
	"projectforge.dev/projectforge/views/vproject"
)

//line views/vbuild/Deps.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vbuild/Deps.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vbuild/Deps.html:12
type Deps struct {
	layout.Basic
	Project      *project.Project
	BuildResult  *action.Result
	Dependencies build.Dependencies
}

//line views/vbuild/Deps.html:19
func (p *Deps) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vbuild/Deps.html:19
	qw422016.N().S(`
`)
//line views/vbuild/Deps.html:20
	prj := p.Project

//line views/vbuild/Deps.html:20
	qw422016.N().S(`  `)
//line views/vbuild/Deps.html:21
	vproject.StreamSummary(qw422016, prj, nil, nil, &action.TypeBuild, ps)
//line views/vbuild/Deps.html:21
	qw422016.N().S(`
  `)
//line views/vbuild/Deps.html:22
	StreamBuildOptions(qw422016, prj.Key)
//line views/vbuild/Deps.html:22
	qw422016.N().S(`

`)
//line views/vbuild/Deps.html:24
	if p.BuildResult != nil && len(p.BuildResult.Errors) > 0 {
//line views/vbuild/Deps.html:24
		qw422016.N().S(`  <div class="card">
    <h3>Error</h3>
`)
//line views/vbuild/Deps.html:27
		for _, e := range p.BuildResult.Errors {
//line views/vbuild/Deps.html:27
			qw422016.N().S(`    <p class="error">`)
//line views/vbuild/Deps.html:28
			qw422016.E().S(e)
//line views/vbuild/Deps.html:28
			qw422016.N().S(`</p>
`)
//line views/vbuild/Deps.html:29
		}
//line views/vbuild/Deps.html:29
		qw422016.N().S(`  </div>
`)
//line views/vbuild/Deps.html:31
	}
//line views/vbuild/Deps.html:31
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="?phase=deps&upgrade=all" onclick="return confirm('You sure about this?')"><button>Upgrade All</button></a></div>
    <h3>Dependencies</h3>
`)
//line views/vbuild/Deps.html:36
	if len(p.Dependencies) == 0 {
//line views/vbuild/Deps.html:36
		qw422016.N().S(`      <em>no dependencies, somehow</em>
`)
//line views/vbuild/Deps.html:38
	} else {
//line views/vbuild/Deps.html:38
		qw422016.N().S(`      <table class="mt">
        <thead>
        <tr>
          <th>Dependency</th>
          <th>Refs</th>
          <th>Version</th>
          <th>Available</th>
          <th class="shrink"></th>
        </tr>
        </thead>
        <tbody>
`)
//line views/vbuild/Deps.html:50
		for _, d := range p.Dependencies {
//line views/vbuild/Deps.html:50
			qw422016.N().S(`        <tr>
          <td>
            <ul class="accordion">
              <li>
                <input id="accordion-`)
//line views/vbuild/Deps.html:55
			qw422016.E().S(d.Key)
//line views/vbuild/Deps.html:55
			qw422016.N().S(`" type="checkbox" hidden />
                <label for="accordion-`)
//line views/vbuild/Deps.html:56
			qw422016.E().S(d.Key)
//line views/vbuild/Deps.html:56
			qw422016.N().S(`">`)
//line views/vbuild/Deps.html:56
			components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vbuild/Deps.html:56
			qw422016.N().S(` `)
//line views/vbuild/Deps.html:56
			qw422016.E().S(d.Key)
//line views/vbuild/Deps.html:56
			qw422016.N().S(`</label>
                <div class="bd">
`)
//line views/vbuild/Deps.html:58
			for _, r := range d.References {
//line views/vbuild/Deps.html:58
				qw422016.N().S(`                  <div>`)
//line views/vbuild/Deps.html:59
				qw422016.E().S(r)
//line views/vbuild/Deps.html:59
				qw422016.N().S(`</div>
`)
//line views/vbuild/Deps.html:60
			}
//line views/vbuild/Deps.html:60
			qw422016.N().S(`                </div>
              </li>
            </ul>
          </td>
          <td>`)
//line views/vbuild/Deps.html:65
			qw422016.N().D(len(d.References))
//line views/vbuild/Deps.html:65
			qw422016.N().S(`</td>
          <td>`)
//line views/vbuild/Deps.html:66
			qw422016.E().S(d.Version)
//line views/vbuild/Deps.html:66
			qw422016.N().S(`</td>
          <td>`)
//line views/vbuild/Deps.html:67
			qw422016.E().S(d.Available)
//line views/vbuild/Deps.html:67
			qw422016.N().S(`</td>
          <td>
`)
//line views/vbuild/Deps.html:69
			if d.Available != "" && d.Available != "(deprecated)" && d.Available != "(retracted)" {
//line views/vbuild/Deps.html:69
				qw422016.N().S(`            <a href="?phase=deps&upgrade=`)
//line views/vbuild/Deps.html:70
				qw422016.N().U(d.Key)
//line views/vbuild/Deps.html:70
				qw422016.N().S(`&old=`)
//line views/vbuild/Deps.html:70
				qw422016.N().U(d.Version)
//line views/vbuild/Deps.html:70
				qw422016.N().S(`&new=`)
//line views/vbuild/Deps.html:70
				qw422016.N().U(d.Available)
//line views/vbuild/Deps.html:70
				qw422016.N().S(`"><button>Upgrade</button></a>
`)
//line views/vbuild/Deps.html:71
			}
//line views/vbuild/Deps.html:71
			qw422016.N().S(`          </td>
        </tr>
`)
//line views/vbuild/Deps.html:74
		}
//line views/vbuild/Deps.html:74
		qw422016.N().S(`        </tbody>
      </table>
`)
//line views/vbuild/Deps.html:77
	}
//line views/vbuild/Deps.html:77
	qw422016.N().S(`  </div>
`)
//line views/vbuild/Deps.html:79
}

//line views/vbuild/Deps.html:79
func (p *Deps) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vbuild/Deps.html:79
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vbuild/Deps.html:79
	p.StreamBody(qw422016, as, ps)
//line views/vbuild/Deps.html:79
	qt422016.ReleaseWriter(qw422016)
//line views/vbuild/Deps.html:79
}

//line views/vbuild/Deps.html:79
func (p *Deps) Body(as *app.State, ps *cutil.PageState) string {
//line views/vbuild/Deps.html:79
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vbuild/Deps.html:79
	p.WriteBody(qb422016, as, ps)
//line views/vbuild/Deps.html:79
	qs422016 := string(qb422016.B)
//line views/vbuild/Deps.html:79
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vbuild/Deps.html:79
	return qs422016
//line views/vbuild/Deps.html:79
}
