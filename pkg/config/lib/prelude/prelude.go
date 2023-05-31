package prelude

import (
	"fmt"
	"go.nc0.fr/svgu/pkg/types"
	"go.starlark.net/starlark"
	"strings"
)

// https://stackoverflow.com/questions/1976007/what-characters-are-forbidden-in-windows-and-linux-directory-names
const invalidName string = "..\\/<>:\"|?* \t\n\r\b\findex"

// Registered represents the index of registered modules.
var Registered types.Index

// InternIndex represents the built-in function "index".
// index(domain) initializes a new index with the given domain.
func InternIndex(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	var domain string
	if err := starlark.UnpackArgs("index", args, kwargs,
		"domain", &domain); err != nil {
		return nil, err
	}

	Registered.SetDomain(domain)

	return starlark.None, nil
}

// InternModule represents the built-in function "module".
// module(name, vcs, repo, dir, file) registers a new module into the
// index.
func InternModule(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {

	var name, vcs, repo, dir, file string
	if err := starlark.UnpackArgs("module", args, kwargs, "name",
		&name, "vcs", &vcs, "repo", &repo, "dir", &dir, "file", &file); err != nil {
		return nil, err
	}

	if Registered.Domain == "" {
		return nil, fmt.Errorf("index not initialized")
	}

	if name == "" {
		return nil, fmt.Errorf("module name cannot be empty")
	}

	if vcs == "" {
		return nil, fmt.Errorf("module %q vcs cannot be empty", name)
	}

	if repo == "" {
		return nil, fmt.Errorf("module %q repo cannot be empty", name)
	}

	// Check for name conditions.
	if strings.Contains(invalidName, name) {
		return nil, fmt.Errorf("module %q name is invalid", name)
	}

	if Registered.CheckModule(name) {
		return nil, fmt.Errorf("module %q already exists", name)
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

	Registered.AddModule(name, types.Module{
		Path: name,
		Vcs:  v,
		Repo: repo,
		Dir:  dir,
		File: file,
	})

	return starlark.None, nil
}
