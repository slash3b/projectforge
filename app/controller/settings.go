// Code generated by Project Forge, see https://projectforge.dev for details.
package controller

import (
	"github.com/valyala/fasthttp"

	"github.com/kyleu/projectforge/app"
	"github.com/kyleu/projectforge/app/controller/cutil"
	"github.com/kyleu/projectforge/app/settings"
	"github.com/kyleu/projectforge/app/user"
	"github.com/kyleu/projectforge/views/vsettings"
)

func Settings(rc *fasthttp.RequestCtx) {
	act("settings", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		current := &settings.Settings{}
		ps.Title = "Settings"
		ps.Data = current
		return render(rc, as, &vsettings.Settings{Settings: current, Perms: user.GetPermissions()}, ps, "settings")
	})
}
