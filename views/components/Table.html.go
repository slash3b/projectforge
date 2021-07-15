// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/components/Table.html:1
package components

//line views/components/Table.html:1
import "fmt"

//line views/components/Table.html:2
import "github.com/valyala/fasthttp"

//line views/components/Table.html:3
import "github.com/kyleu/projectforge/app/controller/cutil"

//line views/components/Table.html:4
import "github.com/kyleu/projectforge/app/filter"

//line views/components/Table.html:5
import "github.com/kyleu/projectforge/views/vutil"

//line views/components/Table.html:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/Table.html:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/Table.html:7
func StreamTableHeader(qw422016 *qt422016.Writer, section string, key string, title string, locs cutil.Locations, params *filter.Params, icon string, u *fasthttp.URI, tooltip string, sortable bool, cls string, ps *cutil.PageState) {
//line views/components/Table.html:8
	if cls != "" {
//line views/components/Table.html:8
		qw422016.N().S(`<th class="`)
//line views/components/Table.html:9
		qw422016.E().S(cls)
//line views/components/Table.html:9
		qw422016.N().S(` `)
//line views/components/Table.html:9
		qw422016.N().S(`no-padding" scope="col">`)
//line views/components/Table.html:10
	} else {
//line views/components/Table.html:10
		qw422016.N().S(`<th class="no-padding" scope="col">`)
//line views/components/Table.html:12
	}
//line views/components/Table.html:12
	qw422016.N().S(`<div class="resize">`)
//line views/components/Table.html:14
	if !sortable {
//line views/components/Table.html:14
		qw422016.N().S(`<div title="`)
//line views/components/Table.html:15
		qw422016.E().S(tooltip)
//line views/components/Table.html:15
		qw422016.N().S(`">`)
//line views/components/Table.html:16
		if icon != "" {
//line views/components/Table.html:17
			qw422016.N().S(` `)
//line views/components/Table.html:18
			StreamSVGRef(qw422016, icon, 16, 16, "icon-block", ps)
//line views/components/Table.html:19
		}
//line views/components/Table.html:20
		qw422016.E().S(title)
//line views/components/Table.html:20
		qw422016.N().S(`</div>`)
//line views/components/Table.html:22
	} else if params == nil {
//line views/components/Table.html:23
		streamthNormal(qw422016, section, key, title, params, icon, u, tooltip, ps)
//line views/components/Table.html:24
	} else {
//line views/components/Table.html:25
		o := params.GetOrdering(key)

//line views/components/Table.html:26
		if o == nil {
//line views/components/Table.html:27
			streamthNormal(qw422016, section, key, title, params, icon, u, tooltip, ps)
//line views/components/Table.html:28
		} else {
//line views/components/Table.html:29
			streamthSorted(qw422016, o.Asc, section, key, title, params, icon, u, tooltip, ps)
//line views/components/Table.html:30
		}
//line views/components/Table.html:31
	}
//line views/components/Table.html:31
	qw422016.N().S(`</div></th>`)
//line views/components/Table.html:33
}

//line views/components/Table.html:33
func WriteTableHeader(qq422016 qtio422016.Writer, section string, key string, title string, locs cutil.Locations, params *filter.Params, icon string, u *fasthttp.URI, tooltip string, sortable bool, cls string, ps *cutil.PageState) {
//line views/components/Table.html:33
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:33
	StreamTableHeader(qw422016, section, key, title, locs, params, icon, u, tooltip, sortable, cls, ps)
//line views/components/Table.html:33
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:33
}

//line views/components/Table.html:33
func TableHeader(section string, key string, title string, locs cutil.Locations, params *filter.Params, icon string, u *fasthttp.URI, tooltip string, sortable bool, cls string, ps *cutil.PageState) string {
//line views/components/Table.html:33
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:33
	WriteTableHeader(qb422016, section, key, title, locs, params, icon, u, tooltip, sortable, cls, ps)
//line views/components/Table.html:33
	qs422016 := string(qb422016.B)
//line views/components/Table.html:33
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:33
	return qs422016
//line views/components/Table.html:33
}

