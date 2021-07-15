// Code generated by qtc from "Form.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/components/Form.html:1
package components

//line views/components/Form.html:1
import "github.com/kyleu/projectforge/app/util"

//line views/components/Form.html:2
import "github.com/kyleu/projectforge/views/vutil"

//line views/components/Form.html:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/Form.html:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/Form.html:4
func StreamFormInput(qw422016 *qt422016.Writer, key string, id string, value string) {
//line views/components/Form.html:5
	if id == "" {
//line views/components/Form.html:5
		qw422016.N().S(`<input name="`)
//line views/components/Form.html:6
		qw422016.E().S(key)
//line views/components/Form.html:6
		qw422016.N().S(`" value="`)
//line views/components/Form.html:6
		qw422016.E().S(value)
//line views/components/Form.html:6
		qw422016.N().S(`" />`)
//line views/components/Form.html:7
	} else {
//line views/components/Form.html:7
		qw422016.N().S(`<input id="`)
//line views/components/Form.html:8
		qw422016.E().S(id)
//line views/components/Form.html:8
		qw422016.N().S(`" name="`)
//line views/components/Form.html:8
		qw422016.E().S(key)
//line views/components/Form.html:8
		qw422016.N().S(`" value="`)
//line views/components/Form.html:8
		qw422016.E().S(value)
//line views/components/Form.html:8
		qw422016.N().S(`" />`)
//line views/components/Form.html:9
	}
//line views/components/Form.html:10
}

//line views/components/Form.html:10
func WriteFormInput(qq422016 qtio422016.Writer, key string, id string, value string) {
//line views/components/Form.html:10
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:10
	StreamFormInput(qw422016, key, id, value)
//line views/components/Form.html:10
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:10
}

//line views/components/Form.html:10
func FormInput(key string, id string, value string) string {
//line views/components/Form.html:10
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:10
	WriteFormInput(qb422016, key, id, value)
//line views/components/Form.html:10
	qs422016 := string(qb422016.B)
//line views/components/Form.html:10
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:10
	return qs422016
//line views/components/Form.html:10
}

//line views/components/Form.html:12
func StreamFormInputNumber(qw422016 *qt422016.Writer, key string, id string, value interface{}) {
//line views/components/Form.html:13
	if id == "" {
//line views/components/Form.html:13
		qw422016.N().S(`<input name="`)
//line views/components/Form.html:14
		qw422016.E().S(key)
//line views/components/Form.html:14
		qw422016.N().S(`" type="number" value="`)
//line views/components/Form.html:14
		qw422016.E().V(value)
//line views/components/Form.html:14
		qw422016.N().S(`" />`)
//line views/components/Form.html:15
	} else {
//line views/components/Form.html:15
		qw422016.N().S(`<input id="`)
//line views/components/Form.html:16
		qw422016.E().S(id)
//line views/components/Form.html:16
		qw422016.N().S(`" name="`)
//line views/components/Form.html:16
		qw422016.E().S(key)
//line views/components/Form.html:16
		qw422016.N().S(`" type="number" value="`)
//line views/components/Form.html:16
		qw422016.E().V(value)
//line views/components/Form.html:16
		qw422016.N().S(`" />`)
//line views/components/Form.html:17
	}
//line views/components/Form.html:18
}

//line views/components/Form.html:18
func WriteFormInputNumber(qq422016 qtio422016.Writer, key string, id string, value interface{}) {
//line views/components/Form.html:18
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:18
	StreamFormInputNumber(qw422016, key, id, value)
//line views/components/Form.html:18
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:18
}

//line views/components/Form.html:18
func FormInputNumber(key string, id string, value interface{}) string {
//line views/components/Form.html:18
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:18
	WriteFormInputNumber(qb422016, key, id, value)
//line views/components/Form.html:18
	qs422016 := string(qb422016.B)
//line views/components/Form.html:18
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:18
	return qs422016
//line views/components/Form.html:18
}

