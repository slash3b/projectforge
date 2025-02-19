package enum

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"

	"projectforge.dev/projectforge/app/util"
)

const defaultIcon = "hammer"

type Enum struct {
	Name           string        `json:"name"`
	Package        string        `json:"package"`
	Group          []string      `json:"group,omitempty"`
	Description    string        `json:"description,omitempty"`
	Icon           string        `json:"icon,omitempty"`
	Values         []string      `json:"values,omitempty"`
	Tags           []string      `json:"tags,omitempty"`
	TitleOverride  string        `json:"title,omitempty"`
	ProperOverride string        `json:"proper,omitempty"`
	RouteOverride  string        `json:"route,omitempty"`
	Config         util.ValueMap `json:"config,omitempty"`
}

func (e *Enum) Title() string {
	if e.TitleOverride == "" {
		return e.Proper()
	}
	return e.TitleOverride
}

func (e *Enum) Proper() string {
	if e.ProperOverride == "" {
		return util.StringToCamel(e.Name)
	}
	return util.StringToCamel(e.ProperOverride)
}

func (e *Enum) IconSafe() string {
	if _, ok := util.SVGLibrary[e.Icon]; ok {
		return e.Icon
	}
	return defaultIcon
}

func (e *Enum) Camel() string {
	return util.StringToLowerCamel(e.Name)
}

func (e *Enum) PackageWithGroup(prefix string) string {
	if len(e.Group) == 0 {
		return prefix + e.Package
	}
	x := make([]string, 0, len(e.Group)+1)
	for _, g := range e.Group {
		x = append(x, prefix+g)
	}
	x = append(x, prefix+e.Package)
	return strings.Join(x, "/")
}

func (e *Enum) HasTag(t string) bool {
	return slices.Contains(e.Tags, t)
}

func (e *Enum) Breadcrumbs() string {
	ret := make([]string, 0, len(e.Group)+1)
	for _, g := range e.Group {
		ret = append(ret, fmt.Sprintf("%q", g))
	}
	ret = append(ret, fmt.Sprintf("%q", e.Package))
	return strings.Join(ret, ", ")
}

type Enums []*Enum

func (e Enums) Get(key string) *Enum {
	for _, x := range e {
		if x.Name == key {
			return x
		}
	}
	return nil
}
