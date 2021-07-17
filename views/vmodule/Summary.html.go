// Code generated by qtc from "Summary.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vmodule/Summary.html:1
package vmodule

//line views/vmodule/Summary.html:1
import (
	"github.com/kyleu/projectforge/app/module"
	"github.com/kyleu/projectforge/views/components"
)

//line views/vmodule/Summary.html:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vmodule/Summary.html:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vmodule/Summary.html:6
func StreamSummary(qw422016 *qt422016.Writer, mod *module.Module) {
//line views/vmodule/Summary.html:6
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="#modal-module"><button type="button">JSON</button></a></div>
    <h3>`)
//line views/vmodule/Summary.html:9
	qw422016.E().S(mod.Title())
//line views/vmodule/Summary.html:9
	qw422016.N().S(`</h3>
  </div>
  `)
//line views/vmodule/Summary.html:11
	components.StreamJSONModal(qw422016, "module", "Module JSON", mod, 1)
//line views/vmodule/Summary.html:11
	qw422016.N().S(`
`)
//line views/vmodule/Summary.html:12
}

//line views/vmodule/Summary.html:12
func WriteSummary(qq422016 qtio422016.Writer, mod *module.Module) {
//line views/vmodule/Summary.html:12
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vmodule/Summary.html:12
	StreamSummary(qw422016, mod)
//line views/vmodule/Summary.html:12
	qt422016.ReleaseWriter(qw422016)
//line views/vmodule/Summary.html:12
}

//line views/vmodule/Summary.html:12
func Summary(mod *module.Module) string {
//line views/vmodule/Summary.html:12
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vmodule/Summary.html:12
	WriteSummary(qb422016, mod)
//line views/vmodule/Summary.html:12
	qs422016 := string(qb422016.B)
//line views/vmodule/Summary.html:12
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vmodule/Summary.html:12
	return qs422016
//line views/vmodule/Summary.html:12
}
