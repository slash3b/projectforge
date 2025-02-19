// Code generated by qtc from "FeatureFiles.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vsite/FeatureFiles.html:1
package vsite

//line views/vsite/FeatureFiles.html:1
import (
	"path/filepath"

	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/app/module"
	"projectforge.dev/projectforge/views/components"
	"projectforge.dev/projectforge/views/layout"
	"projectforge.dev/projectforge/views/vfile"
)

//line views/vsite/FeatureFiles.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsite/FeatureFiles.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsite/FeatureFiles.html:12
type FeatureFiles struct {
	layout.Basic
	Module *module.Module
	Path   []string
}

//line views/vsite/FeatureFiles.html:18
func (p *FeatureFiles) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsite/FeatureFiles.html:18
	qw422016.N().S(`
`)
//line views/vsite/FeatureFiles.html:20
	mod := p.Module
	fs := as.Services.Modules.GetFilesystem(mod.Key)
	u := "/features/" + mod.Key + "/files"

//line views/vsite/FeatureFiles.html:23
	qw422016.N().S(`
`)
//line views/vsite/FeatureFiles.html:25
	if fs.IsDir(filepath.Join(p.Path...)) {
//line views/vsite/FeatureFiles.html:26
		files := fs.ListFiles(filepath.Join(p.Path...), nil, ps.Logger)

//line views/vsite/FeatureFiles.html:26
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vsite/FeatureFiles.html:28
		components.StreamSVGRefIcon(qw422016, p.Module.Icon, ps)
//line views/vsite/FeatureFiles.html:28
		qw422016.E().S(p.Module.Title())
//line views/vsite/FeatureFiles.html:28
		qw422016.N().S(`</h3>
    `)
//line views/vsite/FeatureFiles.html:29
		vfile.StreamList(qw422016, p.Path, files, fs, u, as, ps)
//line views/vsite/FeatureFiles.html:29
		qw422016.N().S(`
  </div>
`)
//line views/vsite/FeatureFiles.html:31
	} else {
//line views/vsite/FeatureFiles.html:33
		b, err := fs.ReadFile(filepath.Join(p.Path...))
		if err != nil {
			panic(err)
		}

//line views/vsite/FeatureFiles.html:37
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vsite/FeatureFiles.html:39
		components.StreamSVGRefIcon(qw422016, p.Module.Icon, ps)
//line views/vsite/FeatureFiles.html:39
		qw422016.E().S(p.Module.Title())
//line views/vsite/FeatureFiles.html:39
		qw422016.N().S(`</h3>
    `)
//line views/vsite/FeatureFiles.html:40
		vfile.StreamDetail(qw422016, p.Path, b, u, as, ps)
//line views/vsite/FeatureFiles.html:40
		qw422016.N().S(`
  </div>
`)
//line views/vsite/FeatureFiles.html:42
	}
//line views/vsite/FeatureFiles.html:43
}

//line views/vsite/FeatureFiles.html:43
func (p *FeatureFiles) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsite/FeatureFiles.html:43
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsite/FeatureFiles.html:43
	p.StreamBody(qw422016, as, ps)
//line views/vsite/FeatureFiles.html:43
	qt422016.ReleaseWriter(qw422016)
//line views/vsite/FeatureFiles.html:43
}

//line views/vsite/FeatureFiles.html:43
func (p *FeatureFiles) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsite/FeatureFiles.html:43
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsite/FeatureFiles.html:43
	p.WriteBody(qb422016, as, ps)
//line views/vsite/FeatureFiles.html:43
	qs422016 := string(qb422016.B)
//line views/vsite/FeatureFiles.html:43
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsite/FeatureFiles.html:43
	return qs422016
//line views/vsite/FeatureFiles.html:43
}