//line views/components/Table.html:35
func StreamTableInput(qw422016 *qt422016.Writer, key string, title string, value string, indent int) {
//line views/components/Table.html:35
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:37
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:37
	qw422016.N().S(`<th class="shrink"><label for="input-`)
//line views/components/Table.html:38
	qw422016.E().S(key)
//line views/components/Table.html:38
	qw422016.N().S(`">`)
//line views/components/Table.html:38
	qw422016.E().S(title)
//line views/components/Table.html:38
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:39
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:39
	qw422016.N().S(`<td>`)
//line views/components/Table.html:40
	StreamFormInput(qw422016, key, "input-"+key, value)
//line views/components/Table.html:40
	qw422016.N().S(`</td>`)
//line views/components/Table.html:41
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:41
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:43
}

//line views/components/Table.html:43
func WriteTableInput(qq422016 qtio422016.Writer, key string, title string, value string, indent int) {
//line views/components/Table.html:43
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:43
	StreamTableInput(qw422016, key, title, value, indent)
//line views/components/Table.html:43
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:43
}

//line views/components/Table.html:43
func TableInput(key string, title string, value string, indent int) string {
//line views/components/Table.html:43
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:43
	WriteTableInput(qb422016, key, title, value, indent)
//line views/components/Table.html:43
	qs422016 := string(qb422016.B)
//line views/components/Table.html:43
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:43
	return qs422016
//line views/components/Table.html:43
}

//line views/components/Table.html:45
func StreamTableInputNumber(qw422016 *qt422016.Writer, key string, title string, value int, indent int) {
//line views/components/Table.html:45
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:47
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:47
	qw422016.N().S(`<th class="shrink"><label for="input-`)
//line views/components/Table.html:48
	qw422016.E().S(key)
//line views/components/Table.html:48
	qw422016.N().S(`">`)
//line views/components/Table.html:48
	qw422016.E().S(title)
//line views/components/Table.html:48
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:49
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:49
	qw422016.N().S(`<td>`)
//line views/components/Table.html:50
	StreamFormInputNumber(qw422016, key, "input-"+key, value)
//line views/components/Table.html:50
	qw422016.N().S(`</td>`)
//line views/components/Table.html:51
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:51
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:53
}

//line views/components/Table.html:53
func WriteTableInputNumber(qq422016 qtio422016.Writer, key string, title string, value int, indent int) {
//line views/components/Table.html:53
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:53
	StreamTableInputNumber(qw422016, key, title, value, indent)
//line views/components/Table.html:53
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:53
}

//line views/components/Table.html:53
func TableInputNumber(key string, title string, value int, indent int) string {
//line views/components/Table.html:53
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:53
	WriteTableInputNumber(qb422016, key, title, value, indent)
//line views/components/Table.html:53
	qs422016 := string(qb422016.B)
//line views/components/Table.html:53
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:53
	return qs422016
//line views/components/Table.html:53
}

//line views/components/Table.html:55
func StreamTableTextarea(qw422016 *qt422016.Writer, key string, title string, value string, indent int) {
//line views/components/Table.html:55
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:57
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:57
	qw422016.N().S(`<th class="shrink"><label for="input-`)
//line views/components/Table.html:58
	qw422016.E().S(key)
//line views/components/Table.html:58
	qw422016.N().S(`">`)
//line views/components/Table.html:58
	qw422016.E().S(title)
//line views/components/Table.html:58
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:59
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:59
	qw422016.N().S(`<td>`)
//line views/components/Table.html:60
	StreamFormTextarea(qw422016, key, "input-"+key, value)
//line views/components/Table.html:60
	qw422016.N().S(`</td>`)
//line views/components/Table.html:61
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:61
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:63
}

//line views/components/Table.html:63
func WriteTableTextarea(qq422016 qtio422016.Writer, key string, title string, value string, indent int) {
//line views/components/Table.html:63
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:63
	StreamTableTextarea(qw422016, key, title, value, indent)
//line views/components/Table.html:63
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:63
}

