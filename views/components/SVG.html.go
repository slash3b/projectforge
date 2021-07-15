// Code generated by qtc from "SVG.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/components/SVG.html:1
package components

//line views/components/SVG.html:1
import "fmt"

//line views/components/SVG.html:2
import "github.com/kyleu/projectforge/app/controller/cutil"

//line views/components/SVG.html:3
import "github.com/kyleu/projectforge/app/util"

//line views/components/SVG.html:5
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/SVG.html:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/SVG.html:5
func StreamSVG(qw422016 *qt422016.Writer, k string) {
//line views/components/SVG.html:5
	qw422016.N().S(util.SVGLibrary[k])
//line views/components/SVG.html:5
}

//line views/components/SVG.html:5
func WriteSVG(qq422016 qtio422016.Writer, k string) {
//line views/components/SVG.html:5
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/SVG.html:5
	StreamSVG(qw422016, k)
//line views/components/SVG.html:5
	qt422016.ReleaseWriter(qw422016)
//line views/components/SVG.html:5
}

//line views/components/SVG.html:5
func SVG(k string) string {
//line views/components/SVG.html:5
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/SVG.html:5
	WriteSVG(qb422016, k)
//line views/components/SVG.html:5
	qs422016 := string(qb422016.B)
//line views/components/SVG.html:5
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/SVG.html:5
	return qs422016
//line views/components/SVG.html:5
}

//line views/components/SVG.html:7
func StreamSVGRef(qw422016 *qt422016.Writer, k string, w int, h int, cls string, ps *cutil.PageState) {
//line views/components/SVG.html:8
	if k != "" {
//line views/components/SVG.html:10
		ps.AddIcon(k)
		if w == 0 {
			w = 20
		}
		if h == 0 {
			h = 20
		}
		style := fmt.Sprintf("width: %dpx; height: %dpx;", w, h)

//line views/components/SVG.html:15
		if cls == "" {
//line views/components/SVG.html:15
			qw422016.N().S(`<svg style="`)
//line views/components/SVG.html:16
			qw422016.E().S(style)
//line views/components/SVG.html:16
			qw422016.N().S(`"><use xlink:href="#svg-`)
//line views/components/SVG.html:16
			qw422016.E().S(k)
//line views/components/SVG.html:16
			qw422016.N().S(`" /></svg>`)
//line views/components/SVG.html:17
		} else {
//line views/components/SVG.html:17
			qw422016.N().S(`<svg class="`)
//line views/components/SVG.html:18
			qw422016.E().S(cls)
//line views/components/SVG.html:18
			qw422016.N().S(`" style="`)
//line views/components/SVG.html:18
			qw422016.E().S(style)
//line views/components/SVG.html:18
			qw422016.N().S(`"><use xlink:href="#svg-`)
//line views/components/SVG.html:18
			qw422016.E().S(k)
//line views/components/SVG.html:18
			qw422016.N().S(`" /></svg>`)
//line views/components/SVG.html:19
		}
//line views/components/SVG.html:20
	}
//line views/components/SVG.html:21
}

//line views/components/SVG.html:21
func WriteSVGRef(qq422016 qtio422016.Writer, k string, w int, h int, cls string, ps *cutil.PageState) {
//line views/components/SVG.html:21
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/SVG.html:21
	StreamSVGRef(qw422016, k, w, h, cls, ps)
//line views/components/SVG.html:21
	qt422016.ReleaseWriter(qw422016)
//line views/components/SVG.html:21
}

//line views/components/SVG.html:21
func SVGRef(k string, w int, h int, cls string, ps *cutil.PageState) string {
//line views/components/SVG.html:21
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/SVG.html:21
	WriteSVGRef(qb422016, k, w, h, cls, ps)
//line views/components/SVG.html:21
	qs422016 := string(qb422016.B)
//line views/components/SVG.html:21
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/SVG.html:21
	return qs422016
//line views/components/SVG.html:21
}
