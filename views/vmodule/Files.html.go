// Code generated by qtc from "Files.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vmodule/Files.html:1
package vmodule

//line views/vmodule/Files.html:1
import (
	"path/filepath"
	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/app/module"
	"projectforge.dev/projectforge/views/layout"
	"projectforge.dev/projectforge/views/vfile"
)

//line views/vmodule/Files.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vmodule/Files.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vmodule/Files.html:10
type Files struct {
	layout.Basic
	Module *module.Module
	Path   []string
}

//line views/vmodule/Files.html:16
func (p *Files) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vmodule/Files.html:16
	qw422016.N().S(`
`)
//line views/vmodule/Files.html:18
	mod := p.Module
	fs := as.Services.Modules.GetFilesystem(mod.Key)
	u := "/m/" + mod.Key + "/fs"

//line views/vmodule/Files.html:21
	qw422016.N().S(`
  `)
//line views/vmodule/Files.html:23
	StreamSummary(qw422016, mod, nil, ps)
//line views/vmodule/Files.html:23
	qw422016.N().S(`

`)
//line views/vmodule/Files.html:25
	if fs.IsDir(filepath.Join(p.Path...)) {
//line views/vmodule/Files.html:26
		files := fs.ListFiles(filepath.Join(p.Path...), nil, ps.Logger)

//line views/vmodule/Files.html:26
		qw422016.N().S(`  <div class="card">
    `)
//line views/vmodule/Files.html:28
		vfile.StreamList(qw422016, p.Path, files, fs, u, as, ps)
//line views/vmodule/Files.html:28
		qw422016.N().S(`
  </div>
`)
//line views/vmodule/Files.html:30
	} else {
//line views/vmodule/Files.html:32
		b, err := fs.ReadFile(filepath.Join(p.Path...))
		if err != nil {
			panic(err)
		}

//line views/vmodule/Files.html:36
		qw422016.N().S(`  `)
//line views/vmodule/Files.html:37
		vfile.StreamDetail(qw422016, p.Path, b, u, as, ps)
//line views/vmodule/Files.html:37
		qw422016.N().S(`
`)
//line views/vmodule/Files.html:38
	}
//line views/vmodule/Files.html:39
}

//line views/vmodule/Files.html:39
func (p *Files) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vmodule/Files.html:39
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vmodule/Files.html:39
	p.StreamBody(qw422016, as, ps)
//line views/vmodule/Files.html:39
	qt422016.ReleaseWriter(qw422016)
//line views/vmodule/Files.html:39
}

//line views/vmodule/Files.html:39
func (p *Files) Body(as *app.State, ps *cutil.PageState) string {
//line views/vmodule/Files.html:39
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vmodule/Files.html:39
	p.WriteBody(qb422016, as, ps)
//line views/vmodule/Files.html:39
	qs422016 := string(qb422016.B)
//line views/vmodule/Files.html:39
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vmodule/Files.html:39
	return qs422016
//line views/vmodule/Files.html:39
}
