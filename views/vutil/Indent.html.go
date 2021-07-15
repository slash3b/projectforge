// Code generated by qtc from "Indent.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vutil/Indent.html:1
package vutil

//line views/vutil/Indent.html:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vutil/Indent.html:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vutil/Indent.html:1
func StreamIndent(qw422016 *qt422016.Writer, br bool, level int) {
//line views/vutil/Indent.html:2
	if br {
//line views/vutil/Indent.html:2
		qw422016.N().S(`
`)
//line views/vutil/Indent.html:2
	}
//line views/vutil/Indent.html:3
	for i := 0; i < level; i++ {
//line views/vutil/Indent.html:4
		qw422016.N().S(` `)
//line views/vutil/Indent.html:4
		qw422016.N().S(` `)
//line views/vutil/Indent.html:5
	}
//line views/vutil/Indent.html:6
}

//line views/vutil/Indent.html:6
func WriteIndent(qq422016 qtio422016.Writer, br bool, level int) {
//line views/vutil/Indent.html:6
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vutil/Indent.html:6
	StreamIndent(qw422016, br, level)
//line views/vutil/Indent.html:6
	qt422016.ReleaseWriter(qw422016)
//line views/vutil/Indent.html:6
}

//line views/vutil/Indent.html:6
func Indent(br bool, level int) string {
//line views/vutil/Indent.html:6
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vutil/Indent.html:6
	WriteIndent(qb422016, br, level)
//line views/vutil/Indent.html:6
	qs422016 := string(qb422016.B)
//line views/vutil/Indent.html:6
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vutil/Indent.html:6
	return qs422016
//line views/vutil/Indent.html:6
}
