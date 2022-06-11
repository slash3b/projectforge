// Code generated by qtc from "ModelDetail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vexport/ModelDetail.html:1
package vexport

//line views/vexport/ModelDetail.html:1
import (
	"strings"

	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/app/export/model"
	"projectforge.dev/projectforge/app/project"
	"projectforge.dev/projectforge/views/components"
	"projectforge.dev/projectforge/views/layout"
)

//line views/vexport/ModelDetail.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vexport/ModelDetail.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vexport/ModelDetail.html:12
type ModelDetail struct {
	layout.Basic
	Project *project.Project
	Model   *model.Model
}

//line views/vexport/ModelDetail.html:18
func (p *ModelDetail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vexport/ModelDetail.html:18
	qw422016.N().S(`
`)
//line views/vexport/ModelDetail.html:19
	m := p.Model

//line views/vexport/ModelDetail.html:19
	qw422016.N().S(`  <div class="card">
    <div class="right"><a href="/p/`)
//line views/vexport/ModelDetail.html:21
	qw422016.E().S(p.Project.Key)
//line views/vexport/ModelDetail.html:21
	qw422016.N().S(`/export/models/`)
//line views/vexport/ModelDetail.html:21
	qw422016.E().S(p.Model.Name)
//line views/vexport/ModelDetail.html:21
	qw422016.N().S(`/edit"><button>Edit</button></a></div>
    <h3>`)
//line views/vexport/ModelDetail.html:22
	components.StreamSVGRefIcon(qw422016, m.IconSafe(), ps)
//line views/vexport/ModelDetail.html:22
	qw422016.E().S(m.Name)
//line views/vexport/ModelDetail.html:22
	qw422016.N().S(`</h3>
    <em>export model</em>
    `)
//line views/vexport/ModelDetail.html:24
	streammodelSummary(qw422016, m, as, ps)
//line views/vexport/ModelDetail.html:24
	qw422016.N().S(`
  </div>
  <div class="card">
    <h3>`)
//line views/vexport/ModelDetail.html:27
	components.StreamSVGRefIcon(qw422016, `first-aid`, ps)
//line views/vexport/ModelDetail.html:27
	qw422016.N().S(` Columns</h3>
    `)
//line views/vexport/ModelDetail.html:28
	streammodelColumns(qw422016, m, as, ps)
//line views/vexport/ModelDetail.html:28
	qw422016.N().S(`
  </div>
`)
//line views/vexport/ModelDetail.html:30
	if len(m.Relations) > 0 {
//line views/vexport/ModelDetail.html:30
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vexport/ModelDetail.html:32
		components.StreamSVGRefIcon(qw422016, `social`, ps)
//line views/vexport/ModelDetail.html:32
		qw422016.N().S(` Relations</h3>
    `)
//line views/vexport/ModelDetail.html:33
		streammodelRelations(qw422016, m, as, ps)
//line views/vexport/ModelDetail.html:33
		qw422016.N().S(`
  </div>`)
//line views/vexport/ModelDetail.html:34
	}
//line views/vexport/ModelDetail.html:34
	if len(m.Indexes) > 0 {
//line views/vexport/ModelDetail.html:34
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vexport/ModelDetail.html:36
		components.StreamSVGRefIcon(qw422016, `star`, ps)
//line views/vexport/ModelDetail.html:36
		qw422016.N().S(` Indexes</h3>
    `)
//line views/vexport/ModelDetail.html:37
		streammodelIndexes(qw422016, m, as, ps)
//line views/vexport/ModelDetail.html:37
		qw422016.N().S(`
  </div>`)
//line views/vexport/ModelDetail.html:38
	}
//line views/vexport/ModelDetail.html:39
}

//line views/vexport/ModelDetail.html:39
func (p *ModelDetail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vexport/ModelDetail.html:39
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vexport/ModelDetail.html:39
	p.StreamBody(qw422016, as, ps)
//line views/vexport/ModelDetail.html:39
	qt422016.ReleaseWriter(qw422016)
//line views/vexport/ModelDetail.html:39
}

//line views/vexport/ModelDetail.html:39
func (p *ModelDetail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vexport/ModelDetail.html:39
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vexport/ModelDetail.html:39
	p.WriteBody(qb422016, as, ps)
//line views/vexport/ModelDetail.html:39
	qs422016 := string(qb422016.B)
//line views/vexport/ModelDetail.html:39
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vexport/ModelDetail.html:39
	return qs422016
//line views/vexport/ModelDetail.html:39
}

