// Code generated by qtc from "Packages.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vbuild/Packages.html:1
package vbuild

//line views/vbuild/Packages.html:1
import (
	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/app/project"
	"projectforge.dev/projectforge/app/project/action"
	"projectforge.dev/projectforge/app/project/build"
	"projectforge.dev/projectforge/views/components"
	"projectforge.dev/projectforge/views/layout"
	"projectforge.dev/projectforge/views/vproject"
)

//line views/vbuild/Packages.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vbuild/Packages.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vbuild/Packages.html:12
type Packages struct {
	layout.Basic
	Project     *project.Project
	BuildResult *action.Result
	Packages    build.Pkgs
}

//line views/vbuild/Packages.html:19
func (p *Packages) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vbuild/Packages.html:19
	qw422016.N().S(`
`)
//line views/vbuild/Packages.html:20
	prj := p.Project

//line views/vbuild/Packages.html:20
	qw422016.N().S(`  `)
//line views/vbuild/Packages.html:21
	vproject.StreamSummary(qw422016, prj, "Packages", nil, nil, &action.TypeBuild, ps)
//line views/vbuild/Packages.html:21
	qw422016.N().S(`
  `)
//line views/vbuild/Packages.html:22
	StreamBuildOptions(qw422016, prj.Key)
//line views/vbuild/Packages.html:22
	qw422016.N().S(`

`)
//line views/vbuild/Packages.html:24
	if p.BuildResult != nil && len(p.BuildResult.Errors) > 0 {
//line views/vbuild/Packages.html:24
		qw422016.N().S(`  <div class="card">
    <h3>Error</h3>
`)
//line views/vbuild/Packages.html:27
		for _, e := range p.BuildResult.Errors {
//line views/vbuild/Packages.html:27
			qw422016.N().S(`    <p class="error">`)
//line views/vbuild/Packages.html:28
			qw422016.E().S(e)
//line views/vbuild/Packages.html:28
			qw422016.N().S(`</p>
`)
//line views/vbuild/Packages.html:29
		}
//line views/vbuild/Packages.html:29
		qw422016.N().S(`  </div>
`)
//line views/vbuild/Packages.html:31
	}
//line views/vbuild/Packages.html:31
	qw422016.N().S(`
  <div class="card">
    <h3>Packages</h3>
    `)
//line views/vbuild/Packages.html:35
	streamrenderPackages(qw422016, p.Project, p.Packages)
//line views/vbuild/Packages.html:35
	qw422016.N().S(`
  </div>
`)
//line views/vbuild/Packages.html:37
}

//line views/vbuild/Packages.html:37
func (p *Packages) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vbuild/Packages.html:37
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vbuild/Packages.html:37
	p.StreamBody(qw422016, as, ps)
//line views/vbuild/Packages.html:37
	qt422016.ReleaseWriter(qw422016)
//line views/vbuild/Packages.html:37
}

//line views/vbuild/Packages.html:37
func (p *Packages) Body(as *app.State, ps *cutil.PageState) string {
//line views/vbuild/Packages.html:37
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vbuild/Packages.html:37
	p.WriteBody(qb422016, as, ps)
//line views/vbuild/Packages.html:37
	qs422016 := string(qb422016.B)
//line views/vbuild/Packages.html:37
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vbuild/Packages.html:37
	return qs422016
//line views/vbuild/Packages.html:37
}

//line views/vbuild/Packages.html:39
type PackagesAll struct {
	layout.Basic
	Projects project.Projects
	Results  map[string]*action.Result
	Packages map[string]build.Pkgs
	Tags     []string
}

