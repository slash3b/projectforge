package gomodel

import (
	"fmt"
	"strings"

	"projectforge.dev/projectforge/app/file"
	"projectforge.dev/projectforge/app/project/export/enum"
	"projectforge.dev/projectforge/app/project/export/files/helper"
	"projectforge.dev/projectforge/app/project/export/golang"
	"projectforge.dev/projectforge/app/project/export/model"
	"projectforge.dev/projectforge/app/util"
)

func History(m *model.Model, args *model.Args, addHeader bool) (*file.File, error) {
	g := golang.NewFile(m.Package, []string{"app", m.PackageWithGroup("")}, m.Camel()+"history")
	for _, imp := range helper.ImportsForTypes("go", m.Columns.Types()...) {
		g.AddImport(imp)
	}
	g.AddImport(helper.ImpJSON, helper.ImpUUID, helper.ImpAppUtil)
	mh, err := modelHistory(m, args.Enums)
	if err != nil {
		return nil, err
	}
	dto, err := modelHistoryDTO(m, args.Enums)
	if err != nil {
		return nil, err
	}
	g.AddBlocks(mh, modelHistoryToData(m), modelHistories(m), dto, modelHistoryDTOToHistory(m), modelHistoryDTOs(m))
	return g.Render(addHeader)
}

func modelHistory(m *model.Model, enums enum.Enums) (*golang.Block, error) {
	ret := golang.NewBlock(m.Proper()+"History", "struct")
	ret.W("type History struct {")
	max := m.PKs().MaxCamelLength() + len(m.Camel())
	tmax := 13
	ret.W("\t%s %s `json:\"id\"`", util.StringPad("ID", max), util.StringPad("uuid.UUID", tmax))
	for _, pk := range m.PKs() {
		gt, err := pk.ToGoType(m.Package, enums)
		if err != nil {
			return nil, err
		}
		goType := util.StringPad(gt, tmax)
		ret.W("\t%s %s `json:\"%s%s\"`", util.StringPad(m.Proper()+pk.Proper(), max), goType, m.Camel(), pk.Proper())
	}
	ret.W("\t%s util.ValueMap `json:\"o,omitempty\"`", util.StringPad("Old", max))
	ret.W("\t%s util.ValueMap `json:\"n,omitempty\"`", util.StringPad("New", max))
	ret.W("\t%s util.Diffs    `json:\"c,omitempty\"`", util.StringPad("Changes", max))
	ret.W("\t%s time.Time     `json:\"created\"`", util.StringPad("Created", max))
	ret.W("}")
	return ret, nil
}

func modelHistoryToData(m *model.Model) *golang.Block {
	ret := golang.NewBlock(m.Proper()+"HistoryToData", "func")
	ret.W("func (h *History) ToData() []any {")
	ret.W("\treturn []any{")
	ret.W("\t\th.ID,")
	for _, pk := range m.PKs() {
		ret.W("\t\th.%s%s,", m.Proper(), pk.Proper())
	}
	ret.W("\t\tutil.ToJSONBytes(h.Old, true),")
	ret.W("\t\tutil.ToJSONBytes(h.New, true),")
	ret.W("\t\tutil.ToJSONBytes(h.Changes, true),")
	ret.W("\t\th.Created,")
	ret.W("\t}")
	ret.W("}")
	return ret
}

func modelHistories(m *model.Model) *golang.Block {
	ret := golang.NewBlock(m.Proper()+"Histories", "struct")
	ret.W("type Histories []*History")
	return ret
}

func modelHistoryDTO(m *model.Model, enums enum.Enums) (*golang.Block, error) {
	ret := golang.NewBlock(m.Proper()+"HistoryDTO", "struct")
	ret.W("type historyDTO struct {")
	max := m.PKs().MaxCamelLength() + len(m.Camel())
	tmax := 15
	ret.W("\t%s %s `db:\"id\"`", util.StringPad("ID", max), util.StringPad("uuid.UUID", tmax))
	for _, pk := range m.PKs() {
		gt, err := pk.ToGoType(m.Package, enums)
		if err != nil {
			return nil, err
		}
		goType := util.StringPad(gt, tmax)
		ret.W("\t%s %s `db:\"%s_%s\"`", util.StringPad(m.Proper()+pk.Proper(), max), goType, m.Name, pk.Name)
	}
	ret.W("\t%s json.RawMessage `db:\"o\"`", util.StringPad("Old", max))
	ret.W("\t%s json.RawMessage `db:\"n\"`", util.StringPad("New", max))
	ret.W("\t%s json.RawMessage `db:\"c\"`", util.StringPad("Changes", max))
	ret.W("\t%s time.Time       `db:\"created\"`", util.StringPad("Created", max))
	ret.W("}")
	return ret, nil
}

func modelHistoryDTOToHistory(m *model.Model) *golang.Block {
	ret := golang.NewBlock(m.Proper()+"HistoryDTOToHistory", "func")
	ret.W("func (h *historyDTO) ToHistory() *History {")
	ret.W("\to := util.ValueMap{}")
	ret.W("\t_ = util.FromJSON(h.Old, &o)")
	ret.W("\tn := util.ValueMap{}")
	ret.W("\t_ = util.FromJSON(h.New, &n)")
	ret.W("\tc := util.Diffs{}")
	ret.W("\t_ = util.FromJSON(h.Changes, &c)")
	pkCalls := make([]string, 0, len(m.PKs()))
	for _, pk := range m.PKs() {
		pkCalls = append(pkCalls, fmt.Sprintf("%s%s: h.%s%s", m.Proper(), pk.Proper(), m.Proper(), pk.Proper()))
	}
	ret.W("\treturn &History{ID: h.ID, %s, Old: o, New: n, Changes: c, Created: h.Created}", strings.Join(pkCalls, ", "))
	ret.W("}")
	return ret
}

func modelHistoryDTOs(m *model.Model) *golang.Block {
	ret := golang.NewBlock(m.Proper()+"HistoryDTOs", "func")
	ret.W("type historyDTOs []*historyDTO")
	ret.W("")
	ret.W("func (h historyDTOs) ToHistories() Histories {")
	ret.W("\tret := make(Histories, 0, len(h))")
	ret.W("\tfor _, x := range h {")
	ret.W("\t\tret = append(ret, x.ToHistory())")
	ret.W("\t}")
	ret.W("\treturn ret")
	ret.W("}")
	return ret
}