//line views/components/Table.html:63
func TableTextarea(key string, title string, value string, indent int) string {
//line views/components/Table.html:63
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:63
	WriteTableTextarea(qb422016, key, title, value, indent)
//line views/components/Table.html:63
	qs422016 := string(qb422016.B)
//line views/components/Table.html:63
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:63
	return qs422016
//line views/components/Table.html:63
}

//line views/components/Table.html:65
func StreamTableSelect(qw422016 *qt422016.Writer, key string, title string, value string, opts []string, titles []string, indent int) {
//line views/components/Table.html:65
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:67
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:67
	qw422016.N().S(`<th class="shrink"><label for="input-`)
//line views/components/Table.html:68
	qw422016.E().S(key)
//line views/components/Table.html:68
	qw422016.N().S(`">`)
//line views/components/Table.html:68
	qw422016.E().S(title)
//line views/components/Table.html:68
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:69
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:69
	qw422016.N().S(`<td>`)
//line views/components/Table.html:70
	StreamFormSelect(qw422016, key, "input-"+key, value, opts, titles, indent)
//line views/components/Table.html:70
	qw422016.N().S(`</td>`)
//line views/components/Table.html:71
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:71
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:73
}

//line views/components/Table.html:73
func WriteTableSelect(qq422016 qtio422016.Writer, key string, title string, value string, opts []string, titles []string, indent int) {
//line views/components/Table.html:73
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:73
	StreamTableSelect(qw422016, key, title, value, opts, titles, indent)
//line views/components/Table.html:73
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:73
}

//line views/components/Table.html:73
func TableSelect(key string, title string, value string, opts []string, titles []string, indent int) string {
//line views/components/Table.html:73
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:73
	WriteTableSelect(qb422016, key, title, value, opts, titles, indent)
//line views/components/Table.html:73
	qs422016 := string(qb422016.B)
//line views/components/Table.html:73
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:73
	return qs422016
//line views/components/Table.html:73
}

//line views/components/Table.html:75
func StreamTableRadio(qw422016 *qt422016.Writer, key string, title string, value string, opts []string, titles []string, indent int) {
//line views/components/Table.html:75
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:77
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:77
	qw422016.N().S(`<th class="shrink"><label>`)
//line views/components/Table.html:78
	qw422016.E().S(title)
//line views/components/Table.html:78
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:79
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:79
	qw422016.N().S(`<td>`)
//line views/components/Table.html:81
	StreamFormRadio(qw422016, key, value, opts, titles, indent+2)
//line views/components/Table.html:82
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:82
	qw422016.N().S(`</td>`)
//line views/components/Table.html:84
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:84
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:86
}

//line views/components/Table.html:86
func WriteTableRadio(qq422016 qtio422016.Writer, key string, title string, value string, opts []string, titles []string, indent int) {
//line views/components/Table.html:86
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:86
	StreamTableRadio(qw422016, key, title, value, opts, titles, indent)
//line views/components/Table.html:86
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:86
}

//line views/components/Table.html:86
func TableRadio(key string, title string, value string, opts []string, titles []string, indent int) string {
//line views/components/Table.html:86
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:86
	WriteTableRadio(qb422016, key, title, value, opts, titles, indent)
//line views/components/Table.html:86
	qs422016 := string(qb422016.B)
//line views/components/Table.html:86
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:86
	return qs422016
//line views/components/Table.html:86
}

//line views/components/Table.html:88
func StreamTableBoolean(qw422016 *qt422016.Writer, key string, title string, value bool, indent int) {
//line views/components/Table.html:89
	StreamTableRadio(qw422016, key, title, fmt.Sprintf("%v", value), []string{"true", "false"}, []string{"True", "False"}, indent)
//line views/components/Table.html:90
}

//line views/components/Table.html:90
func WriteTableBoolean(qq422016 qtio422016.Writer, key string, title string, value bool, indent int) {
//line views/components/Table.html:90
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:90
	StreamTableBoolean(qw422016, key, title, value, indent)
//line views/components/Table.html:90
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:90
}