//line views/components/Form.html:20
func StreamFormInputTimestamp(qw422016 *qt422016.Writer, key string, id string, value interface{}) {
//line views/components/Form.html:21
	if id == "" {
//line views/components/Form.html:21
		qw422016.N().S(`<input name="`)
//line views/components/Form.html:22
		qw422016.E().S(key)
//line views/components/Form.html:22
		qw422016.N().S(`" type="datetime-local" value="`)
//line views/components/Form.html:22
		qw422016.E().V(value)
//line views/components/Form.html:22
		qw422016.N().S(`" />`)
//line views/components/Form.html:23
	} else {
//line views/components/Form.html:23
		qw422016.N().S(`<input id="`)
//line views/components/Form.html:24
		qw422016.E().S(id)
//line views/components/Form.html:24
		qw422016.N().S(`" name="`)
//line views/components/Form.html:24
		qw422016.E().S(key)
//line views/components/Form.html:24
		qw422016.N().S(`" type="datetime-local" value="`)
//line views/components/Form.html:24
		qw422016.E().V(value)
//line views/components/Form.html:24
		qw422016.N().S(`" />`)
//line views/components/Form.html:25
	}
//line views/components/Form.html:26
}

//line views/components/Form.html:26
func WriteFormInputTimestamp(qq422016 qtio422016.Writer, key string, id string, value interface{}) {
//line views/components/Form.html:26
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:26
	StreamFormInputTimestamp(qw422016, key, id, value)
//line views/components/Form.html:26
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:26
}

//line views/components/Form.html:26
func FormInputTimestamp(key string, id string, value interface{}) string {
//line views/components/Form.html:26
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:26
	WriteFormInputTimestamp(qb422016, key, id, value)
//line views/components/Form.html:26
	qs422016 := string(qb422016.B)
//line views/components/Form.html:26
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:26
	return qs422016
//line views/components/Form.html:26
}

//line views/components/Form.html:28
func StreamFormTextarea(qw422016 *qt422016.Writer, key string, id string, value string) {
//line views/components/Form.html:29
	if id == "" {
//line views/components/Form.html:29
		qw422016.N().S(`<textarea name="`)
//line views/components/Form.html:30
		qw422016.E().S(key)
//line views/components/Form.html:30
		qw422016.N().S(`">`)
//line views/components/Form.html:30
		qw422016.E().S(value)
//line views/components/Form.html:30
		qw422016.N().S(`</textarea>`)
//line views/components/Form.html:31
	} else {
//line views/components/Form.html:31
		qw422016.N().S(`<textarea id="`)
//line views/components/Form.html:32
		qw422016.E().S(id)
//line views/components/Form.html:32
		qw422016.N().S(`" name="`)
//line views/components/Form.html:32
		qw422016.E().S(key)
//line views/components/Form.html:32
		qw422016.N().S(`">`)
//line views/components/Form.html:32
		qw422016.E().S(value)
//line views/components/Form.html:32
		qw422016.N().S(`</textarea>`)
//line views/components/Form.html:33
	}
//line views/components/Form.html:34
}

//line views/components/Form.html:34
func WriteFormTextarea(qq422016 qtio422016.Writer, key string, id string, value string) {
//line views/components/Form.html:34
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:34
	StreamFormTextarea(qw422016, key, id, value)
//line views/components/Form.html:34
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:34
}

//line views/components/Form.html:34
func FormTextarea(key string, id string, value string) string {
//line views/components/Form.html:34
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:34
	WriteFormTextarea(qb422016, key, id, value)
//line views/components/Form.html:34
	qs422016 := string(qb422016.B)
//line views/components/Form.html:34
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:34
	return qs422016
//line views/components/Form.html:34
}

