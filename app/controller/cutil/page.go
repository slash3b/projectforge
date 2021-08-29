// Code generated by Project Forge, see https://projectforge.dev for details.
package cutil

import (
	"context"
	"fmt"

	"github.com/go-gem/sessions"
	"github.com/valyala/fasthttp"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"github.com/kyleu/projectforge/app/menu"
	"github.com/kyleu/projectforge/app/user"
	"github.com/kyleu/projectforge/app/util"
)

type PageState struct {
	Title         string             `json:"title"`
	Description   string             `json:"description"`
	Method        string             `json:"method"`
	URI           *fasthttp.URI      `json:"-"`
	Menu          menu.Items         `json:"menu"`
	Breadcrumbs   Breadcrumbs        `json:"breadcrumbs"`
	Flashes       []string           `json:"flashes"`
	Session       *sessions.Session  `json:"-"`
	Profile       *user.Profile      `json:"profile"`
	Accounts      user.Accounts      `json:"accounts"`
	Icons         []string           `json:"icons"`
	RootIcon      string             `json:"rootIcon"`
	RootPath      string             `json:"rootPath"`
	RootTitle     string             `json:"rootTitle"`
	SearchPath    string             `json:"searchPath"`
	ProfilePath   string             `json:"profilePath"`
	Data          interface{}        `json:"data"`
	Logger        *zap.SugaredLogger `json:"-"`
	Context       context.Context    `json:"-"`
	Span          *trace.Span        `json:"-"`
	RenderElapsed float64            `json:"renderElapsed"`
}

func (p *PageState) AddIcon(n string) {
	for _, icon := range p.Icons {
		if icon == n {
			return
		}
	}
	p.Icons = append(p.Icons, n)
}

func (p *PageState) TitleString() string {
	if p.Title == "" {
		return util.AppName
	}
	return fmt.Sprintf("%s - %s", p.Title, util.AppName)
}

func (p *PageState) Close() {
	if p.Span != nil {
		(*p.Span).End()
	}
}