//line views/components/Table.html:90
func TableBoolean(key string, title string, value bool, indent int) string {
//line views/components/Table.html:90
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:90
	WriteTableBoolean(qb422016, key, title, value, indent)
//line views/components/Table.html:90
	qs422016 := string(qb422016.B)
//line views/components/Table.html:90
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:90
	return qs422016
//line views/components/Table.html:90
}

//line views/components/Table.html:92
func StreamTableCheckbox(qw422016 *qt422016.Writer, key string, title string, values []string, opts []string, titles []string, indent int) {
//line views/components/Table.html:92
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:94
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:94
	qw422016.N().S(`<th class="shrink"><label>`)
//line views/components/Table.html:95
	qw422016.E().S(title)
//line views/components/Table.html:95
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:96
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:96
	qw422016.N().S(`<td>`)
//line views/components/Table.html:98
	StreamFormCheckbox(qw422016, key, values, opts, titles, indent+2)
//line views/components/Table.html:99
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:99
	qw422016.N().S(`</td>`)
//line views/components/Table.html:101
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:101
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:103
}

//line views/components/Table.html:103
func WriteTableCheckbox(qq422016 qtio422016.Writer, key string, title string, values []string, opts []string, titles []string, indent int) {
//line views/components/Table.html:103
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:103
	StreamTableCheckbox(qw422016, key, title, values, opts, titles, indent)
//line views/components/Table.html:103
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:103
}

//line views/components/Table.html:103
func TableCheckbox(key string, title string, values []string, opts []string, titles []string, indent int) string {
//line views/components/Table.html:103
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:103
	WriteTableCheckbox(qb422016, key, title, values, opts, titles, indent)
//line views/components/Table.html:103
	qs422016 := string(qb422016.B)
//line views/components/Table.html:103
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:103
	return qs422016
//line views/components/Table.html:103
}

//line views/components/Table.html:105
func StreamTableIcons(qw422016 *qt422016.Writer, key string, title string, value string, ps *cutil.PageState, indent int) {
//line views/components/Table.html:105
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:107
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:107
	qw422016.N().S(`<th class="shrink"><label>`)
//line views/components/Table.html:108
	qw422016.E().S(title)
//line views/components/Table.html:108
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:109
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:109
	qw422016.N().S(`<td>`)
//line views/components/Table.html:110
	StreamIconPicker(qw422016, value, key, ps, indent+2)
//line views/components/Table.html:110
	qw422016.N().S(`</td>`)
//line views/components/Table.html:112
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:112
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:114
}

//line views/components/Table.html:114
func WriteTableIcons(qq422016 qtio422016.Writer, key string, title string, value string, ps *cutil.PageState, indent int) {
//line views/components/Table.html:114
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:114
	StreamTableIcons(qw422016, key, title, value, ps, indent)
//line views/components/Table.html:114
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:114
}

//line views/components/Table.html:114
func TableIcons(key string, title string, value string, ps *cutil.PageState, indent int) string {
//line views/components/Table.html:114
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:114
	WriteTableIcons(qb422016, key, title, value, ps, indent)
//line views/components/Table.html:114
	qs422016 := string(qb422016.B)
//line views/components/Table.html:114
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:114
	return qs422016
//line views/components/Table.html:114
}

//line views/components/Table.html:116
func streamthNormal(qw422016 *qt422016.Writer, section string, key string, title string, params *filter.Params, icon string, u *fasthttp.URI, tooltip string, ps *cutil.PageState) {
//line views/components/Table.html:116
	qw422016.N().S(`<a class="sort-hover" href="?`)
//line views/components/Table.html:117
	qw422016.N().S(params.CloneOrdering(&filter.Ordering{Column: key, Asc: true}).ToQueryString(u))
//line views/components/Table.html:117
	qw422016.N().S(`" title="`)
//line views/components/Table.html:117
	qw422016.E().S(tooltip)
//line views/components/Table.html:117
	qw422016.N().S(`"><div class="sort-icon" title="click to sort by this column, ascending">`)
//line views/components/Table.html:118
	StreamSVGRef(qw422016, `down`, 0, 0, ``, ps)
//line views/components/Table.html:118
	qw422016.N().S(`</div><div class="sort-title">`)
//line views/components/Table.html:120
	if icon != "" {
//line views/components/Table.html:121
		qw422016.N().S(` `)
//line views/components/Table.html:122
		StreamSVGRef(qw422016, icon, 16, 16, "icon-block", ps)
//line views/components/Table.html:123
	}
//line views/components/Table.html:124
	qw422016.E().S(title)
//line views/components/Table.html:124
	qw422016.N().S(`</div></a>`)
//line views/components/Table.html:127
}