//line views/components/Form.html:36
func StreamFormSelect(qw422016 *qt422016.Writer, key string, id string, value string, opts []string, titles []string, indent int) {
//line views/components/Form.html:37
	if id == "" {
//line views/components/Form.html:37
		qw422016.N().S(`<select name="`)
//line views/components/Form.html:38
		qw422016.E().S(key)
//line views/components/Form.html:38
		qw422016.N().S(`">`)
//line views/components/Form.html:39
	} else {
//line views/components/Form.html:39
		qw422016.N().S(`<select name="`)
//line views/components/Form.html:40
		qw422016.E().S(key)
//line views/components/Form.html:40
		qw422016.N().S(`" id="`)
//line views/components/Form.html:40
		qw422016.E().S(id)
//line views/components/Form.html:40
		qw422016.N().S(`">`)
//line views/components/Form.html:41
	}
//line views/components/Form.html:42
	for idx, opt := range opts {
//line views/components/Form.html:44
		title := opt
		if idx < len(titles) {
			title = titles[idx]
		}

//line views/components/Form.html:49
		vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Form.html:50
		if opt == value {
//line views/components/Form.html:50
			qw422016.N().S(`<option selected="selected" value="`)
//line views/components/Form.html:51
			qw422016.E().S(opt)
//line views/components/Form.html:51
			qw422016.N().S(`">`)
//line views/components/Form.html:51
			qw422016.E().S(title)
//line views/components/Form.html:51
			qw422016.N().S(`</option>`)
//line views/components/Form.html:52
		} else {
//line views/components/Form.html:52
			qw422016.N().S(`<option value="`)
//line views/components/Form.html:53
			qw422016.E().S(opt)
//line views/components/Form.html:53
			qw422016.N().S(`">`)
//line views/components/Form.html:53
			qw422016.E().S(title)
//line views/components/Form.html:53
			qw422016.N().S(`</option>`)
//line views/components/Form.html:54
		}
//line views/components/Form.html:55
	}
//line views/components/Form.html:56
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Form.html:56
	qw422016.N().S(`</select>`)
//line views/components/Form.html:58
}

//line views/components/Form.html:58
func WriteFormSelect(qq422016 qtio422016.Writer, key string, id string, value string, opts []string, titles []string, indent int) {
//line views/components/Form.html:58
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:58
	StreamFormSelect(qw422016, key, id, value, opts, titles, indent)
//line views/components/Form.html:58
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:58
}

//line views/components/Form.html:58
func FormSelect(key string, id string, value string, opts []string, titles []string, indent int) string {
//line views/components/Form.html:58
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:58
	WriteFormSelect(qb422016, key, id, value, opts, titles, indent)
//line views/components/Form.html:58
	qs422016 := string(qb422016.B)
//line views/components/Form.html:58
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:58
	return qs422016
//line views/components/Form.html:58
}

//line views/components/Form.html:60
func StreamFormRadio(qw422016 *qt422016.Writer, key string, value string, opts []string, titles []string, indent int) {
//line views/components/Form.html:61
	for idx, opt := range opts {
//line views/components/Form.html:63
		title := opt
		if idx < len(titles) {
			title = titles[idx]
		}

//line views/components/Form.html:68
		vutil.StreamIndent(qw422016, true, indent)
//line views/components/Form.html:69
		if opt == value {
//line views/components/Form.html:69
			qw422016.N().S(`<label><input type="radio" name="`)
//line views/components/Form.html:70
			qw422016.E().S(key)
//line views/components/Form.html:70
			qw422016.N().S(`" value="`)
//line views/components/Form.html:70
			qw422016.E().S(opt)
//line views/components/Form.html:70
			qw422016.N().S(`" checked="checked" />`)
//line views/components/Form.html:70
			qw422016.N().S(` `)
//line views/components/Form.html:70
			qw422016.E().S(title)
//line views/components/Form.html:70
			qw422016.N().S(`</label>`)
//line views/components/Form.html:71
		} else {
//line views/components/Form.html:71
			qw422016.N().S(`<label><input type="radio" name="`)
//line views/components/Form.html:72
			qw422016.E().S(key)
//line views/components/Form.html:72
			qw422016.N().S(`" value="`)
//line views/components/Form.html:72
			qw422016.E().S(opt)
//line views/components/Form.html:72
			qw422016.N().S(`" />`)
//line views/components/Form.html:72
			qw422016.N().S(` `)
//line views/components/Form.html:72
			qw422016.E().S(title)
//line views/components/Form.html:72
			qw422016.N().S(`</label>`)
//line views/components/Form.html:73
		}
//line views/components/Form.html:74
	}
//line views/components/Form.html:75
}

