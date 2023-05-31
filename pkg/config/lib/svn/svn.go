package svn

import (
	_ "embed"
	"go.nc0.fr/svgu/pkg/config/lib/prelude"
	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
	"sync"
)

var (
	once = sync.Once{}
	svn  = starlark.StringDict{}
	//go:embed svn.star
	svnFile string
	svnErr  error
)

// LoadSvnModule loads the Subversion module.
func LoadSvnModule(t *starlark.Thread) (starlark.StringDict, error) {
	once.Do(func() {
		env := starlark.StringDict{
			"module": starlark.NewBuiltin("module",
				prelude.InternModule),
			"make_module": starlark.NewBuiltin("mod",
				starlarkstruct.MakeModule),
		}
		svn, svnErr = starlark.ExecFile(t, "svn.star", svnFile, env)
	})

	return svn, svnErr
}
