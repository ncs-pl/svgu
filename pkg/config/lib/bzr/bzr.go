package bzr

import (
	_ "embed"
	"go.nc0.fr/svgu/pkg/config/lib/prelude"
	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
	"sync"
)

var (
	once = sync.Once{}
	bzr  = starlark.StringDict{}
	//go:embed bzr.star
	bzrFile string
	bzrErr  error
)

// LoadBzrModule loads the Bazaar module.
func LoadBzrModule(t *starlark.Thread) (starlark.StringDict, error) {
	once.Do(func() {
		env := starlark.StringDict{
			"module": starlark.NewBuiltin("module",
				prelude.InternModule),
			"make_module": starlark.NewBuiltin("mod",
				starlarkstruct.MakeModule),
		}
		bzr, bzrErr = starlark.ExecFile(t, "bzr.star", bzrFile, env)
	})

	return bzr, bzrErr
}
