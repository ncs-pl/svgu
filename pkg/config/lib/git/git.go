// Package git provides Starlark macros to declare modules hosted on Git
// repositories.
package git

import (
	_ "embed"
	"go.nc0.fr/svgu/pkg/config/lib/prelude"
	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
	"sync"
)

var (
	once = sync.Once{}
	git  = starlark.StringDict{}
	//go:embed git.star
	gitFile string
	gitErr  error
)

// LoadGitModule loads the git module.
func LoadGitModule(t *starlark.Thread) (starlark.StringDict, error) {
	once.Do(func() {
		env := starlark.StringDict{
			"module": starlark.NewBuiltin("module",
				prelude.InternModule),
			"make_module": starlark.NewBuiltin("mod",
				starlarkstruct.MakeModule),
		}
		git, gitErr = starlark.ExecFile(t, "git.star", gitFile, env)
	})

	return git, gitErr
}