//line views/vexport/ModelDetail.html:41
func streammodelSummary(qw422016 *qt422016.Writer, model *model.Model, as *app.State, ps *cutil.PageState) {
//line views/vexport/ModelDetail.html:41
	qw422016.N().S(`
  <table class="mt min-200 full-width">
    <tbody>
      <tr><th class="shrink">Name</th><td>`)
//line views/vexport/ModelDetail.html:44
	qw422016.E().S(model.Name)
//line views/vexport/ModelDetail.html:44
	qw422016.N().S(`</td></tr>
      <tr><th>Package</th><td>`)
//line views/vexport/ModelDetail.html:45
	qw422016.E().S(model.Package)
//line views/vexport/ModelDetail.html:45
	qw422016.N().S(`</td></tr>
      <tr><th>Description</th><td>`)
//line views/vexport/ModelDetail.html:46
	qw422016.E().S(model.Description)
//line views/vexport/ModelDetail.html:46
	qw422016.N().S(`</td></tr>
      <tr>
        <th>Ordering</th>
        <td>
`)
//line views/vexport/ModelDetail.html:50
	for _, x := range model.Ordering {
//line views/vexport/ModelDetail.html:50
		qw422016.N().S(`          <div>`)
//line views/vexport/ModelDetail.html:51
		qw422016.E().S(x.String())
//line views/vexport/ModelDetail.html:51
		qw422016.N().S(`</div>
`)
//line views/vexport/ModelDetail.html:52
	}
//line views/vexport/ModelDetail.html:52
	qw422016.N().S(`        </td>
      </tr>
`)
//line views/vexport/ModelDetail.html:55
	if model.TitleOverride != "" {
//line views/vexport/ModelDetail.html:55
		qw422016.N().S(`      <tr><th>Title Override</th><td>`)
//line views/vexport/ModelDetail.html:56
		qw422016.E().S(model.TitleOverride)
//line views/vexport/ModelDetail.html:56
		qw422016.N().S(`</td></tr>
`)
//line views/vexport/ModelDetail.html:57
	}
//line views/vexport/ModelDetail.html:58
	if model.ProperOverride != "" {
//line views/vexport/ModelDetail.html:58
		qw422016.N().S(`      <tr><th>Proper Override</th><td>`)
//line views/vexport/ModelDetail.html:59
		qw422016.E().S(model.ProperOverride)
//line views/vexport/ModelDetail.html:59
		qw422016.N().S(`</td></tr>
`)
//line views/vexport/ModelDetail.html:60
	}
//line views/vexport/ModelDetail.html:61
	if model.RouteOverride != "" {
//line views/vexport/ModelDetail.html:61
		qw422016.N().S(`      <tr><th>Route Override</th><td>`)
//line views/vexport/ModelDetail.html:62
		qw422016.E().S(model.RouteOverride)
//line views/vexport/ModelDetail.html:62
		qw422016.N().S(`</td></tr>
`)
//line views/vexport/ModelDetail.html:63
	}
//line views/vexport/ModelDetail.html:64
	if len(model.GroupedColumns()) > 0 {
//line views/vexport/ModelDetail.html:64
		qw422016.N().S(`      <tr><th>Groupings</th><td>`)
//line views/vexport/ModelDetail.html:65
		qw422016.E().S(strings.Join(model.GroupedColumns().Names(), ", "))
//line views/vexport/ModelDetail.html:65
		qw422016.N().S(`</td></tr>
`)
//line views/vexport/ModelDetail.html:66
	}
//line views/vexport/ModelDetail.html:67
	if len(model.Tags) > 0 {
//line views/vexport/ModelDetail.html:67
		qw422016.N().S(`      <tr><th>Tags</th><td>`)
//line views/vexport/ModelDetail.html:68
		qw422016.E().S(strings.Join(model.Tags, ", "))
//line views/vexport/ModelDetail.html:68
		qw422016.N().S(`</td></tr>
`)
//line views/vexport/ModelDetail.html:69
	}
//line views/vexport/ModelDetail.html:70
	if len(model.Ordering) > 0 {
//line views/vexport/ModelDetail.html:70
		qw422016.N().S(`      <tr><th>Ordering</th><td>`)
//line views/vexport/ModelDetail.html:71
		components.StreamJSON(qw422016, model.Ordering)
//line views/vexport/ModelDetail.html:71
		qw422016.N().S(`</td></tr>
`)
//line views/vexport/ModelDetail.html:72
	}
//line views/vexport/ModelDetail.html:73
	if len(model.Search) > 0 {
//line views/vexport/ModelDetail.html:73
		qw422016.N().S(`      <tr><th>Search</th><td>`)
//line views/vexport/ModelDetail.html:74
		qw422016.E().S(strings.Join(model.Search, ", "))
//line views/vexport/ModelDetail.html:74
		qw422016.N().S(`</td></tr>
`)
//line views/vexport/ModelDetail.html:75
	}
//line views/vexport/ModelDetail.html:76
	if len(model.History) > 0 {
//line views/vexport/ModelDetail.html:76
		qw422016.N().S(`      <tr><th>History</th><td>`)
//line views/vexport/ModelDetail.html:77
		qw422016.E().S(model.History)
//line views/vexport/ModelDetail.html:77
		qw422016.N().S(`</td></tr>
`)
//line views/vexport/ModelDetail.html:78
	}
//line views/vexport/ModelDetail.html:79
	if len(model.Config) > 0 {
//line views/vexport/ModelDetail.html:79
		qw422016.N().S(`      <tr><th>Config</th><td>`)
//line views/vexport/ModelDetail.html:80
		components.StreamJSON(qw422016, model.Config)
//line views/vexport/ModelDetail.html:80
		qw422016.N().S(`</td></tr>
`)
//line views/vexport/ModelDetail.html:81
	}
//line views/vexport/ModelDetail.html:81
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vexport/ModelDetail.html:84
}