//line views/components/Table.html:127
func writethNormal(qq422016 qtio422016.Writer, section string, key string, title string, params *filter.Params, icon string, u *fasthttp.URI, tooltip string, ps *cutil.PageState) {
//line views/components/Table.html:127
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:127
	streamthNormal(qw422016, section, key, title, params, icon, u, tooltip, ps)
//line views/components/Table.html:127
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:127
}

//line views/components/Table.html:127
func thNormal(section string, key string, title string, params *filter.Params, icon string, u *fasthttp.URI, tooltip string, ps *cutil.PageState) string {
//line views/components/Table.html:127
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:127
	writethNormal(qb422016, section, key, title, params, icon, u, tooltip, ps)
//line views/components/Table.html:127
	qs422016 := string(qb422016.B)
//line views/components/Table.html:127
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:127
	return qs422016
//line views/components/Table.html:127
}

//line views/components/Table.html:129
func streamthSorted(qw422016 *qt422016.Writer, asc bool, section string, key string, title string, params *filter.Params, icon string, u *fasthttp.URI, tooltip string, ps *cutil.PageState) {
//line views/components/Table.html:131
	ascStr := "ascending"
	dirStr := "up"
	if asc {
		ascStr = "descending"
		dirStr = "down"
	}

//line views/components/Table.html:137
	qw422016.N().S(`<a href="?`)
//line views/components/Table.html:138
	qw422016.N().S(params.CloneOrdering(&filter.Ordering{Column: key, Asc: !asc}).ToQueryString(u))
//line views/components/Table.html:138
	qw422016.N().S(`" title="`)
//line views/components/Table.html:138
	qw422016.E().S(tooltip)
//line views/components/Table.html:138
	qw422016.N().S(`"><div class="sort-icon" title="click to sort by this column,`)
//line views/components/Table.html:139
	qw422016.N().S(` `)
//line views/components/Table.html:139
	qw422016.E().S(ascStr)
//line views/components/Table.html:139
	qw422016.N().S(`">`)
//line views/components/Table.html:139
	StreamSVGRef(qw422016, dirStr, 0, 0, ``, ps)
//line views/components/Table.html:139
	qw422016.N().S(`</div><div class="sort-title">`)
//line views/components/Table.html:141
	if icon != "" {
//line views/components/Table.html:142
		qw422016.N().S(` `)
//line views/components/Table.html:143
		StreamSVGRef(qw422016, icon, 16, 16, "icon-block", ps)
//line views/components/Table.html:144
	}
//line views/components/Table.html:145
	qw422016.E().S(title)
//line views/components/Table.html:145
	qw422016.N().S(`</div></a>`)
//line views/components/Table.html:148
}

//line views/components/Table.html:148
func writethSorted(qq422016 qtio422016.Writer, asc bool, section string, key string, title string, params *filter.Params, icon string, u *fasthttp.URI, tooltip string, ps *cutil.PageState) {
//line views/components/Table.html:148
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:148
	streamthSorted(qw422016, asc, section, key, title, params, icon, u, tooltip, ps)
//line views/components/Table.html:148
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:148
}

//line views/components/Table.html:148
func thSorted(asc bool, section string, key string, title string, params *filter.Params, icon string, u *fasthttp.URI, tooltip string, ps *cutil.PageState) string {
//line views/components/Table.html:148
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:148
	writethSorted(qb422016, asc, section, key, title, params, icon, u, tooltip, ps)
//line views/components/Table.html:148
	qs422016 := string(qb422016.B)
//line views/components/Table.html:148
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:148
	return qs422016
//line views/components/Table.html:148
}
