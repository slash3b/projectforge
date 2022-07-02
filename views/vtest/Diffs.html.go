// Code generated by qtc from "Diffs.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vtest/Diffs.html:1
package vtest

//line views/vtest/Diffs.html:1
import (
	"fmt"

	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/app/file/diff"
	"projectforge.dev/projectforge/app/util"
	"projectforge.dev/projectforge/views/layout"
)

//line views/vtest/Diffs.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vtest/Diffs.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vtest/Diffs.html:11
type Diffs struct {
	layout.Basic
	Results diff.Results
}

//line views/vtest/Diffs.html:16
func (p *Diffs) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vtest/Diffs.html:16
	qw422016.N().S(`
  <div class="card">
    <h3>Diffs</h3>
    <table>
      <thead>
        <tr>
          <th>Filename</th>
          <th>Source</th>
          <th>Target</th>
          <th>Edits</th>
          <th>Changes</th>
          <th>Patch</th>
        </tr>
      </thead>
      <tbody>
`)
//line views/vtest/Diffs.html:31
	for _, r := range p.Results {
//line views/vtest/Diffs.html:31
		qw422016.N().S(`        <tr>
          <td>`)
//line views/vtest/Diffs.html:33
		qw422016.E().S(r.Filename)
//line views/vtest/Diffs.html:33
		qw422016.N().S(`</td>
          <td>`)
//line views/vtest/Diffs.html:34
		qw422016.E().S(r.Src)
//line views/vtest/Diffs.html:34
		qw422016.N().S(`</td>
          <td>`)
//line views/vtest/Diffs.html:35
		qw422016.E().S(r.Tgt)
//line views/vtest/Diffs.html:35
		qw422016.N().S(`</td>
          <td><pre>`)
//line views/vtest/Diffs.html:36
		qw422016.E().S(util.ToJSON(r.Edits))
//line views/vtest/Diffs.html:36
		qw422016.N().S(`</pre></td>
          <td><pre>`)
//line views/vtest/Diffs.html:37
		qw422016.E().S(util.ToJSON(r.Changes))
//line views/vtest/Diffs.html:37
		qw422016.N().S(`</pre></td>
          <td><pre>`)
//line views/vtest/Diffs.html:38
		qw422016.E().S(fmt.Sprint(r.Patch))
//line views/vtest/Diffs.html:38
		qw422016.N().S(`</pre></td>
        </tr>
`)
//line views/vtest/Diffs.html:40
	}
//line views/vtest/Diffs.html:40
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vtest/Diffs.html:44
}

//line views/vtest/Diffs.html:44
func (p *Diffs) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vtest/Diffs.html:44
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vtest/Diffs.html:44
	p.StreamBody(qw422016, as, ps)
//line views/vtest/Diffs.html:44
	qt422016.ReleaseWriter(qw422016)
//line views/vtest/Diffs.html:44
}

//line views/vtest/Diffs.html:44
func (p *Diffs) Body(as *app.State, ps *cutil.PageState) string {
//line views/vtest/Diffs.html:44
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vtest/Diffs.html:44
	p.WriteBody(qb422016, as, ps)
//line views/vtest/Diffs.html:44
	qs422016 := string(qb422016.B)
//line views/vtest/Diffs.html:44
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vtest/Diffs.html:44
	return qs422016
//line views/vtest/Diffs.html:44
}
