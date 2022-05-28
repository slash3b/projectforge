package git

import (
	"context"
	"strings"

	"projectforge.dev/projectforge/app/util"
)

func gitBranch(ctx context.Context, path string, logger util.Logger) string {
	out, err := gitCmd(ctx, "branch --show-current", path, logger)
	if err != nil {
		if isNoRepo(err) {
			return "norepo"
		}
		return "error: " + err.Error()
	}
	return strings.TrimSpace(out)
}
