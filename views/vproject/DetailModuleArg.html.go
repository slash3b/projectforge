// Code generated by qtc from "DetailModuleArg.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vproject/DetailModuleArg.html:1
package vproject

//line views/vproject/DetailModuleArg.html:1
import (
	"projectforge.dev/app"
	"projectforge.dev/app/controller/cutil"
	"projectforge.dev/app/export/model"
	"projectforge.dev/app/module"
	"projectforge.dev/app/util"
	"projectforge.dev/views/components"
)

//line views/vproject/DetailModuleArg.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vproject/DetailModuleArg.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vproject/DetailModuleArg.html:10
func StreamDetailModuleArg(qw422016 *qt422016.Writer, mod *module.Module, arg interface{}, as *app.State, ps *cutil.PageState) {
//line views/vproject/DetailModuleArg.html:10
	qw422016.N().S(`
`)
//line views/vproject/DetailModuleArg.html:11
	switch mod.Key {
//line views/vproject/DetailModuleArg.html:12
	case "export":
//line views/vproject/DetailModuleArg.html:14
		ea := &model.Args{}
		err := util.CycleJSON(arg, ea)
		if err != nil {
			panic(err)
		}

//line views/vproject/DetailModuleArg.html:20
		if ea.Config != nil {
//line views/vproject/DetailModuleArg.html:20
			qw422016.N().S(`    <h3>Config Settings</h3>
    `)
//line views/vproject/DetailModuleArg.html:22
			components.StreamJSON(qw422016, ea.Config)
//line views/vproject/DetailModuleArg.html:22
			qw422016.N().S(`
`)
//line views/vproject/DetailModuleArg.html:23
		}
//line views/vproject/DetailModuleArg.html:23
		qw422016.N().S(`    <h3 class="mt">Models</h3>
    <ul class="accordion">
`)
//line views/vproject/DetailModuleArg.html:26
		for _, m := range ea.Models {
//line views/vproject/DetailModuleArg.html:26
			qw422016.N().S(`      <li>
        <input id="export-accordion-`)
//line views/vproject/DetailModuleArg.html:28
			qw422016.E().S(m.Name)
//line views/vproject/DetailModuleArg.html:28
			qw422016.N().S(`" type="checkbox" hidden />
        <label for="export-accordion-`)
//line views/vproject/DetailModuleArg.html:29
			qw422016.E().S(m.Name)
//line views/vproject/DetailModuleArg.html:29
			qw422016.N().S(`">
          `)
//line views/vproject/DetailModuleArg.html:30
			components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vproject/DetailModuleArg.html:30
			qw422016.N().S(`
          `)
//line views/vproject/DetailModuleArg.html:31
			components.StreamSVGRef(qw422016, m.IconSafe(), 15, 15, ``, ps)
//line views/vproject/DetailModuleArg.html:31
			qw422016.N().S(`
          `)
//line views/vproject/DetailModuleArg.html:32
			qw422016.E().S(m.Title())
//line views/vproject/DetailModuleArg.html:32
			qw422016.N().S(`
        </label>
        <div class="bd">
          `)
//line views/vproject/DetailModuleArg.html:35
			StreamDetailExportModel(qw422016, m, as, ps)
//line views/vproject/DetailModuleArg.html:35
			qw422016.N().S(`
        </div>
      </li>
`)
//line views/vproject/DetailModuleArg.html:38
		}
//line views/vproject/DetailModuleArg.html:38
		qw422016.N().S(`    </ul>
`)
//line views/vproject/DetailModuleArg.html:40
	default:
//line views/vproject/DetailModuleArg.html:40
		qw422016.N().S(`    `)
//line views/vproject/DetailModuleArg.html:41
		components.StreamJSON(qw422016, arg)
//line views/vproject/DetailModuleArg.html:41
		qw422016.N().S(`
`)
//line views/vproject/DetailModuleArg.html:42
	}
//line views/vproject/DetailModuleArg.html:43
}

//line views/vproject/DetailModuleArg.html:43
func WriteDetailModuleArg(qq422016 qtio422016.Writer, mod *module.Module, arg interface{}, as *app.State, ps *cutil.PageState) {
//line views/vproject/DetailModuleArg.html:43
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vproject/DetailModuleArg.html:43
	StreamDetailModuleArg(qw422016, mod, arg, as, ps)
//line views/vproject/DetailModuleArg.html:43
	qt422016.ReleaseWriter(qw422016)
//line views/vproject/DetailModuleArg.html:43
}

//line views/vproject/DetailModuleArg.html:43
func DetailModuleArg(mod *module.Module, arg interface{}, as *app.State, ps *cutil.PageState) string {
//line views/vproject/DetailModuleArg.html:43
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vproject/DetailModuleArg.html:43
	WriteDetailModuleArg(qb422016, mod, arg, as, ps)
//line views/vproject/DetailModuleArg.html:43
	qs422016 := string(qb422016.B)
//line views/vproject/DetailModuleArg.html:43
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vproject/DetailModuleArg.html:43
	return qs422016
//line views/vproject/DetailModuleArg.html:43
}
