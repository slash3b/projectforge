package action

import (
	"github.com/kyleu/projectforge/app/codegen"
	"github.com/kyleu/projectforge/app/filesystem"
	"github.com/kyleu/projectforge/app/module"
	"github.com/kyleu/projectforge/app/project"
	"github.com/kyleu/projectforge/app/util"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Params struct {
	Span       trace.Span
	ProjectKey string
	T          Type
	Cfg        util.ValueMap
	RootFiles  filesystem.FileLoader
	MSvc       *module.Service
	PSvc       *project.Service
	CSvc       *codegen.Service
	CLI        bool
	Logger     *zap.SugaredLogger
}
