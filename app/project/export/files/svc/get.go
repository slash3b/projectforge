package svc

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"

	"projectforge.dev/projectforge/app/file"
	"projectforge.dev/projectforge/app/project/export/enum"
	"projectforge.dev/projectforge/app/project/export/files/helper"
	"projectforge.dev/projectforge/app/project/export/golang"
	"projectforge.dev/projectforge/app/project/export/model"
)

//nolint:gocognit
func ServiceGet(m *model.Model, args *model.Args, addHeader bool) (*file.File, error) {
	dbRef := args.DBRef()
	g := golang.NewFile(m.Package, []string{"app", m.PackageWithGroup("")}, "serviceget")
	for _, imp := range helper.ImportsForTypes("go", m.PKs().Types()...) {
		g.AddImport(imp)
	}
	if len(m.PKs()) > 1 {
		g.AddImport(helper.ImpFmt)
	}
	g.AddImport(helper.ImpAppUtil, helper.ImpContext, helper.ImpErrors, helper.ImpSQLx, helper.ImpFilter, helper.ImpDatabase)
	if err := helper.SpecialImports(g, m.IndexedColumns(), m.PackageWithGroup(""), args.Enums); err != nil {
		return nil, err
	}
	g.AddBlocks(serviceList(m, args.DBRef()))
	g.AddBlocks(serviceCount(g, m, args.DBRef()))
	pkLen := len(m.PKs())
	if pkLen > 0 {
		gbpk, err := serviceGetByPK(m, dbRef, args.Enums)
		if err != nil {
			return nil, err
		}
		g.AddBlocks(gbpk)
		if pkLen == 1 {
			x, err := serviceGetMultipleSinglePK(m, dbRef, args.Enums)
			if err != nil {
				return nil, err
			}
			g.AddBlocks(x)
		} else {
			g.AddBlocks(serviceGetMultipleManyPKs(m, dbRef))
		}
	}
	getBys := map[string]model.Columns{}
	if pkLen > 1 {
		for _, pkCol := range m.PKs() {
			getBys[pkCol.Name] = model.Columns{pkCol}
		}
	}
	for _, grp := range m.GroupedColumns() {
		g.AddImport(helper.ImpAppUtil)
		g.AddBlocks(serviceGrouped(m, grp, args.DBRef()))
		getBys[grp.Name] = model.Columns{grp}
	}
	for _, rel := range m.Relations {
		cols := rel.SrcColumns(m)
		colStr := strings.Join(cols.Names(), ",")
		getBys[colStr] = cols
	}
	for _, col := range m.Columns {
		if col.Indexed {
			getBys[col.Name] = model.Columns{col}
		}
	}
	keys := make([]string, 0, len(getBys))
	for k := range getBys {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, key := range keys {
		cols := getBys[key]
		name := "GetBy" + strings.Join(cols.ProperNames(), "")
		for _, imp := range helper.ImportsForTypes("go", cols.Types()...) {
			g.AddImport(imp)
		}
		returnMultiple := false
		for _, x := range cols {
			if !x.HasTag("unique") {
				returnMultiple = true
			}
		}
		gb, err := serviceGetBy(name, m, cols, returnMultiple, dbRef, args.Enums)
		if err != nil {
			return nil, err
		}
		g.AddBlocks(gb)
	}

	if len(m.AllSearches()) > 0 {
		g.AddImport(helper.ImpStrings)
		ss, err := serviceSearch(m, nil, dbRef, args.Enums)
		if err != nil {
			return nil, err
		}
		g.AddBlocks(ss)
	}
	for _, grp := range m.GroupedColumns() {
		if !grp.PK {
			if len(m.AllSearches()) > 0 {
				ss, err := serviceSearch(m, grp, dbRef, args.Enums)
				if err != nil {
					return nil, err
				}
				g.AddBlocks(ss)
			}
		}
	}
	g.AddBlocks(serviceListSQL(m, args.DBRef()))
	return g.Render(addHeader)
}

func serviceGrouped(m *model.Model, grp *model.Column, dbRef string) *golang.Block {
	name := "Get" + grp.ProperPlural()
	ret := golang.NewBlock(name, "func")
	ret.W("func (s *Service) %s(ctx context.Context, tx *sqlx.Tx%s, logger util.Logger) ([]*util.KeyValInt, error) {", name, getSuffix(m))
	ret.W("\twc := \"\"")
	if m.IsSoftDelete() {
		delCols := m.Columns.WithTag("deleted")
		ret.W("\tif !includeDeleted {")
		ret.W("\t\twc = %q", delCols[0].NameQuoted()+" is null")
		ret.W("\t}")
	}
	cols := fmt.Sprintf("%q as key, count(*) as val", grp.Name)
	ret.W("\tq := database.SQLSelectGrouped(%q, %s, wc, %q, %q, 0, 0)", cols, tableClauseFor(m), `"`+grp.Name+`"`, `"`+grp.Name+`"`)
	ret.W("\tvar ret []*util.KeyValInt")
	ret.W("\terr := s.%s.Select(ctx, &ret, q, tx, logger)", dbRef)
	ret.W("\tif err != nil {")
	ret.W("\t\treturn nil, errors.Wrap(err, \"unable to get %s by %s\")", m.TitlePluralLower(), grp.TitleLower())
	ret.W("\t}")
	ret.W("\treturn ret, nil")
	ret.W("}")
	return ret
}

func serviceList(m *model.Model, dbRef string) *golang.Block {
	ret := golang.NewBlock("List", "func")
	ret.W("func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params%s, logger util.Logger) (%s, error) {", getSuffix(m), m.ProperPlural())
	ret.W("\tparams = filters(params)")
	ret.W("\twc := \"\"")
	if m.IsSoftDelete() {
		delCols := m.Columns.WithTag("deleted")
		ret.W("\tif !includeDeleted {")
		ret.W("\t\twc = %q", delCols[0].NameQuoted()+" is null")
		ret.W("\t}")
	}
	ret.W("\tq := database.SQLSelect(columnsString, %s, wc, params.OrderByString(), params.Limit, params.Offset)", tableClauseFor(m))
	ret.W("\tret := dtos{}")
	ret.W("\terr := s.%s.Select(ctx, &ret, q, tx, logger)", dbRef)
	ret.W("\tif err != nil {")
	ret.W("\t\treturn nil, errors.Wrap(err, \"unable to get %s\")", m.TitlePluralLower())
	ret.W("\t}")
	ret.W("\treturn ret.To%s(), nil", m.ProperPlural())
	ret.W("}")
	return ret
}

func serviceCount(g *golang.File, m *model.Model, dbRef string) *golang.Block {
	g.AddImport(helper.ImpStrings)
	ret := golang.NewBlock("Count", "func")
	ret.W("func (s *Service) Count(ctx context.Context, tx *sqlx.Tx, whereClause string%s, logger util.Logger, args ...any) (int, error) {", getSuffix(m))
	ret.W("\tif strings.Contains(whereClause, \"'\") || strings.Contains(whereClause, \";\") {")
	ret.W("\t\treturn 0, errors.Errorf(\"invalid where clause [%%s]\", whereClause)")
	ret.W("\t}")
	if m.IsSoftDelete() {
		delCols := m.Columns.WithTag("deleted")
		ret.W("\tif !includeDeleted {")
		ret.W("\t\tif whereClause == \"\" {")
		ret.W("\t\t\twhereClause = %q", delCols[0].NameQuoted()+" is null")
		ret.W("\t\t} else {")
		ret.W("\t\t\twhereClause += \"and \" + %q", delCols[0].NameQuoted()+" is null")
		ret.W("\t\t}")
		ret.W("\t}")
	}
	ret.W("\tq := database.SQLSelectSimple(columnsString, %s, whereClause)", tableClauseFor(m))
	ret.W("\tret, err := s.%s.SingleInt(ctx, q, tx, logger, args...)", dbRef)
	ret.W("\tif err != nil {")
	ret.W("\t\treturn 0, errors.Wrap(err, \"unable to get count of %s\")", m.TitlePluralLower())
	ret.W("\t}")
	ret.W("\treturn int(ret), nil")
	ret.W("}")
	return ret
}

func serviceGetByPK(m *model.Model, dbRef string, enums enum.Enums) (*golang.Block, error) {
	return serviceGetBy("Get", m, m.PKs(), false, dbRef, enums)
}

func serviceGetBy(key string, m *model.Model, cols model.Columns, returnMultiple bool, dbRef string, enums enum.Enums) (*golang.Block, error) {
	if returnMultiple {
		return serviceGetByCols(key, m, cols, dbRef, enums)
	}
	return serviceGet(key, m, cols, dbRef, enums)
}

func serviceGet(key string, m *model.Model, cols model.Columns, dbRef string, enums enum.Enums) (*golang.Block, error) {
	if key == "" {
		key = "GetBy" + cols.Smushed()
	}
	ret := golang.NewBlock(key, "func")
	msg := "func (s *Service) %s(ctx context.Context, tx *sqlx.Tx, %s%s, logger util.Logger) (*%s, error) {"
	args, err := cols.Args(m.Package, enums)
	if err != nil {
		return nil, err
	}
	ret.W(msg, key, args, getSuffix(m), m.Proper())
	if slices.Equal(m.PKs().Names(), cols.Names()) {
		ret.W("\twc := defaultWC(0)")
	} else {
		wc := make([]string, 0, len(cols))
		for idx, col := range cols {
			wc = append(wc, fmt.Sprintf("%q = $%d", col.Name, idx+1))
		}
		ret.W("\twc := %q", strings.Join(wc, " and "))
	}
	if m.IsSoftDelete() {
		ret.W("\twc = addDeletedClause(wc, includeDeleted)")
	}
	ret.W("\tret := &dto{}")
	ret.W("\tq := database.SQLSelectSimple(columnsString, %s, wc)", tableClauseFor(m))
	ret.W("\terr := s.%s.Get(ctx, ret, q, tx, logger, %s)", dbRef, strings.Join(cols.CamelNames(), ", "))
	ret.W("\tif err != nil {")
	sj := strings.Join(cols.CamelNames(), ", ")
	decls := make([]string, 0, len(cols))
	for _, c := range cols {
		decls = append(decls, c.Camel()+" [%%v]")
	}
	ret.W("\t\treturn nil, errors.Wrapf(err, \"unable to get %s by %s\", %s)", m.Camel(), strings.Join(decls, ", "), sj)
	ret.W("\t}")
	ret.W("\treturn ret.To%s(), nil", m.Proper())
	ret.W("}")
	return ret, nil
}

func serviceGetByCols(key string, m *model.Model, cols model.Columns, dbRef string, enums enum.Enums) (*golang.Block, error) {
	if key == "" {
		key = "GetBy" + cols.Smushed()
	}
	ret := golang.NewBlock(key, "func")
	args, err := cols.Args(m.Package, enums)
	if err != nil {
		return nil, err
	}
	msg := "func (s *Service) %s(ctx context.Context, tx *sqlx.Tx, %s, params *filter.Params%s, logger util.Logger) (%s, error) {"
	msg = fmt.Sprintf(msg, key, args, getSuffix(m), m.ProperPlural())
	if len(msg) > 160 {
		msg += " //nolint:lll"
	}
	ret.W(msg)
	ret.W("\tparams = filters(params)")
	ret.W("\twc := %q", cols.WhereClause(0))
	if m.IsSoftDelete() {
		ret.W("\twc = addDeletedClause(wc, includeDeleted)")
	}
	ret.W("\tq := database.SQLSelect(columnsString, %s, wc, params.OrderByString(), params.Limit, params.Offset)", tableClauseFor(m))
	ret.W("\tret := dtos{}")
	ret.W("\terr := s.%s.Select(ctx, &ret, q, tx, logger, %s)", dbRef, strings.Join(cols.CamelNames(), ", "))
	ret.W("\tif err != nil {")
	sj := strings.Join(cols.CamelNames(), ", ")
	decls := make([]string, 0, len(cols))
	for _, c := range cols {
		decls = append(decls, c.Camel()+" [%%v]")
	}
	ret.W("\t\treturn nil, errors.Wrapf(err, \"unable to get %s by %s\", %s)", m.Plural(), strings.Join(decls, ", "), sj)
	ret.W("\t}")
	ret.W("\treturn ret.To%s(), nil", m.ProperPlural())
	ret.W("}")
	return ret, nil
}

func serviceListSQL(m *model.Model, dbRef string) *golang.Block {
	ret := golang.NewBlock("ListSQL", "func")
	ret.W("func (s *Service) ListSQL(ctx context.Context, tx *sqlx.Tx, sql string, logger util.Logger, values ...any) (%s, error) {", m.ProperPlural())
	ret.W("\tret := dtos{}")
	ret.W("\terr := s.%s.Select(ctx, &ret, sql, tx, logger, values...)", dbRef)
	ret.W("\tif err != nil {")
	ret.W("\t\treturn nil, errors.Wrap(err, \"unable to get %s using custom SQL\")", m.TitlePluralLower())
	ret.W("\t}")
	ret.W("\treturn ret.To%s(), nil", m.ProperPlural())
	ret.W("}")
	return ret
}

func tableClauseFor(m *model.Model) string {
	if m.IsRevision() {
		return "tablesJoined"
	}
	return "tableQuoted"
}