//line views/vbuild/Packages.html:47
func (p *PackagesAll) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vbuild/Packages.html:47
	qw422016.N().S(`
  <div class="card">
    <h3>Package Graphs</h3>
    <div class="mt">
      <ul class="accordion">
`)
//line views/vbuild/Packages.html:52
	for _, prj := range p.Projects {
//line views/vbuild/Packages.html:54
		res := p.Results[prj.Key]
		pkgs := p.Packages[prj.Key]

//line views/vbuild/Packages.html:56
		qw422016.N().S(`        <li>
          <input id="accordion-`)
//line views/vbuild/Packages.html:58
		qw422016.E().S(prj.Key)
//line views/vbuild/Packages.html:58
		qw422016.N().S(`" type="checkbox" hidden />
          <label for="accordion-`)
//line views/vbuild/Packages.html:59
		qw422016.E().S(prj.Key)
//line views/vbuild/Packages.html:59
		qw422016.N().S(`">
            `)
//line views/vbuild/Packages.html:60
		components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vbuild/Packages.html:60
		qw422016.N().S(` `)
//line views/vbuild/Packages.html:60
		components.StreamSVGRef(qw422016, prj.IconSafe(), 16, 16, "icon", ps)
//line views/vbuild/Packages.html:60
		qw422016.N().S(` `)
//line views/vbuild/Packages.html:60
		qw422016.E().S(prj.Title())
//line views/vbuild/Packages.html:60
		qw422016.N().S(`
          </label>
          <div class="bd">
            `)
//line views/vbuild/Packages.html:63
		vproject.StreamSummary(qw422016, prj, "Packages", nil, nil, nil, ps)
//line views/vbuild/Packages.html:63
		qw422016.N().S(`
`)
//line views/vbuild/Packages.html:64
		if res != nil && len(res.Errors) > 0 {
//line views/vbuild/Packages.html:64
			qw422016.N().S(`            <div class="card">
              <h3>Error</h3>
`)
//line views/vbuild/Packages.html:67
			for _, e := range res.Errors {
//line views/vbuild/Packages.html:67
				qw422016.N().S(`              <p class="error">`)
//line views/vbuild/Packages.html:68
				qw422016.E().S(e)
//line views/vbuild/Packages.html:68
				qw422016.N().S(`</p>
`)
//line views/vbuild/Packages.html:69
			}
//line views/vbuild/Packages.html:69
			qw422016.N().S(`            </div>
`)
//line views/vbuild/Packages.html:71
		}
//line views/vbuild/Packages.html:71
		qw422016.N().S(`            <div class="card">
              <h3>Package Graph</h3>
              `)
//line views/vbuild/Packages.html:74
		streamrenderPackages(qw422016, prj, pkgs)
//line views/vbuild/Packages.html:74
		qw422016.N().S(`
            </div>
          </div>
        </li>
`)
//line views/vbuild/Packages.html:78
	}
//line views/vbuild/Packages.html:78
	qw422016.N().S(`      </ul>
    </div>
  </div>
`)
//line views/vbuild/Packages.html:82
}

//line views/vbuild/Packages.html:82
func (p *PackagesAll) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vbuild/Packages.html:82
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vbuild/Packages.html:82
	p.StreamBody(qw422016, as, ps)
//line views/vbuild/Packages.html:82
	qt422016.ReleaseWriter(qw422016)
//line views/vbuild/Packages.html:82
}

//line views/vbuild/Packages.html:82
func (p *PackagesAll) Body(as *app.State, ps *cutil.PageState) string {
//line views/vbuild/Packages.html:82
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vbuild/Packages.html:82
	p.WriteBody(qb422016, as, ps)
//line views/vbuild/Packages.html:82
	qs422016 := string(qb422016.B)
//line views/vbuild/Packages.html:82
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vbuild/Packages.html:82
	return qs422016
//line views/vbuild/Packages.html:82
}

//line views/vbuild/Packages.html:84
func streamrenderPackages(qw422016 *qt422016.Writer, prj *project.Project, p build.Pkgs) {
//line views/vbuild/Packages.html:84
	qw422016.N().S(`
  `)
//line views/vbuild/Packages.html:85
	components.StreamJSON(qw422016, p)
//line views/vbuild/Packages.html:85
	qw422016.N().S(`
  <pre>`)
//line views/vbuild/Packages.html:86
	qw422016.E().S(p.ToGraph(prj.Package))
//line views/vbuild/Packages.html:86
	qw422016.N().S(`</pre>
`)
//line views/vbuild/Packages.html:87
}

//line views/vbuild/Packages.html:87
func writerenderPackages(qq422016 qtio422016.Writer, prj *project.Project, p build.Pkgs) {
//line views/vbuild/Packages.html:87
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vbuild/Packages.html:87
	streamrenderPackages(qw422016, prj, p)
//line views/vbuild/Packages.html:87
	qt422016.ReleaseWriter(qw422016)
//line views/vbuild/Packages.html:87
}

//line views/vbuild/Packages.html:87
func renderPackages(prj *project.Project, p build.Pkgs) string {
//line views/vbuild/Packages.html:87
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vbuild/Packages.html:87
	writerenderPackages(qb422016, prj, p)
//line views/vbuild/Packages.html:87
	qs422016 := string(qb422016.B)
//line views/vbuild/Packages.html:87
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vbuild/Packages.html:87
	return qs422016
//line views/vbuild/Packages.html:87
}
