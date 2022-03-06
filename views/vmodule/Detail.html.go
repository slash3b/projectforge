// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vmodule/Detail.html:1
package vmodule

//line views/vmodule/Detail.html:1
import (
	"projectforge.dev/app"
	"projectforge.dev/app/controller/cutil"
	"projectforge.dev/app/module"
	"projectforge.dev/app/project"
	"projectforge.dev/views/layout"
)

//line views/vmodule/Detail.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vmodule/Detail.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vmodule/Detail.html:9
type Detail struct {
	layout.Basic
	Module *module.Module
	Usages project.Projects
}

//line views/vmodule/Detail.html:15
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vmodule/Detail.html:15
	qw422016.N().S(`
`)
//line views/vmodule/Detail.html:16
	mod := p.Module

//line views/vmodule/Detail.html:16
	qw422016.N().S(`  `)
//line views/vmodule/Detail.html:17
	StreamSummary(qw422016, mod, nil, ps)
//line views/vmodule/Detail.html:17
	qw422016.N().S(`
  <div class="card">
    <h3>Details</h3>
    <table class="mt">
      <tbody>
        <tr>
          <th>Key</th>
          <td>`)
//line views/vmodule/Detail.html:24
	qw422016.E().S(mod.Key)
//line views/vmodule/Detail.html:24
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th>Name</th>
          <td>`)
//line views/vmodule/Detail.html:28
	qw422016.E().S(mod.Name)
//line views/vmodule/Detail.html:28
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th>Description</th>
          <td>`)
//line views/vmodule/Detail.html:32
	qw422016.E().S(mod.Description)
//line views/vmodule/Detail.html:32
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th>Author</th>
          <td><a href="mailto:`)
//line views/vmodule/Detail.html:36
	qw422016.E().S(mod.AuthorEmail)
//line views/vmodule/Detail.html:36
	qw422016.N().S(`">`)
//line views/vmodule/Detail.html:36
	qw422016.E().S(mod.AuthorName)
//line views/vmodule/Detail.html:36
	qw422016.N().S(`</a></td>
        </tr>
        <tr>
          <th>License</th>
          <td>`)
//line views/vmodule/Detail.html:40
	qw422016.E().S(mod.License)
//line views/vmodule/Detail.html:40
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th>Sourcecode</th>
          <td><a href="`)
//line views/vmodule/Detail.html:44
	qw422016.E().S(mod.Sourcecode)
//line views/vmodule/Detail.html:44
	qw422016.N().S(`" target="_blank">`)
//line views/vmodule/Detail.html:44
	qw422016.E().S(mod.Sourcecode)
//line views/vmodule/Detail.html:44
	qw422016.N().S(`</a></td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vmodule/Detail.html:49
	if p.Module.UsageHTML() != "" {
//line views/vmodule/Detail.html:49
		qw422016.N().S(`  <div class="card">
    `)
//line views/vmodule/Detail.html:51
		qw422016.N().S(p.Module.UsageHTML())
//line views/vmodule/Detail.html:51
		qw422016.N().S(`
  </div>
`)
//line views/vmodule/Detail.html:53
	}
//line views/vmodule/Detail.html:53
	qw422016.N().S(`  <div class="card">
    <h3>Project Usages</h3>
    <ul>
`)
//line views/vmodule/Detail.html:57
	if len(p.Usages) == 0 {
//line views/vmodule/Detail.html:57
		qw422016.N().S(`      <li><em>not referenced</em></li>
`)
//line views/vmodule/Detail.html:59
	}
//line views/vmodule/Detail.html:60
	for _, x := range p.Usages {
//line views/vmodule/Detail.html:60
		qw422016.N().S(`      <li><a href="/p/`)
//line views/vmodule/Detail.html:61
		qw422016.E().S(x.Key)
//line views/vmodule/Detail.html:61
		qw422016.N().S(`">`)
//line views/vmodule/Detail.html:61
		qw422016.E().S(x.Title())
//line views/vmodule/Detail.html:61
		qw422016.N().S(`</a></li>
`)
//line views/vmodule/Detail.html:62
	}
//line views/vmodule/Detail.html:62
	qw422016.N().S(`    </ul>
  </div>
`)
//line views/vmodule/Detail.html:65
}

//line views/vmodule/Detail.html:65
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vmodule/Detail.html:65
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vmodule/Detail.html:65
	p.StreamBody(qw422016, as, ps)
//line views/vmodule/Detail.html:65
	qt422016.ReleaseWriter(qw422016)
//line views/vmodule/Detail.html:65
}

//line views/vmodule/Detail.html:65
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vmodule/Detail.html:65
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vmodule/Detail.html:65
	p.WriteBody(qb422016, as, ps)
//line views/vmodule/Detail.html:65
	qs422016 := string(qb422016.B)
//line views/vmodule/Detail.html:65
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vmodule/Detail.html:65
	return qs422016
//line views/vmodule/Detail.html:65
}
