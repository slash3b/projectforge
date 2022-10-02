// Code generated by qtc from "Settings.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vadmin/Settings.html:2
package vadmin

//line views/vadmin/Settings.html:2
import (
	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/app/lib/user"
	"projectforge.dev/projectforge/app/util"
	"projectforge.dev/projectforge/views/components"
	"projectforge.dev/projectforge/views/layout"
)

//line views/vadmin/Settings.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/Settings.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/Settings.html:11
type Settings struct {
	layout.Basic
	Perms user.Permissions
}

//line views/vadmin/Settings.html:16
func (p *Settings) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Settings.html:16
	qw422016.N().S(`
  <div class="card">
`)
//line views/vadmin/Settings.html:18
	if util.AppSource != "" {
//line views/vadmin/Settings.html:18
		qw422016.N().S(`    <div class="right"><a href="`)
//line views/vadmin/Settings.html:19
		qw422016.E().S(util.AppSource)
//line views/vadmin/Settings.html:19
		qw422016.N().S(`"><button>Github</button></a></div>
`)
//line views/vadmin/Settings.html:20
	}
//line views/vadmin/Settings.html:20
	qw422016.N().S(`    <h3 title="github:`)
//line views/vadmin/Settings.html:21
	qw422016.E().S(as.BuildInfo.Commit)
//line views/vadmin/Settings.html:21
	qw422016.N().S(`">`)
//line views/vadmin/Settings.html:21
	components.StreamSVGRefIcon(qw422016, `cog`, ps)
//line views/vadmin/Settings.html:21
	qw422016.E().S(util.AppName)
//line views/vadmin/Settings.html:21
	qw422016.N().S(` `)
//line views/vadmin/Settings.html:21
	qw422016.E().S(as.BuildInfo.String())
//line views/vadmin/Settings.html:21
	qw422016.N().S(`</h3>
`)
//line views/vadmin/Settings.html:22
	if util.AppLegal != "" {
//line views/vadmin/Settings.html:22
		qw422016.N().S(`    <div class="mt">`)
//line views/vadmin/Settings.html:23
		qw422016.N().S(util.AppLegal)
//line views/vadmin/Settings.html:23
		qw422016.N().S(`</div>
`)
//line views/vadmin/Settings.html:24
	}
//line views/vadmin/Settings.html:25
	if util.AppURL != "" {
//line views/vadmin/Settings.html:25
		qw422016.N().S(`    <p><a href="`)
//line views/vadmin/Settings.html:26
		qw422016.N().S(util.AppURL)
//line views/vadmin/Settings.html:26
		qw422016.N().S(`">`)
//line views/vadmin/Settings.html:26
		qw422016.N().S(util.AppURL)
//line views/vadmin/Settings.html:26
		qw422016.N().S(`</a></p>
`)
//line views/vadmin/Settings.html:27
	}
//line views/vadmin/Settings.html:27
	qw422016.N().S(`  </div>

  <div class="flex-wrap" style="align-items: stretch;">
    <div class="card" style="flex-grow: 1;">
      <h3>`)
//line views/vadmin/Settings.html:32
	components.StreamSVGRefIcon(qw422016, `archive`, ps)
//line views/vadmin/Settings.html:32
	qw422016.N().S(`Admin Functions</h3>
      <ul class="mt">
        <li><a href="/admin/modules">View Go modules</a></li>
        <li><a href="/admin/gc">Collect garbage</a></li>
        <li><a href="/theme">Edit Themes</a></li>
        <li><a href="/admin/exec">Managed Processes</a></li>
      </ul>
    </div>
    <div class="card" style="flex-grow: 1;">
      <h3>`)
//line views/vadmin/Settings.html:41
	components.StreamSVGRefIcon(qw422016, `bolt`, ps)
//line views/vadmin/Settings.html:41
	qw422016.N().S(`HTTP Methods</h3>
      <ul class="mt">
        <li><a href="/admin/sitemap">Sitemap</a></li>
        <li><a href="/admin/routes">View HTTP routes</a></li>
        <li><a href="/admin/session">Parse and display session</a></li>
        <li><a href="/admin/request">Debug HTTP request</a></li>
      </ul>
    </div>
    <div class="card" style="flex-grow: 1;">
      <h3>`)
//line views/vadmin/Settings.html:50
	components.StreamSVGRefIcon(qw422016, `cog`, ps)
//line views/vadmin/Settings.html:50
	qw422016.N().S(`App Profiling</h3>
      <ul class="mt">
        <li><a href="/admin/cpu/start">Start CPU profile</a></li>
        <li><a href="/admin/cpu/stop">Stop CPU profile</a></li>
        <li><a href="/admin/heap">Write memory dump</a></li>
        <li><a href="/admin/memusage">Memory Usage</a></li>
      </ul>
    </div>
  </div>
`)
//line views/vadmin/Settings.html:59
}

//line views/vadmin/Settings.html:59
func (p *Settings) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Settings.html:59
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Settings.html:59
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Settings.html:59
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Settings.html:59
}

//line views/vadmin/Settings.html:59
func (p *Settings) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Settings.html:59
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Settings.html:59
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Settings.html:59
	qs422016 := string(qb422016.B)
//line views/vadmin/Settings.html:59
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Settings.html:59
	return qs422016
//line views/vadmin/Settings.html:59
}
