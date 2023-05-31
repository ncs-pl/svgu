package config

import (
	"fmt"
	"go.nc0.fr/svgu/pkg/config/lib/bzr"
	"go.nc0.fr/svgu/pkg/config/lib/fossil"
	"go.nc0.fr/svgu/pkg/config/lib/git"
	"go.nc0.fr/svgu/pkg/config/lib/hg"
	"go.nc0.fr/svgu/pkg/config/lib/svn"
	"go.nc0.fr/svgu/pkg/types"
	"go.starlark.net/starlark"
)

var registered types.Index

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
		"index":  starlark.NewBuiltin("index", InternIndex),
		"module": starlark.NewBuiltin("module", InternModule),
	}

	registered = types.Index{}
	if _, err := starlark.ExecFile(th, fl, nil, env); err != nil {
		return &types.Index{}, err
	}

	return &registered, nil
}

// load loads a module from the given path.
func load(t *starlark.Thread, module string) (starlark.StringDict, error) {
	switch module {
	case "git.star": // git
		return git.LoadGitModule(t)
	case "hg.star": // mercurial
		return hg.LoadHgModule(t)
	case "svn.star": // subversion
		return svn.LoadSvnModule(t)
	case "fossil.star": // fossil
		return fossil.LoadFossilModule(t)
	case "bzr.star": // bazaar
		return bzr.LoadBzrModule(t)
	default:
		return nil, fmt.Errorf("unknown module %q", module)
	}
}

// Injected built-ins.

// InternIndex represents the built-in function "index".
// index(domain) initializes a new index with the given domain.
func InternIndex(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	var domain string
	if err := starlark.UnpackArgs("index", args, kwargs,
		"domain", &domain); err != nil {
		return nil, err
	}

	registered.SetDomain(domain)

	return starlark.None, nil
}

// InternModule represents the built-in function "module".
// module(name, vcs, repo, dir=None, file=None) registers a new module into the
// index.
func InternModule(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {

	var name, vcs, repo, dir, file string
	if err := starlark.UnpackArgs("module", args, kwargs, "name",
		&name, "vcs", &vcs, "repo", &repo, "dir?", &dir, "file?", &file); err != nil {
		return nil, err
	}

	var v types.Vcs
	switch vcs {
	case "git":
		v = types.VcsGit
	case "hg":
		v = types.VcsMercurial
	case "svn":
		v = types.VcsSubversion
	case "fossil":
		v = types.VcsFossil
	case "bzr":
		v = types.VcsBazaar
	default:
		return nil, fmt.Errorf("unknown vcs %q", vcs)
	}

	registered.AddModule(name, types.Module{
		Path: name,
		Vcs:  v,
		Repo: repo,
		Dir:  dir,
		File: file,
	})

	return starlark.None, nil
}
