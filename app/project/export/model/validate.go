package model

import (
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
)

var goKeywords = []string{
	"break", "case", "chan", "const", "continue", "default", "defer", "else",
	"fallthrough", "for", "func", "go", "goto", "if", "import", "interface",
	"map", "package", "range", "return", "select", "struct", "switch", "type", "var",
}

var reservedNames = map[string][]string{"audit": {"audit", "audit_record"}}

//nolint:gocognit
func (m *Model) Validate(mods []string, models Models, groups Groups) error {
	if len(m.PKs()) == 0 {
		return errors.Errorf("model [%s] has no primary key", m.Name)
	}
	if m.Package == "vendor" {
		return errors.Errorf("model [%s] uses [vendor] package, which is reserved by Go", m.Name)
	}
	for _, mod := range mods {
		if slices.Contains(reservedNames[mod], m.Name) {
			return errors.Errorf("model [%s] uses name which is reserved by [%s]", m.Name, mod)
		}
	}
	if m.IsRevision() {
		hc := m.HistoryColumns(true)
		if hc.Err != nil {
			return hc.Err
		}
		hc = m.HistoryColumns(false)
		if hc.Err != nil {
			return hc.Err
		}
	}
	if m.IsSoftDelete() {
		if d := m.Columns.WithTag("deleted"); len(d) != 1 {
			return errors.New("when set to soft delete, model must have one column tagged [deleted]")
		}
	}
	for _, col := range m.Columns {
		if slices.Contains(goKeywords, col.Name) {
			return errors.Errorf("column [%s] uses reserved keyword", col.Name)
		}
	}
	for _, rel := range m.Relations {
		for _, s := range rel.Src {
			if m.Columns.Get(s) == nil {
				return errors.Errorf("relation [%s] references missing source column [%s]", rel.Name, s)
			}
		}
	}
	if len(m.Group) > 0 {
		if groups.Get(m.Group...) == nil {
			if len(m.Group) != 1 && models.Get(m.Group[0]) == nil {
				return errors.Errorf("model [%s] references undefined group [%s]", m.Name, strings.Join(m.Group, "/"))
			}
		}
	}
	return nil
}
