// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vfile/List.html:2
package vfile

//line views/vfile/List.html:2
import (
	"os"
	"path/filepath"

	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/app/lib/filesystem"
	"projectforge.dev/projectforge/views/components"
)

//line views/vfile/List.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vfile/List.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vfile/List.html:12
func StreamList(qw422016 *qt422016.Writer, path []string, files []os.DirEntry, fs filesystem.FileLoader, urlPrefix string, as *app.State, ps *cutil.PageState) {
//line views/vfile/List.html:12
	qw422016.N().S(`
  <h3><a href="`)
//line views/vfile/List.html:13
	qw422016.E().S(urlPrefix)
//line views/vfile/List.html:13
	qw422016.N().S(`">.</a>`)
//line views/vfile/List.html:13
	for idx, p := range path {
//line views/vfile/List.html:13
		qw422016.N().S(`/<a href="`)
//line views/vfile/List.html:13
		qw422016.E().S(urlPrefix)
//line views/vfile/List.html:13
		qw422016.N().S(`/`)
//line views/vfile/List.html:13
		qw422016.E().S(filepath.Join(path[:idx+1]...))
//line views/vfile/List.html:13
		qw422016.N().S(`">`)
//line views/vfile/List.html:13
		qw422016.E().S(p)
//line views/vfile/List.html:13
		qw422016.N().S(`</a>`)
//line views/vfile/List.html:13
	}
//line views/vfile/List.html:13
	qw422016.N().S(`</h3>
  <div class="mt">
`)
//line views/vfile/List.html:15
	for _, f := range files {
//line views/vfile/List.html:17
		icon := "file-text"
		if f.IsDir() {
			icon = "folder"
		}
		x := []string{urlPrefix}
		x = append(x, path...)
		x = append(x, f.Name())
		u := filepath.Join(x...)

//line views/vfile/List.html:25
		qw422016.N().S(`    <div><a href="`)
//line views/vfile/List.html:26
		qw422016.E().S(u)
//line views/vfile/List.html:26
		qw422016.N().S(`">`)
//line views/vfile/List.html:26
		components.StreamSVGRef(qw422016, icon, 16, 16, `icon`, ps)
//line views/vfile/List.html:26
		qw422016.E().S(f.Name())
//line views/vfile/List.html:26
		qw422016.N().S(`</a></div>
`)
//line views/vfile/List.html:27
	}
//line views/vfile/List.html:27
	qw422016.N().S(`  </div>
`)
//line views/vfile/List.html:29
}

//line views/vfile/List.html:29
func WriteList(qq422016 qtio422016.Writer, path []string, files []os.DirEntry, fs filesystem.FileLoader, urlPrefix string, as *app.State, ps *cutil.PageState) {
//line views/vfile/List.html:29
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vfile/List.html:29
	StreamList(qw422016, path, files, fs, urlPrefix, as, ps)
//line views/vfile/List.html:29
	qt422016.ReleaseWriter(qw422016)
//line views/vfile/List.html:29
}

//line views/vfile/List.html:29
func List(path []string, files []os.DirEntry, fs filesystem.FileLoader, urlPrefix string, as *app.State, ps *cutil.PageState) string {
//line views/vfile/List.html:29
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vfile/List.html:29
	WriteList(qb422016, path, files, fs, urlPrefix, as, ps)
//line views/vfile/List.html:29
	qs422016 := string(qb422016.B)
//line views/vfile/List.html:29
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vfile/List.html:29
	return qs422016
//line views/vfile/List.html:29
}
