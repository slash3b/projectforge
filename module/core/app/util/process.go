package util

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func RunProcess(cmd string, path string, logger *zap.SugaredLogger, in io.Reader, out io.Writer, er io.Writer) (int, error) {
	logger.Info(fmt.Sprintf("Running [" + cmd + "] in [" + path + "]"))

	args := strings.Split(cmd, " ")
	firstArg := args[0]

	var err error
	if !strings.Contains(firstArg, "/") {
		firstArg, err = exec.LookPath(firstArg)
		if err != nil {
			return -1, errors.Wrap(err, "unable to look up cmd ["+firstArg+"]")
		}
	}

	if in == nil {
		in = os.Stdin
	}
	if out == nil {
		out = os.Stdout
	}
	if er == nil {
		er = os.Stderr
	}

	c := exec.Cmd{Path: firstArg, Args: args, Stdin: in, Stdout: out, Stderr: er, Dir: path}
	err = c.Start()
	if err != nil {
		return -1, errors.Wrap(err, fmt.Sprintf("unable to start [%v] (%T)", cmd, err))
	}
	err = c.Wait()
	if err != nil {
		ec, ok := err.(*exec.ExitError) // nolint
		if ok {
			return ec.ExitCode(), nil
		}
		return -1, errors.Wrap(err, fmt.Sprintf("unable to run [%v] (%T)", cmd, err))
	}
	return 0, nil
}

func RunProcessSimple(cmd string, path string, logger *zap.SugaredLogger) (int, string, error) {
	var buf bytes.Buffer
	ec, err := RunProcess(cmd, path, logger, nil, &buf, &buf)
	if err != nil {
		return -1, "", err
	}
	return ec, buf.String(), nil
}
