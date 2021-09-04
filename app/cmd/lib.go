// Code generated by Project Forge, see https://projectforge.dev for details.
package cmd

import (
	"github.com/pkg/errors"

	"github.com/kyleu/projectforge/app"
	"github.com/kyleu/projectforge/app/util"
)

// Lib starts the application as a library, returning the actual TCP port the server is listening on (as an int32 to make interop easier).
func Lib() (int32, error) {
	if _buildInfo == nil {
		_buildInfo = &app.BuildInfo{Version: "TODO", Commit: "TODO", Date: "TODO"}
	}
	f := &Flags{Address: "0.0.0.0", Port: 0}

	if err := initIfNeeded(); err != nil {
		return 0, errors.Wrap(err, "error initializing application")
	}

	r, logger, err := loadServer(f, _logger)
	if err != nil {
		return 0, err
	}

	port, listener, err := listen(f.Address, f.Port)
	if err != nil {
		return 0, err
	}

	logger.Infof("%s library started on port [%d]", util.AppName, port)

	go func() {
		e := serve(util.AppKey, listener, r)
		if e != nil {
			panic(errors.WithStack(e))
		}
	}()

	return int32(port), nil
}
