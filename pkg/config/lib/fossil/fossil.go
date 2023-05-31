package fossil

import (
	_ "embed"
	"go.nc0.fr/svgu/pkg/config/lib/prelude"
	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
	"sync"
)

var (
	once   = sync.Once{}
	fossil = starlark.StringDict{}
	//go:embed fossil.star
	fossilFile string
	fossilErr  error
)

// LoadFossilModule loads the Fossil module.
func LoadFossilModule(t *starlark.Thread) (starlark.StringDict, error) {
	once.Do(func() {
		env := starlark.StringDict{
			"module": starlark.NewBuiltin("module",
				prelude.InternModule),
			"make_module": starlark.NewBuiltin("mod",
				starlarkstruct.MakeModule),
		}
		fossil, fossilErr = starlark.ExecFile(t, "fossil.star", fossilFile, env)
	})

	return fossil, fossilErr
}
