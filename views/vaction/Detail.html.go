// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vaction/Detail.html:1
package vaction

//line views/vaction/Detail.html:1
import (
	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/action"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/app/util"
)

//line views/vaction/Detail.html:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vaction/Detail.html:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vaction/Detail.html:8
func StreamDetail(qw422016 *qt422016.Writer, cfg util.ValueMap, res *action.Result, includeSkipped bool, as *app.State, ps *cutil.PageState) {
//line views/vaction/Detail.html:8
	qw422016.N().S(`
`)
//line views/vaction/Detail.html:9
	if len(cfg) > 0 {
//line views/vaction/Detail.html:9
		qw422016.N().S(`  <div class="card">
    <h3>Config</h3>
    <table>
      <tbody>
`)
//line views/vaction/Detail.html:14
		for _, k := range cfg.Keys() {
//line views/vaction/Detail.html:14
			qw422016.N().S(`        <tr>
          <th class="shrink">`)
//line views/vaction/Detail.html:16
			qw422016.E().S(k)
//line views/vaction/Detail.html:16
			qw422016.N().S(`</th>
          <td>`)
//line views/vaction/Detail.html:17
			qw422016.E().V(cfg[k])
//line views/vaction/Detail.html:17
			qw422016.N().S(`</td>
        </tr>
`)
//line views/vaction/Detail.html:19
		}
//line views/vaction/Detail.html:19
		qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vaction/Detail.html:23
	}
//line views/vaction/Detail.html:24
	if len(res.Errors) > 0 {
//line views/vaction/Detail.html:24
		qw422016.N().S(`  <div class="card">
    <div class="right">`)
//line views/vaction/Detail.html:26
		qw422016.N().D(len(res.Errors))
//line views/vaction/Detail.html:26
		qw422016.N().S(` `)
//line views/vaction/Detail.html:26
		qw422016.E().S(util.StringPluralMaybe("error", len(res.Errors)))
//line views/vaction/Detail.html:26
		qw422016.N().S(`</div>
    <h3>Errors</h3>
    <ul>
`)
//line views/vaction/Detail.html:29
		for _, e := range res.Errors {
//line views/vaction/Detail.html:29
			qw422016.N().S(`      <li class="error">`)
//line views/vaction/Detail.html:30
			qw422016.E().S(e)
//line views/vaction/Detail.html:30
			qw422016.N().S(`</li>
`)
//line views/vaction/Detail.html:31
		}
//line views/vaction/Detail.html:31
		qw422016.N().S(`    </ul>
  </div>
`)
//line views/vaction/Detail.html:34
	}
//line views/vaction/Detail.html:35
	if len(res.Logs) > 0 {
//line views/vaction/Detail.html:35
		qw422016.N().S(`  <div class="card">
    <h3>Logs</h3>
    <table>
      <tbody>
`)
//line views/vaction/Detail.html:40
		for _, l := range res.Logs {
//line views/vaction/Detail.html:40
			qw422016.N().S(`        <tr>
          <td><pre>`)
//line views/vaction/Detail.html:42
			qw422016.E().S(l)
//line views/vaction/Detail.html:42
			qw422016.N().S(`</pre></td>
        </tr>
`)
//line views/vaction/Detail.html:44
		}
//line views/vaction/Detail.html:44
		qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vaction/Detail.html:48
	}
//line views/vaction/Detail.html:49
	for _, mr := range res.Modules {
//line views/vaction/Detail.html:51
		links := ""
		for mIdx, mk := range mr.Keys {
			if mIdx > 0 {
				links += ", "
			}
			links += `<a href="/m/` + mk + `">` + mk + `</a>`
		}

//line views/vaction/Detail.html:58
		qw422016.N().S(`    <div class="card">
      <div class="right">`)
//line views/vaction/Detail.html:60
		qw422016.E().S(util.MicrosToMillis(mr.Duration))
//line views/vaction/Detail.html:60
		qw422016.N().S(`</div>
      <h3>`)
//line views/vaction/Detail.html:61
		qw422016.E().S(util.StringPluralMaybe("Module", len(mr.Keys)))
//line views/vaction/Detail.html:61
		qw422016.N().S(` [`)
//line views/vaction/Detail.html:61
		qw422016.N().S(links)
//line views/vaction/Detail.html:61
		qw422016.N().S(`]</h3>
      <div class="right">`)
//line views/vaction/Detail.html:62
		qw422016.N().S(res.StatusLog())
//line views/vaction/Detail.html:62
		qw422016.N().S(`</div>
      <em>`)
//line views/vaction/Detail.html:63
		qw422016.E().S(mr.Status)
//line views/vaction/Detail.html:63
		qw422016.N().S(`</em>
`)
//line views/vaction/Detail.html:64
		if len(mr.Actions) > 0 {
//line views/vaction/Detail.html:64
			qw422016.N().S(`      <h4>Actions</h4>
`)
//line views/vaction/Detail.html:66
			for _, a := range mr.Actions {
//line views/vaction/Detail.html:66
				qw422016.N().S(`        <a href="`)
//line views/vaction/Detail.html:67
				qw422016.E().S(a.URL())
//line views/vaction/Detail.html:67
				qw422016.N().S(`"><button>`)
//line views/vaction/Detail.html:67
				qw422016.E().S(a.Title)
//line views/vaction/Detail.html:67
				qw422016.N().S(`</button></a>
`)
//line views/vaction/Detail.html:68
			}
//line views/vaction/Detail.html:69
		}
//line views/vaction/Detail.html:70
		diffs := mr.DiffsFiltered(includeSkipped)

//line views/vaction/Detail.html:71
		if len(diffs) > 0 {
//line views/vaction/Detail.html:71
			qw422016.N().S(`      `)
//line views/vaction/Detail.html:72
			streamrenderDiffs(qw422016, res.Project.Key, res.Action, diffs, as, ps)
//line views/vaction/Detail.html:72
			qw422016.N().S(`
`)
//line views/vaction/Detail.html:73
		}
//line views/vaction/Detail.html:73
		qw422016.N().S(`    </div>
`)
//line views/vaction/Detail.html:75
	}
//line views/vaction/Detail.html:76
}

//line views/vaction/Detail.html:76
func WriteDetail(qq422016 qtio422016.Writer, cfg util.ValueMap, res *action.Result, includeSkipped bool, as *app.State, ps *cutil.PageState) {
//line views/vaction/Detail.html:76
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vaction/Detail.html:76
	StreamDetail(qw422016, cfg, res, includeSkipped, as, ps)
//line views/vaction/Detail.html:76
	qt422016.ReleaseWriter(qw422016)
//line views/vaction/Detail.html:76
}

//line views/vaction/Detail.html:76
func Detail(cfg util.ValueMap, res *action.Result, includeSkipped bool, as *app.State, ps *cutil.PageState) string {
//line views/vaction/Detail.html:76
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vaction/Detail.html:76
	WriteDetail(qb422016, cfg, res, includeSkipped, as, ps)
//line views/vaction/Detail.html:76
	qs422016 := string(qb422016.B)
//line views/vaction/Detail.html:76
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vaction/Detail.html:76
	return qs422016
//line views/vaction/Detail.html:76
}
