// Code generated by qtc from "DetailSettings.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vproject/DetailSettings.html:1
package vproject

//line views/vproject/DetailSettings.html:1
import (
	"projectforge.dev/app"
	"projectforge.dev/app/controller/cutil"
	"projectforge.dev/app/project"
	"projectforge.dev/views/components"
	"strings"
)

//line views/vproject/DetailSettings.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vproject/DetailSettings.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vproject/DetailSettings.html:9
func StreamDetailSettings(qw422016 *qt422016.Writer, prj *project.Project, as *app.State, ps *cutil.PageState) {
//line views/vproject/DetailSettings.html:9
	qw422016.N().S(`
  <table class="min-200 full-width">
    <tbody>
      <tr>
        <th class="shrink">Key</th>
        <td>`)
//line views/vproject/DetailSettings.html:14
	qw422016.E().S(prj.Key)
//line views/vproject/DetailSettings.html:14
	qw422016.N().S(`</td>
      </tr>
      <tr>
        <th>Name</th>
        <td>`)
//line views/vproject/DetailSettings.html:18
	qw422016.E().S(prj.Name)
//line views/vproject/DetailSettings.html:18
	qw422016.N().S(`</td>
      </tr>
      <tr>
        <th>Version</th>
        <td>`)
//line views/vproject/DetailSettings.html:22
	qw422016.E().S(prj.Version)
//line views/vproject/DetailSettings.html:22
	qw422016.N().S(`</td>
      </tr>
      <tr>
        <th>Package</th>
        <td>`)
//line views/vproject/DetailSettings.html:26
	qw422016.E().S(prj.Package)
//line views/vproject/DetailSettings.html:26
	qw422016.N().S(`</td>
      </tr>
      <tr>
        <th>Args</th>
        <td>`)
//line views/vproject/DetailSettings.html:30
	qw422016.E().S(prj.Args)
//line views/vproject/DetailSettings.html:30
	qw422016.N().S(`</td>
      </tr>
      <tr>
        <th>Port</th>
        <td><a href="http://localhost:`)
//line views/vproject/DetailSettings.html:34
	qw422016.N().D(prj.Port)
//line views/vproject/DetailSettings.html:34
	qw422016.N().S(`" target="_blank">`)
//line views/vproject/DetailSettings.html:34
	qw422016.N().D(prj.Port)
//line views/vproject/DetailSettings.html:34
	qw422016.N().S(`</a></td>
      </tr>
      <tr>
        <th>Ignore</th>
        <td>`)
//line views/vproject/DetailSettings.html:38
	qw422016.E().S(strings.Join(prj.Ignore, ", "))
//line views/vproject/DetailSettings.html:38
	qw422016.N().S(`</td>
      </tr>
      <tr>
        <th>Path</th>
        <td>`)
//line views/vproject/DetailSettings.html:42
	qw422016.E().S(prj.Path)
//line views/vproject/DetailSettings.html:42
	qw422016.N().S(`</td>
      </tr>
      <tr>
        <th>JSON</th>
        <td><a href="#modal-`)
//line views/vproject/DetailSettings.html:46
	qw422016.E().S(prj.Key)
//line views/vproject/DetailSettings.html:46
	qw422016.N().S(`"><button type="button">JSON</button></a></td>
      </tr>
    </tbody>
  </table>
  `)
//line views/vproject/DetailSettings.html:50
	components.StreamJSONModal(qw422016, prj.Key, "Project JSON", prj, 1)
//line views/vproject/DetailSettings.html:50
	qw422016.N().S(`
`)
//line views/vproject/DetailSettings.html:51
}

//line views/vproject/DetailSettings.html:51
func WriteDetailSettings(qq422016 qtio422016.Writer, prj *project.Project, as *app.State, ps *cutil.PageState) {
//line views/vproject/DetailSettings.html:51
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vproject/DetailSettings.html:51
	StreamDetailSettings(qw422016, prj, as, ps)
//line views/vproject/DetailSettings.html:51
	qt422016.ReleaseWriter(qw422016)
//line views/vproject/DetailSettings.html:51
}

//line views/vproject/DetailSettings.html:51
func DetailSettings(prj *project.Project, as *app.State, ps *cutil.PageState) string {
//line views/vproject/DetailSettings.html:51
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vproject/DetailSettings.html:51
	WriteDetailSettings(qb422016, prj, as, ps)
//line views/vproject/DetailSettings.html:51
	qs422016 := string(qb422016.B)
//line views/vproject/DetailSettings.html:51
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vproject/DetailSettings.html:51
	return qs422016
//line views/vproject/DetailSettings.html:51
}
