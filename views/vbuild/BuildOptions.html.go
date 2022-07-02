// Code generated by qtc from "BuildOptions.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vbuild/BuildOptions.html:1
package vbuild

//line views/vbuild/BuildOptions.html:1
import (
	"projectforge.dev/projectforge/app/project/action"
)

//line views/vbuild/BuildOptions.html:5
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vbuild/BuildOptions.html:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vbuild/BuildOptions.html:5
func StreamBuildOptions(qw422016 *qt422016.Writer, key string) {
//line views/vbuild/BuildOptions.html:5
	qw422016.N().S(`
  <div class="card">
    <h3>Build your project</h3>
    <table class="mt">
      <thead>
        <tr>
          <th>Action</th>
          <th>Description</th>
        </tr>
      </thead>
      <tbody>
`)
//line views/vbuild/BuildOptions.html:16
	for _, b := range action.AllBuilds {
//line views/vbuild/BuildOptions.html:16
		qw422016.N().S(`        <tr>
          <td><a href="/run/`)
//line views/vbuild/BuildOptions.html:18
		qw422016.E().S(key)
//line views/vbuild/BuildOptions.html:18
		qw422016.N().S(`/build?phase=`)
//line views/vbuild/BuildOptions.html:18
		qw422016.E().S(b.Key)
//line views/vbuild/BuildOptions.html:18
		qw422016.N().S(`" title="`)
//line views/vbuild/BuildOptions.html:18
		qw422016.E().S(b.Description)
//line views/vbuild/BuildOptions.html:18
		qw422016.N().S(`"><button>`)
//line views/vbuild/BuildOptions.html:18
		qw422016.E().S(b.Title)
//line views/vbuild/BuildOptions.html:18
		qw422016.N().S(`</button></a></td>
          <td>`)
//line views/vbuild/BuildOptions.html:19
		qw422016.E().S(b.Description)
//line views/vbuild/BuildOptions.html:19
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vbuild/BuildOptions.html:21
	}
//line views/vbuild/BuildOptions.html:21
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vbuild/BuildOptions.html:25
}

//line views/vbuild/BuildOptions.html:25
func WriteBuildOptions(qq422016 qtio422016.Writer, key string) {
//line views/vbuild/BuildOptions.html:25
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vbuild/BuildOptions.html:25
	StreamBuildOptions(qw422016, key)
//line views/vbuild/BuildOptions.html:25
	qt422016.ReleaseWriter(qw422016)
//line views/vbuild/BuildOptions.html:25
}

//line views/vbuild/BuildOptions.html:25
func BuildOptions(key string) string {
//line views/vbuild/BuildOptions.html:25
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vbuild/BuildOptions.html:25
	WriteBuildOptions(qb422016, key)
//line views/vbuild/BuildOptions.html:25
	qs422016 := string(qb422016.B)
//line views/vbuild/BuildOptions.html:25
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vbuild/BuildOptions.html:25
	return qs422016
//line views/vbuild/BuildOptions.html:25
}
