package svc

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

func ServiceHistory(m *model.Model, args *model.Args, addHeader bool) (*file.File, error) {
	dbRef := args.DBRef()
	g := golang.NewFile(m.Package, []string{"app", m.PackageWithGroup("")}, "servicehistory")
	g.AddImport(helper.ImpContext, helper.ImpUUID, helper.ImpErrors, helper.ImpFmt, helper.ImpTime, helper.ImpStrings)
	g.AddImport(helper.ImpSQLx, helper.ImpAppUtil, helper.ImpDatabase)
	gh, err := serviceHistoryGetHistories(m, dbRef, args.Enums)
	if err != nil {
		return nil, err
	}
	g.AddBlocks(serviceHistoryVars(m), serviceHistoryGetHistory(m, dbRef), gh, serviceHistorySaveHistory(m))
	return g.Render(addHeader)
}

func serviceHistoryVars(m *model.Model) *golang.Block {
	ret := golang.NewBlock("HistoryVars", "func")
	ret.W("var (")
	xx := make(model.Columns, 0, len(m.PKs()))
	for _, pk := range m.PKs() {
		x := pk.Clone()
		x.Name = m.Name + "_" + x.Name
		xx = append(xx, x)
	}
	ret.W("\thistoryColumns       = "+`[]string{"id", %s, "o", "n", "c", "created"}`, strings.Join(xx.NamesQuoted(), ", "))
	ret.W("\thistoryColumnsQuoted = util.StringArrayQuoted(historyColumns)")
	ret.W("\thistoryColumnsString = strings.Join(historyColumnsQuoted, \", \")")
	ret.W("")
	ret.W("\thistoryTable       = table + \"_history\"")
	ret.W("\thistoryTableQuoted = fmt.Sprintf(\"%%q\", historyTable)")
	ret.W(")")
	return ret
}

func serviceHistoryGetHistory(m *model.Model, dbRef string) *golang.Block {
	ret := golang.NewBlock("GetHistory", "func")
	ret.W("func (s *Service) GetHistory(ctx context.Context, tx *sqlx.Tx, id uuid.UUID, logger util.Logger) (*History, error) {")
	ret.W("\tq := database.SQLSelectSimple(historyColumnsString, historyTableQuoted, \"id = $1\")")
	ret.W("\tret := historyDTO{}")
	ret.W("\terr := s.%s.Get(ctx, &ret, q, tx, logger, id)", dbRef)
	ret.W("\tif err != nil {")
	ret.W("\t\treturn nil, errors.Wrapf(err, \"unable to get %s history [%%%%s]\", id.String())", m.TitleLower())
	ret.W("\t}")
	ret.W("\treturn ret.ToHistory(), nil")
	ret.W("}")
	return ret
}

func serviceHistoryGetHistories(m *model.Model, dbRef string, enums enum.Enums) (*golang.Block, error) {
	ret := golang.NewBlock("GetHistories", "func")
	msg := "func (s *Service) GetHistories(ctx context.Context, tx *sqlx.Tx, %s, logger util.Logger) (Histories, error) {"
	args, err := m.PKs().Args(m.Package, enums)
	if err != nil {
		return nil, err
	}
	ret.W(msg, args)
	pks := m.PKs()
	joins := make([]string, 0, len(pks))
	logs := make([]string, 0, len(pks))
	for idx, pk := range pks {
		joins = append(joins, fmt.Sprintf("%s_%s = $%d", m.Name, pk.Name, idx+1))
		logs = append(logs, pk.Camel()+" [%%v]")
	}
	ret.W("\tq := database.SQLSelectSimple(historyColumnsString, historyTableQuoted, %q)", strings.Join(joins, " and "))
	ret.W("\tret := historyDTOs{}")
	ret.W("\terr := s.%s.Select(ctx, &ret, q, tx, logger, %s)", dbRef, strings.Join(pks.CamelNames(), ", "))
	ret.W("\tif err != nil {")
	const msg2 = "\t\treturn nil, errors.Wrapf(err, \"unable to get %s by %s\", %s)"
	ret.W(msg2, m.TitlePluralLower(), strings.Join(logs, ", "), strings.Join(pks.CamelNames(), ", "))
	ret.W("\t}")
	ret.W("\treturn ret.ToHistories(), nil")
	ret.W("}")
	return ret, nil
}

func serviceHistorySaveHistory(m *model.Model) *golang.Block {
	ret := golang.NewBlock("SaveHistory", "func")
	const decl = "func (s *Service) SaveHistory(ctx context.Context, tx *sqlx.Tx, o *%s, n *%s, logger util.Logger) (*History, error) {"
	ret.W(decl, m.Proper(), m.Proper())
	ret.W("\tdiffs := o.Diff(n)")
	ret.W("\tif len(diffs) == 0 {")
	ret.W("\t\treturn nil, nil")
	ret.W("\t}")
	ret.W("\tq := database.SQLInsert(historyTableQuoted, historyColumns, 1, \"\")")
	ret.W("\th := &historyDTO{")
	max := m.PKs().MaxCamelLength() + len(m.Proper()) + 1
	ret.W("\t\t%s util.UUID(),", util.StringPad("ID:", max))
	for _, pk := range m.PKs() {
		ret.W("\t\t%s o.%s,", util.StringPad(m.Proper()+pk.Proper()+":", max), pk.Proper())
	}
	ret.W("\t\t%s util.ToJSONBytes(o, true),", util.StringPad("Old:", max))
	ret.W("\t\t%s util.ToJSONBytes(n, true),", util.StringPad("New:", max))
	ret.W("\t\t%s util.ToJSONBytes(diffs, true),", util.StringPad("Changes:", max))
	ret.W("\t\t%s time.Now(),", util.StringPad("Created:", max))
	ret.W("\t}")
	ret.W("\thist := h.ToHistory()")
	ret.W("\terr := s.db.Insert(ctx, q, tx, logger, hist.ToData()...)")
	ret.W("\tif err != nil {")
	ret.W("\t\treturn nil, errors.Wrap(err, \"unable to insert %s\")", m.TitleLower())
	ret.W("\t}")
	ret.W("\treturn hist, nil")
	ret.W("}")
	return ret
}
