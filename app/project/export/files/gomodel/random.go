package gomodel

import (
	"fmt"

	"projectforge.dev/projectforge/app/lib/types"
	"projectforge.dev/projectforge/app/project/export/enum"
	"projectforge.dev/projectforge/app/project/export/golang"
	"projectforge.dev/projectforge/app/project/export/model"
	"projectforge.dev/projectforge/app/util"
)

const nilKey = "nil"

func modelRandom(m *model.Model, enums enum.Enums) *golang.Block {
	ret := golang.NewBlock(m.Proper()+"Random", "struct")
	ret.W("func Random() *%s {", m.Proper())
	ret.W("\treturn &%s{", m.Proper())
	maxColLength := m.Columns.MaxCamelLength() + 1
	for _, col := range m.Columns {
		ret.W("\t\t%s %s,", util.StringPad(col.Proper()+":", maxColLength), randFor(col, m.PackageWithGroup(""), enums))
	}
	ret.W("\t}")
	ret.W("}")
	return ret
}

func randFor(col *model.Column, pkg string, enums enum.Enums) string {
	switch col.Type.Key() {
	case types.KeyAny:
		return types.KeyNil
	case types.KeyBool:
		return "util.RandomBool()"
	case types.KeyEnum:
		et, err := model.AsEnumInstance(col.Type, enums)
		if err != nil {
			return "ERROR:" + err.Error()
		}
		if pkg == et.PackageWithGroup("") {
			return fmt.Sprintf("%s(util.RandomString(12))", et.Proper())
		}
		return fmt.Sprintf("%s.%s(util.RandomString(12))", et.Package, et.Proper())
	case types.KeyInt:
		return "util.RandomInt(10000)"
	case types.KeyFloat:
		return "util.RandomFloat(1000)"
	case types.KeyList:
		return nilKey
	case types.KeyMap, types.KeyValueMap:
		return "util.RandomValueMap(4)"
	case types.KeyReference:
		return nilKey
	case types.KeyString:
		switch col.Format {
		case "url":
			return "\"https://\" + util.RandomString(6) + \".com/\" + util.RandomString(6)"
		default:
			return "util.RandomString(12)"
		}
	case types.KeyDate, types.KeyTimestamp:
		if col.HasTag("deleted") {
			return types.KeyNil
		}
		if col.Nullable {
			return "util.NowPointer()"
		}
		return "time.Now()"
	case types.KeyUUID:
		if col.Nullable {
			return "util.UUIDP()"
		}
		return "util.UUID()"
	default:
		return "TODO"
	}
}
