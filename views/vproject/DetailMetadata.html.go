// Code generated by qtc from "DetailMetadata.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vproject/DetailMetadata.html:1
package vproject

//line views/vproject/DetailMetadata.html:1
import (
	"strings"

	"projectforge.dev/app"
	"projectforge.dev/app/controller/cutil"
	"projectforge.dev/app/project"
)

//line views/vproject/DetailMetadata.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vproject/DetailMetadata.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vproject/DetailMetadata.html:9
func StreamDetailMetadata(qw422016 *qt422016.Writer, info *project.Info, as *app.State, ps *cutil.PageState) {
//line views/vproject/DetailMetadata.html:9
	qw422016.N().S(`
  <table class="min-200 full-width">
    <tbody>
      <tr class="shrink">
        <th class="shrink">Organization</th>
        <td>`)
//line views/vproject/DetailMetadata.html:14
	qw422016.E().S(info.Org)
//line views/vproject/DetailMetadata.html:14
	qw422016.N().S(`</td>
      </tr>
      <tr>
        <th>Author ID</th>
        <td>`)
//line views/vproject/DetailMetadata.html:18
	qw422016.E().S(info.AuthorID)
//line views/vproject/DetailMetadata.html:18
	qw422016.N().S(`</td>
      </tr>
      <tr>
        <th>Author</th>
        <td>`)
//line views/vproject/DetailMetadata.html:22
	qw422016.E().S(info.AuthorName)
//line views/vproject/DetailMetadata.html:22
	qw422016.N().S(` &lt;<a href="mailto:`)
//line views/vproject/DetailMetadata.html:22
	qw422016.E().S(info.AuthorEmail)
//line views/vproject/DetailMetadata.html:22
	qw422016.N().S(`">`)
//line views/vproject/DetailMetadata.html:22
	qw422016.E().S(info.AuthorEmail)
//line views/vproject/DetailMetadata.html:22
	qw422016.N().S(`</a>&gt;</td>
      </tr>
      <tr>
        <th>License</th>
        <td>`)
//line views/vproject/DetailMetadata.html:26
	qw422016.E().S(info.License)
//line views/vproject/DetailMetadata.html:26
	qw422016.N().S(`</td>
      </tr>
      <tr>
        <th>Homepage</th>
        <td>`)
//line views/vproject/DetailMetadata.html:30
	qw422016.E().S(info.Homepage)
//line views/vproject/DetailMetadata.html:30
	qw422016.N().S(`</td>
      </tr>
      <tr>
        <th>Source Code</th>
        <td>`)
//line views/vproject/DetailMetadata.html:34
	qw422016.E().S(info.Sourcecode)
//line views/vproject/DetailMetadata.html:34
	qw422016.N().S(`</td>
      </tr>
      <tr>
        <th>Summary</th>
        <td>`)
//line views/vproject/DetailMetadata.html:38
	qw422016.E().S(info.Summary)
//line views/vproject/DetailMetadata.html:38
	qw422016.N().S(`</td>
      </tr>
      <tr>
        <th>Description</th>
        <td>`)
//line views/vproject/DetailMetadata.html:42
	qw422016.E().S(info.Description)
//line views/vproject/DetailMetadata.html:42
	qw422016.N().S(`</td>
      </tr>
      <tr>
        <th>CI</th>
        <td>`)
//line views/vproject/DetailMetadata.html:46
	qw422016.E().S(info.CI)
//line views/vproject/DetailMetadata.html:46
	qw422016.N().S(`</td>
      </tr>
      <tr>
        <th>Homebrew</th>
        <td>`)
//line views/vproject/DetailMetadata.html:50
	qw422016.E().S(info.Homebrew)
//line views/vproject/DetailMetadata.html:50
	qw422016.N().S(`</td>
      </tr>
      <tr>
        <th>Bundle</th>
        <td>`)
//line views/vproject/DetailMetadata.html:54
	qw422016.E().S(info.Bundle)
//line views/vproject/DetailMetadata.html:54
	qw422016.N().S(`</td>
      </tr>
      <tr>
        <th>Signing Identity</th>
        <td>`)
//line views/vproject/DetailMetadata.html:58
	qw422016.E().S(info.SigningIdentity)
//line views/vproject/DetailMetadata.html:58
	qw422016.N().S(`</td>
      </tr>
      <tr>
        <th>Slack Webhook</th>
        <td>`)
//line views/vproject/DetailMetadata.html:62
	qw422016.E().S(info.Slack)
//line views/vproject/DetailMetadata.html:62
	qw422016.N().S(`</td>
      </tr>
      <tr>
        <th>Java Package</th>
        <td>`)
//line views/vproject/DetailMetadata.html:66
	qw422016.E().S(info.JavaPackage)
//line views/vproject/DetailMetadata.html:66
	qw422016.N().S(`</td>
      </tr>
      <tr>
        <th>Go Binary</th>
        <td>`)
//line views/vproject/DetailMetadata.html:70
	qw422016.E().S(info.GoBinary)
//line views/vproject/DetailMetadata.html:70
	qw422016.N().S(`</td>
      </tr>
      <tr>
        <th>Go Version</th>
        <td>`)
//line views/vproject/DetailMetadata.html:74
	qw422016.E().S(info.GoVersion)
//line views/vproject/DetailMetadata.html:74
	qw422016.N().S(`</td>
      </tr>
      <tr>
        <th>Extra Files</th>
        <td>`)
//line views/vproject/DetailMetadata.html:78
	qw422016.E().S(strings.Join(info.ExtraFiles, ", "))
//line views/vproject/DetailMetadata.html:78
	qw422016.N().S(`</td>
      </tr>
`)
//line views/vproject/DetailMetadata.html:80
	if len(info.ModuleDefs) > 0 {
//line views/vproject/DetailMetadata.html:80
		qw422016.N().S(`      <tr>
        <th>Module Definitions</th>
        <td>
`)
//line views/vproject/DetailMetadata.html:84
		for _, x := range info.ModuleDefs {
//line views/vproject/DetailMetadata.html:84
			qw422016.N().S(`          <a href="/m/`)
//line views/vproject/DetailMetadata.html:85
			qw422016.E().S(x.Key)
//line views/vproject/DetailMetadata.html:85
			qw422016.N().S(`">`)
//line views/vproject/DetailMetadata.html:85
			qw422016.E().S(x.Key)
//line views/vproject/DetailMetadata.html:85
			qw422016.N().S(`</a> <em>(`)
//line views/vproject/DetailMetadata.html:85
			qw422016.E().S(x.Path)
//line views/vproject/DetailMetadata.html:85
			qw422016.N().S(`)</em><br />
`)
//line views/vproject/DetailMetadata.html:86
		}
//line views/vproject/DetailMetadata.html:86
		qw422016.N().S(`        </td>
      </tr>
`)
//line views/vproject/DetailMetadata.html:89
	}
//line views/vproject/DetailMetadata.html:89
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vproject/DetailMetadata.html:92
}

//line views/vproject/DetailMetadata.html:92
func WriteDetailMetadata(qq422016 qtio422016.Writer, info *project.Info, as *app.State, ps *cutil.PageState) {
//line views/vproject/DetailMetadata.html:92
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vproject/DetailMetadata.html:92
	StreamDetailMetadata(qw422016, info, as, ps)
//line views/vproject/DetailMetadata.html:92
	qt422016.ReleaseWriter(qw422016)
//line views/vproject/DetailMetadata.html:92
}

//line views/vproject/DetailMetadata.html:92
func DetailMetadata(info *project.Info, as *app.State, ps *cutil.PageState) string {
//line views/vproject/DetailMetadata.html:92
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vproject/DetailMetadata.html:92
	WriteDetailMetadata(qb422016, info, as, ps)
//line views/vproject/DetailMetadata.html:92
	qs422016 := string(qb422016.B)
//line views/vproject/DetailMetadata.html:92
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vproject/DetailMetadata.html:92
	return qs422016
//line views/vproject/DetailMetadata.html:92
}
