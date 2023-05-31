package config

import (
	"fmt"
	"go.nc0.fr/svgu/pkg/config/lib/bzr"
	"go.nc0.fr/svgu/pkg/config/lib/fossil"
	"go.nc0.fr/svgu/pkg/config/lib/git"
	"go.nc0.fr/svgu/pkg/config/lib/hg"
	"go.nc0.fr/svgu/pkg/config/lib/prelude"
	"go.nc0.fr/svgu/pkg/config/lib/svn"
	"go.nc0.fr/svgu/pkg/types"
	"go.starlark.net/starlark"
)

// ExecConfig configures the Starlark environment and executes the given
// configuration file "fl".
// The function returns a list of registered modules, or an error if something
// went wrong.
func ExecConfig(fl string) (*types.Index, error) {
	th := &starlark.Thread{
		Name: "exec " + fl,
		Load: load,
	}

	// TODO(nc0): add built-ins
	env := starlark.StringDict{
		"index":  starlark.NewBuiltin("index", prelude.InternIndex),
		"module": starlark.NewBuiltin("module", prelude.InternModule),
	}

	prelude.Registered = types.Index{
		Domain:  "",
		Modules: make(map[string]types.Module),
	}
	if _, err := starlark.ExecFile(th, fl, nil, env); err != nil {
		return &types.Index{}, err
	}

	return &prelude.Registered, nil
}

// load loads a module from the given path.
func load(t *starlark.Thread, module string) (starlark.StringDict, error) {
	switch module {
	case "@svgu/git.star": // git
		return git.LoadGitModule(t)
	case "@svgu/hg.star": // mercurial
		return hg.LoadHgModule(t)
	case "@svgu/svn.star": // subversion
		return svn.LoadSvnModule(t)
	case "@svgu/fossil.star": // fossil
		return fossil.LoadFossilModule(t)
	case "@svgu/bzr.star": // bazaar
		return bzr.LoadBzrModule(t)
	default:
		return nil, fmt.Errorf("unknown module %q", module)
	}
}
