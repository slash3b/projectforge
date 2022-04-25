// Code generated by qtc from "DepMap.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vbuild/DepMap.html:1
package vbuild

//line views/vbuild/DepMap.html:1
import (
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/views/components"
	"projectforge.dev/projectforge/views/layout"
	"projectforge.dev/projectforge/views/vaction"
)

//line views/vbuild/DepMap.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vbuild/DepMap.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vbuild/DepMap.html:12
type DepMap struct {
	layout.Basic
	Message string
	Result  map[string]map[string][]string
}

//line views/vbuild/DepMap.html:18
func (p *DepMap) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vbuild/DepMap.html:18
	qw422016.N().S(`
  <div class="card">
    `)
//line views/vbuild/DepMap.html:20
	vaction.StreamSharedActions(qw422016, "Dependency Conflicts", as, ps)
//line views/vbuild/DepMap.html:20
	qw422016.N().S(`
`)
//line views/vbuild/DepMap.html:21
	if p.Message != "" {
//line views/vbuild/DepMap.html:21
		qw422016.N().S(`    <h4 class="mt mb">`)
//line views/vbuild/DepMap.html:22
		qw422016.E().S(p.Message)
//line views/vbuild/DepMap.html:22
		qw422016.N().S(`</h4>
`)
//line views/vbuild/DepMap.html:23
	}
//line views/vbuild/DepMap.html:24
	if len(p.Result) == 0 {
//line views/vbuild/DepMap.html:24
		qw422016.N().S(`    <div class="mt"><em>No dependency conflicts, nice work!</em></div>
`)
//line views/vbuild/DepMap.html:26
	}
//line views/vbuild/DepMap.html:26
	qw422016.N().S(`    <ul class="accordion mt">
`)
//line views/vbuild/DepMap.html:29
	depKeys := maps.Keys(p.Result)
	slices.Sort(depKeys)

//line views/vbuild/DepMap.html:32
	for _, k := range depKeys {
//line views/vbuild/DepMap.html:33
		v := p.Result[k]

//line views/vbuild/DepMap.html:33
		qw422016.N().S(`
      <li>
        <input id="accordion-`)
//line views/vbuild/DepMap.html:35
		qw422016.E().S(k)
//line views/vbuild/DepMap.html:35
		qw422016.N().S(`" type="checkbox" hidden />
        <label for="accordion-`)
//line views/vbuild/DepMap.html:36
		qw422016.E().S(k)
//line views/vbuild/DepMap.html:36
		qw422016.N().S(`">
          <div class="right">`)
//line views/vbuild/DepMap.html:37
		qw422016.N().D(len(v))
//line views/vbuild/DepMap.html:37
		qw422016.N().S(`</div>
          `)
//line views/vbuild/DepMap.html:38
		components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vbuild/DepMap.html:38
		qw422016.N().S(` `)
//line views/vbuild/DepMap.html:38
		qw422016.E().S(k)
//line views/vbuild/DepMap.html:38
		qw422016.N().S(`
        </label>
        <div class="bd">
          <ul>
`)
//line views/vbuild/DepMap.html:42
		for vers, prjs := range v {
//line views/vbuild/DepMap.html:42
			qw422016.N().S(`            <li><a href="?key=`)
//line views/vbuild/DepMap.html:43
			qw422016.E().S(k)
//line views/vbuild/DepMap.html:43
			qw422016.N().S(`&version=`)
//line views/vbuild/DepMap.html:43
			qw422016.E().S(vers)
//line views/vbuild/DepMap.html:43
			qw422016.N().S(`">`)
//line views/vbuild/DepMap.html:43
			qw422016.E().S(vers)
//line views/vbuild/DepMap.html:43
			qw422016.N().S(`</a>:
              <ul>
`)
//line views/vbuild/DepMap.html:45
			for _, prj := range prjs {
//line views/vbuild/DepMap.html:45
				qw422016.N().S(`                <li><a href="/p/`)
//line views/vbuild/DepMap.html:46
				qw422016.E().S(prj)
//line views/vbuild/DepMap.html:46
				qw422016.N().S(`">`)
//line views/vbuild/DepMap.html:46
				qw422016.E().S(prj)
//line views/vbuild/DepMap.html:46
				qw422016.N().S(`</a></li>
`)
//line views/vbuild/DepMap.html:47
			}
//line views/vbuild/DepMap.html:47
			qw422016.N().S(`              </ul>
            </li>
`)
//line views/vbuild/DepMap.html:50
		}
//line views/vbuild/DepMap.html:50
		qw422016.N().S(`          </ul>
        </div>
      </li>
`)
//line views/vbuild/DepMap.html:54
	}
//line views/vbuild/DepMap.html:54
	qw422016.N().S(`    </ul>
  </div>
`)
//line views/vbuild/DepMap.html:57
}

//line views/vbuild/DepMap.html:57
func (p *DepMap) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vbuild/DepMap.html:57
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vbuild/DepMap.html:57
	p.StreamBody(qw422016, as, ps)
//line views/vbuild/DepMap.html:57
	qt422016.ReleaseWriter(qw422016)
//line views/vbuild/DepMap.html:57
}

//line views/vbuild/DepMap.html:57
func (p *DepMap) Body(as *app.State, ps *cutil.PageState) string {
//line views/vbuild/DepMap.html:57
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vbuild/DepMap.html:57
	p.WriteBody(qb422016, as, ps)
//line views/vbuild/DepMap.html:57
	qs422016 := string(qb422016.B)
//line views/vbuild/DepMap.html:57
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vbuild/DepMap.html:57
	return qs422016
//line views/vbuild/DepMap.html:57
}
