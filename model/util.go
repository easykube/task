package model

import (
	"os"
	"strings"

	"github.com/flosch/pongo2"
)

func AddArgsToContext(args Args, context pongo2.Context) {
	for _, arg := range args {
		context[arg.Arg] = arg.Val
	}
}

func AddEnvVarsToContext(context pongo2.Context) {
	lines := os.Environ()
	for _, line := range lines {
		list := strings.Split(line, "=")
		if len(list) > 1 && list[0] != "" && list[1] != "" {
			context[list[0]] = list[1]
		}
	}
}
