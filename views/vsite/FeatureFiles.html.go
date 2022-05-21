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
	"projectforge.dev/projectforge/views/layout"
	"projectforge.dev/projectforge/views/vfile"
)

//line views/vsite/FeatureFiles.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsite/FeatureFiles.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsite/FeatureFiles.html:10
type FeatureFiles struct {
	layout.Basic
	Module *module.Module
	Path   []string
}

//line views/vsite/FeatureFiles.html:16
func (p *FeatureFiles) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsite/FeatureFiles.html:16
	qw422016.N().S(`
`)
//line views/vsite/FeatureFiles.html:18
	mod := p.Module
	fs := as.Services.Modules.GetFilesystem(mod.Key)
	u := "/features/" + mod.Key + "/files"

//line views/vsite/FeatureFiles.html:21
	qw422016.N().S(`
`)
//line views/vsite/FeatureFiles.html:23
	if fs.IsDir(filepath.Join(p.Path...)) {
//line views/vsite/FeatureFiles.html:24
		files := fs.ListFiles(filepath.Join(p.Path...), nil)

//line views/vsite/FeatureFiles.html:24
		qw422016.N().S(`  <div class="card">
    `)
//line views/vsite/FeatureFiles.html:26
		vfile.StreamList(qw422016, p.Path, files, fs, u, as, ps)
//line views/vsite/FeatureFiles.html:26
		qw422016.N().S(`
  </div>
`)
//line views/vsite/FeatureFiles.html:28
	} else {
//line views/vsite/FeatureFiles.html:30
		b, err := fs.ReadFile(filepath.Join(p.Path...))
		if err != nil {
			panic(err)
		}

//line views/vsite/FeatureFiles.html:34
		qw422016.N().S(`  `)
//line views/vsite/FeatureFiles.html:35
		vfile.StreamDetail(qw422016, p.Path, b, u, as, ps)
//line views/vsite/FeatureFiles.html:35
		qw422016.N().S(`
`)
//line views/vsite/FeatureFiles.html:36
	}
//line views/vsite/FeatureFiles.html:37
}

//line views/vsite/FeatureFiles.html:37
func (p *FeatureFiles) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsite/FeatureFiles.html:37
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsite/FeatureFiles.html:37
	p.StreamBody(qw422016, as, ps)
//line views/vsite/FeatureFiles.html:37
	qt422016.ReleaseWriter(qw422016)
//line views/vsite/FeatureFiles.html:37
}

//line views/vsite/FeatureFiles.html:37
func (p *FeatureFiles) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsite/FeatureFiles.html:37
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsite/FeatureFiles.html:37
	p.WriteBody(qb422016, as, ps)
//line views/vsite/FeatureFiles.html:37
	qs422016 := string(qb422016.B)
//line views/vsite/FeatureFiles.html:37
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsite/FeatureFiles.html:37
	return qs422016
//line views/vsite/FeatureFiles.html:37
}
