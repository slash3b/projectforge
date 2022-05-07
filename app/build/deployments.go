package build

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"projectforge.dev/projectforge/app/diff"
	"projectforge.dev/projectforge/app/lib/filesystem"
	"projectforge.dev/projectforge/app/util"
)

func Deployments(ctx context.Context, curr string, fs filesystem.FileLoader, fix bool, path string, deps []string) ([]string, diff.Diffs, error) {
	var logs []string
	var ret diff.Diffs

	if len(deps) == 0 {
		logs = append(logs, "no deployments defined for this project")
	}

	for _, dep := range deps {
		b, err := fs.ReadFile(dep)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "deployment file [%s] does not exist", dep)
		}

		hit := util.ValueMap{}
		lines := strings.Split(string(b), "\n")
		final := make([]string, 0, len(lines))

		df := &diff.Diff{Path: dep, Status: diff.StatusIdentical}

		for _, line := range lines {
			if idx := strings.Index(line, "tag: "); idx > -1 {
				v := strings.TrimSpace(line[idx+5:])
				if curr == v || curr == strings.TrimPrefix(v, "v") {
					final = append(final, line)
				} else {
					hit[v] = struct{}{}
					if fix {
						final = append(final, strings.ReplaceAll(line, v, curr))
					} else {
						final = append(final, line)
					}
				}
			} else {
				final = append(final, line)
			}
		}

		if len(hit) > 0 {
			ret = append(ret, df)
			hits := strings.Join(hit.Keys(), ", ")
			if fix && (path == "" || path == dep) {
				logs = append(logs, fmt.Sprintf("updated version [%s] to [%s]", hits, curr))
			} else {
				df.Status = diff.StatusDifferent
				df.Patch = hits + " -> " + curr
			}
		}
	}

	return logs, ret, nil
}