//line views/vexport/ModelDetail.html:84
func writemodelSummary(qq422016 qtio422016.Writer, model *model.Model, as *app.State, ps *cutil.PageState) {
//line views/vexport/ModelDetail.html:84
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vexport/ModelDetail.html:84
	streammodelSummary(qw422016, model, as, ps)
//line views/vexport/ModelDetail.html:84
	qt422016.ReleaseWriter(qw422016)
//line views/vexport/ModelDetail.html:84
}

//line views/vexport/ModelDetail.html:84
func modelSummary(model *model.Model, as *app.State, ps *cutil.PageState) string {
//line views/vexport/ModelDetail.html:84
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vexport/ModelDetail.html:84
	writemodelSummary(qb422016, model, as, ps)
//line views/vexport/ModelDetail.html:84
	qs422016 := string(qb422016.B)
//line views/vexport/ModelDetail.html:84
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vexport/ModelDetail.html:84
	return qs422016
//line views/vexport/ModelDetail.html:84
}

//line views/vexport/ModelDetail.html:86
func streammodelColumns(qw422016 *qt422016.Writer, model *model.Model, as *app.State, ps *cutil.PageState) {
//line views/vexport/ModelDetail.html:86
	qw422016.N().S(`
`)
//line views/vexport/ModelDetail.html:87
	if len(model.Columns) == 0 {
//line views/vexport/ModelDetail.html:87
		qw422016.N().S(`  <p><em>no columns</em></p>
`)
//line views/vexport/ModelDetail.html:89
	} else {
//line views/vexport/ModelDetail.html:89
		qw422016.N().S(`  <table class="mt min-200 full-width">
    <thead>
      <tr>
        <th class="shrink">Name</th>
        <th>Type</th>
        <th>Format</th>
        <th>Tags</th>
      </tr>
    </thead>
    <tbody>
`)
//line views/vexport/ModelDetail.html:100
		for _, col := range model.Columns {
//line views/vexport/ModelDetail.html:100
			qw422016.N().S(`      <tr>
        <td>`)
//line views/vexport/ModelDetail.html:102
			qw422016.E().S(col.Name)
//line views/vexport/ModelDetail.html:102
			qw422016.N().S(`</td>
        <td>`)
//line views/vexport/ModelDetail.html:103
			qw422016.E().S(col.Type.String())
//line views/vexport/ModelDetail.html:103
			qw422016.N().S(`</td>
        <td>`)
//line views/vexport/ModelDetail.html:104
			qw422016.E().S(col.Format)
//line views/vexport/ModelDetail.html:104
			qw422016.N().S(`</td>
        <td>`)
//line views/vexport/ModelDetail.html:105
			qw422016.E().S(strings.Join(col.Tags, ", "))
//line views/vexport/ModelDetail.html:105
			qw422016.N().S(`</td>
      </tr>
`)
//line views/vexport/ModelDetail.html:107
		}
//line views/vexport/ModelDetail.html:107
		qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vexport/ModelDetail.html:110
	}
//line views/vexport/ModelDetail.html:111
}

//line views/vexport/ModelDetail.html:111
func writemodelColumns(qq422016 qtio422016.Writer, model *model.Model, as *app.State, ps *cutil.PageState) {
//line views/vexport/ModelDetail.html:111
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vexport/ModelDetail.html:111
	streammodelColumns(qw422016, model, as, ps)
//line views/vexport/ModelDetail.html:111
	qt422016.ReleaseWriter(qw422016)
//line views/vexport/ModelDetail.html:111
}

//line views/vexport/ModelDetail.html:111
func modelColumns(model *model.Model, as *app.State, ps *cutil.PageState) string {
//line views/vexport/ModelDetail.html:111
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vexport/ModelDetail.html:111
	writemodelColumns(qb422016, model, as, ps)
//line views/vexport/ModelDetail.html:111
	qs422016 := string(qb422016.B)
//line views/vexport/ModelDetail.html:111
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vexport/ModelDetail.html:111
	return qs422016
//line views/vexport/ModelDetail.html:111
}

