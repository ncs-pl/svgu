package hg

import (
	_ "embed"
	"go.nc0.fr/svgu/pkg/config/lib/prelude"
	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
	"sync"
)

var (
	once = sync.Once{}
	hg   = starlark.StringDict{}
	//go:embed hg.star
	hgFile string
	hgErr  error
)

// LoadHgModule loads the Mercurial module.
func LoadHgModule(t *starlark.Thread) (starlark.StringDict, error) {
	once.Do(func() {
		env := starlark.StringDict{
			"module": starlark.NewBuiltin("module",
				prelude.InternModule),
			"make_module": starlark.NewBuiltin("mod",
				starlarkstruct.MakeModule),
		}
		hg, hgErr = starlark.ExecFile(t, "hg.star", hgFile, env)
	})

	return hg, hgErr
}
