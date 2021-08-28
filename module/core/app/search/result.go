package search

import (
	"sort"
)

type Result struct {
	ID    string      `json:"id"`
	Type  string      `json:"type"`
	Title string      `json:"title"`
	URL   string      `json:"url"`
	Match string      `json:"match"`
	Data  interface{} `json:"data"`
	HTML  string      `json:"-"`
}

type Results []*Result

func (rs Results) Sort() {
	sort.Slice(rs, func(i, j int) bool {
		l, r := rs[i], rs[j]
		if l.Type == r.Type {
			return l.Title < r.Title
		}
		return l.Type < r.Type
	})
}