//line views/components/Form.html:75
func WriteFormRadio(qq422016 qtio422016.Writer, key string, value string, opts []string, titles []string, indent int) {
//line views/components/Form.html:75
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:75
	StreamFormRadio(qw422016, key, value, opts, titles, indent)
//line views/components/Form.html:75
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:75
}

//line views/components/Form.html:75
func FormRadio(key string, value string, opts []string, titles []string, indent int) string {
//line views/components/Form.html:75
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:75
	WriteFormRadio(qb422016, key, value, opts, titles, indent)
//line views/components/Form.html:75
	qs422016 := string(qb422016.B)
//line views/components/Form.html:75
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:75
	return qs422016
//line views/components/Form.html:75
}

//line views/components/Form.html:77
func StreamFormCheckbox(qw422016 *qt422016.Writer, key string, values []string, opts []string, titles []string, indent int) {
//line views/components/Form.html:78
	for idx, opt := range opts {
//line views/components/Form.html:80
		title := opt
		if idx < len(titles) {
			title = titles[idx]
		}

//line views/components/Form.html:85
		vutil.StreamIndent(qw422016, true, indent)
//line views/components/Form.html:86
		if util.StringArrayContains(values, opt) {
//line views/components/Form.html:86
			qw422016.N().S(`<label><input type="checkbox" name="`)
//line views/components/Form.html:87
			qw422016.E().S(key)
//line views/components/Form.html:87
			qw422016.N().S(`" value="`)
//line views/components/Form.html:87
			qw422016.E().S(opt)
//line views/components/Form.html:87
			qw422016.N().S(`" checked="checked" />`)
//line views/components/Form.html:87
			qw422016.N().S(` `)
//line views/components/Form.html:87
			qw422016.E().S(title)
//line views/components/Form.html:87
			qw422016.N().S(`</label>`)
//line views/components/Form.html:88
		} else {
//line views/components/Form.html:88
			qw422016.N().S(`<label><input type="checkbox" name="`)
//line views/components/Form.html:89
			qw422016.E().S(key)
//line views/components/Form.html:89
			qw422016.N().S(`" value="`)
//line views/components/Form.html:89
			qw422016.E().S(opt)
//line views/components/Form.html:89
			qw422016.N().S(`" />`)
//line views/components/Form.html:89
			qw422016.N().S(` `)
//line views/components/Form.html:89
			qw422016.E().S(title)
//line views/components/Form.html:89
			qw422016.N().S(`</label>`)
//line views/components/Form.html:90
		}
//line views/components/Form.html:91
	}
//line views/components/Form.html:92
}

//line views/components/Form.html:92
func WriteFormCheckbox(qq422016 qtio422016.Writer, key string, values []string, opts []string, titles []string, indent int) {
//line views/components/Form.html:92
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:92
	StreamFormCheckbox(qw422016, key, values, opts, titles, indent)
//line views/components/Form.html:92
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:92
}

//line views/components/Form.html:92
func FormCheckbox(key string, values []string, opts []string, titles []string, indent int) string {
//line views/components/Form.html:92
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:92
	WriteFormCheckbox(qb422016, key, values, opts, titles, indent)
//line views/components/Form.html:92
	qs422016 := string(qb422016.B)
//line views/components/Form.html:92
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:92
	return qs422016
//line views/components/Form.html:92
}
