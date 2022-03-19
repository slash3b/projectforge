// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vmodule/Table.html:1
package vmodule

//line views/vmodule/Table.html:1
import (
	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/app/module"
	"projectforge.dev/projectforge/app/util"
	"projectforge.dev/projectforge/views/components"
)

//line views/vmodule/Table.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vmodule/Table.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vmodule/Table.html:9
func StreamTable(qw422016 *qt422016.Writer, mods module.Modules, as *app.State, ps *cutil.PageState) {
//line views/vmodule/Table.html:9
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vmodule/Table.html:11
	qw422016.E().S(util.StringPlural(len(mods), "Available Module"))
//line views/vmodule/Table.html:11
	qw422016.N().S(`</h3>
    <table class="mt min-200">
      <tbody>
`)
//line views/vmodule/Table.html:14
	for _, mod := range mods {
//line views/vmodule/Table.html:14
		qw422016.N().S(`        <tr>
          <td class="shrink"><a href="/m/`)
//line views/vmodule/Table.html:16
		qw422016.E().S(mod.Key)
//line views/vmodule/Table.html:16
		qw422016.N().S(`">`)
//line views/vmodule/Table.html:16
		components.StreamSVGRef(qw422016, mod.IconSafe(), 16, 16, "icon", ps)
//line views/vmodule/Table.html:16
		qw422016.N().S(` `)
//line views/vmodule/Table.html:16
		qw422016.E().S(mod.Title())
//line views/vmodule/Table.html:16
		qw422016.N().S(`</a></td>
          <td>`)
//line views/vmodule/Table.html:17
		qw422016.E().S(mod.Description)
//line views/vmodule/Table.html:17
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vmodule/Table.html:19
	}
//line views/vmodule/Table.html:19
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vmodule/Table.html:23
}

//line views/vmodule/Table.html:23
func WriteTable(qq422016 qtio422016.Writer, mods module.Modules, as *app.State, ps *cutil.PageState) {
//line views/vmodule/Table.html:23
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vmodule/Table.html:23
	StreamTable(qw422016, mods, as, ps)
//line views/vmodule/Table.html:23
	qt422016.ReleaseWriter(qw422016)
//line views/vmodule/Table.html:23
}

//line views/vmodule/Table.html:23
func Table(mods module.Modules, as *app.State, ps *cutil.PageState) string {
//line views/vmodule/Table.html:23
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vmodule/Table.html:23
	WriteTable(qb422016, mods, as, ps)
//line views/vmodule/Table.html:23
	qs422016 := string(qb422016.B)
//line views/vmodule/Table.html:23
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vmodule/Table.html:23
	return qs422016
//line views/vmodule/Table.html:23
}
