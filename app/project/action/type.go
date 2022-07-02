package action

import (
	"projectforge.dev/projectforge/app/util"
)

type Type struct {
	Key         string
	Title       string
	Icon        string
	Description string
	Hidden      bool
}

var (
	TypeAudit    = Type{Key: "audit", Title: "Audit", Icon: "scale", Description: "Audits the project files, detecting invalid files and empty folders"}
	TypeBuild    = Type{Key: "build", Title: "Build", Icon: "hammer", Description: "Builds the project, many options available"}
	TypeCreate   = Type{Key: "create", Title: "Create", Icon: "folder-plus", Description: "Creates a new project"}
	TypeDebug    = Type{Key: "debug", Title: "Debug", Icon: "bug", Description: "Dumps a ton of information about the project"}
	TypeDoctor   = Type{Key: "doctor", Title: "Doctor", Icon: "first-aid", Description: "Makes sure your machine has the required dependencies"}
	TypeGenerate = Type{Key: "generate", Title: "Generate", Icon: "forward", Description: "Applies pending changes to files as required"}
	TypePreview  = Type{Key: "preview", Title: "Preview", Icon: "play", Description: "Show what would happen if you generate"}
	TypeSVG      = Type{Key: "svg", Title: "SVG", Icon: "icons", Description: "Builds the project's SVG files"}
	TypeValidate = Type{Key: "validate", Title: "Validate", Icon: "glasses", Description: "Validates the project config"}
	TypeTest     = Type{Key: "test", Title: "Test", Icon: "wrench", Description: "Runs internal tests, you probably don't want this", Hidden: true}
)

var (
	AllTypes     = []Type{TypeAudit, TypeBuild, TypeCreate, TypeDebug, TypeDoctor, TypeGenerate, TypePreview, TypeSVG, TypeTest, TypeValidate}
	ProjectTypes = []Type{TypePreview, TypeGenerate, TypeAudit, TypeValidate, TypeBuild}
)

func TypeFromString(s string) Type {
	for _, t := range AllTypes {
		if t.Key == s {
			return t
		}
	}
	return Type{Key: s, Title: "Error", Icon: "star", Description: "No action type available with key [" + s + "]"}
}

func (t *Type) String() string {
	return t.Key
}

func (t *Type) MarshalJSON() ([]byte, error) {
	return util.ToJSONBytes(t.Key, false), nil
}

func (t *Type) UnmarshalJSON(data []byte) error {
	var s string
	if err := util.FromJSON(data, &s); err != nil {
		return err
	}
	x := TypeFromString(s)
	*t = x
	return nil
}
