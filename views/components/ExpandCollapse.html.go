// Code generated by qtc from "ExpandCollapse.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/components/ExpandCollapse.html:1
package components

//line views/components/ExpandCollapse.html:1
import "github.com/kyleu/projectforge/app/controller/cutil"

//line views/components/ExpandCollapse.html:2
import "github.com/kyleu/projectforge/views/vutil"

//line views/components/ExpandCollapse.html:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/ExpandCollapse.html:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/ExpandCollapse.html:4
func StreamExpandCollapse(qw422016 *qt422016.Writer, indent int, ps *cutil.PageState) {
//line views/components/ExpandCollapse.html:5
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/ExpandCollapse.html:6
	StreamSVGRef(qw422016, `right`, 15, 15, `expand`, ps)
//line views/components/ExpandCollapse.html:7
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/ExpandCollapse.html:8
	StreamSVGRef(qw422016, `down`, 15, 15, `collapse`, ps)
//line views/components/ExpandCollapse.html:9
}

//line views/components/ExpandCollapse.html:9
func WriteExpandCollapse(qq422016 qtio422016.Writer, indent int, ps *cutil.PageState) {
//line views/components/ExpandCollapse.html:9
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/ExpandCollapse.html:9
	StreamExpandCollapse(qw422016, indent, ps)
//line views/components/ExpandCollapse.html:9
	qt422016.ReleaseWriter(qw422016)
//line views/components/ExpandCollapse.html:9
}

//line views/components/ExpandCollapse.html:9
func ExpandCollapse(indent int, ps *cutil.PageState) string {
//line views/components/ExpandCollapse.html:9
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/ExpandCollapse.html:9
	WriteExpandCollapse(qb422016, indent, ps)
//line views/components/ExpandCollapse.html:9
	qs422016 := string(qb422016.B)
//line views/components/ExpandCollapse.html:9
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/ExpandCollapse.html:9
	return qs422016
//line views/components/ExpandCollapse.html:9
}