//line views/vexport/ModelDetail.html:113
func streammodelRelations(qw422016 *qt422016.Writer, model *model.Model, as *app.State, ps *cutil.PageState) {
//line views/vexport/ModelDetail.html:113
	qw422016.N().S(`
  <table class="mt min-200 full-width">
    <thead>
      <tr>
        <th class="shrink">Name</th>
        <th>Source</th>
        <th>Table</th>
        <th>Target</th>
      </tr>
    </thead>
    <tbody>
`)
//line views/vexport/ModelDetail.html:124
	for _, rel := range model.Relations {
//line views/vexport/ModelDetail.html:124
		qw422016.N().S(`      <tr>
        <td>`)
//line views/vexport/ModelDetail.html:126
		qw422016.E().S(rel.Name)
//line views/vexport/ModelDetail.html:126
		qw422016.N().S(`</td>
        <td>`)
//line views/vexport/ModelDetail.html:127
		qw422016.E().S(strings.Join(rel.Src, ", "))
//line views/vexport/ModelDetail.html:127
		qw422016.N().S(`</td>
        <td>`)
//line views/vexport/ModelDetail.html:128
		qw422016.E().S(rel.Table)
//line views/vexport/ModelDetail.html:128
		qw422016.N().S(`</td>
        <td>`)
//line views/vexport/ModelDetail.html:129
		qw422016.E().S(strings.Join(rel.Tgt, ", "))
//line views/vexport/ModelDetail.html:129
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vexport/ModelDetail.html:131
	}
//line views/vexport/ModelDetail.html:131
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vexport/ModelDetail.html:134
}

//line views/vexport/ModelDetail.html:134
func writemodelRelations(qq422016 qtio422016.Writer, model *model.Model, as *app.State, ps *cutil.PageState) {
//line views/vexport/ModelDetail.html:134
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vexport/ModelDetail.html:134
	streammodelRelations(qw422016, model, as, ps)
//line views/vexport/ModelDetail.html:134
	qt422016.ReleaseWriter(qw422016)
//line views/vexport/ModelDetail.html:134
}

//line views/vexport/ModelDetail.html:134
func modelRelations(model *model.Model, as *app.State, ps *cutil.PageState) string {
//line views/vexport/ModelDetail.html:134
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vexport/ModelDetail.html:134
	writemodelRelations(qb422016, model, as, ps)
//line views/vexport/ModelDetail.html:134
	qs422016 := string(qb422016.B)
//line views/vexport/ModelDetail.html:134
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vexport/ModelDetail.html:134
	return qs422016
//line views/vexport/ModelDetail.html:134
}

//line views/vexport/ModelDetail.html:136
func streammodelIndexes(qw422016 *qt422016.Writer, model *model.Model, as *app.State, ps *cutil.PageState) {
//line views/vexport/ModelDetail.html:136
	qw422016.N().S(`
  <table class="mt min-200 full-width">
    <thead>
      <tr>
        <th class="shrink">Name</th>
        <th>Declaration</th>
      </tr>
    </thead>
    <tbody>
`)
//line views/vexport/ModelDetail.html:145
	for _, idx := range model.Indexes {
//line views/vexport/ModelDetail.html:145
		qw422016.N().S(`      <tr>
        <td>`)
//line views/vexport/ModelDetail.html:147
		qw422016.E().S(idx.Name)
//line views/vexport/ModelDetail.html:147
		qw422016.N().S(`</td>
        <td>`)
//line views/vexport/ModelDetail.html:148
		qw422016.E().S(idx.Decl)
//line views/vexport/ModelDetail.html:148
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vexport/ModelDetail.html:150
	}
//line views/vexport/ModelDetail.html:150
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vexport/ModelDetail.html:153
}

//line views/vexport/ModelDetail.html:153
func writemodelIndexes(qq422016 qtio422016.Writer, model *model.Model, as *app.State, ps *cutil.PageState) {
//line views/vexport/ModelDetail.html:153
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vexport/ModelDetail.html:153
	streammodelIndexes(qw422016, model, as, ps)
//line views/vexport/ModelDetail.html:153
	qt422016.ReleaseWriter(qw422016)
//line views/vexport/ModelDetail.html:153
}

//line views/vexport/ModelDetail.html:153
func modelIndexes(model *model.Model, as *app.State, ps *cutil.PageState) string {
//line views/vexport/ModelDetail.html:153
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vexport/ModelDetail.html:153
	writemodelIndexes(qb422016, model, as, ps)
//line views/vexport/ModelDetail.html:153
	qs422016 := string(qb422016.B)
//line views/vexport/ModelDetail.html:153
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vexport/ModelDetail.html:153
	return qs422016
//line views/vexport/ModelDetail.html:153
}
