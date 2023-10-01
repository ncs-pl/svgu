// Copyright Nicolas Paul (2023)
//
// * Nicolas Paul
//
// This software is a computer program whose purpose is to allow the hosting
// and sharing of Go modules using a personal domain.
//
// This software is governed by the CeCILL license under French law and
// abiding by the rules of distribution of free software.  You can  use,
// modify and/ or redistribute the software under the terms of the CeCILL
// license as circulated by CEA, CNRS and INRIA at the following URL
// "http://www.cecill.info".
//
// As a counterpart to the access to the source code and  rights to copy,
// modify and redistribute granted by the license, users are provided only
// with a limited warranty  and the software's author,  the holder of the
// economic rights,  and the successive licensors  have only  limited
// liability.
//
// In this respect, the user's attention is drawn to the risks associated
// with loading,  using,  modifying and/or developing or reproducing the
// software by the user in light of its specific status of free software,
// that may mean  that it is complicated to manipulate,  and  that  also
// therefore means  that it is reserved for developers  and  experienced
// professionals having in-depth computer knowledge. Users are therefore
// encouraged to load and test the software's suitability as regards their
// requirements in conditions enabling the security of their systems and/or
// data to be ensured and,  more generally, to use and operate it in the
// same conditions as regards security.
//
// The fact that you are presently reading this means that you have had
// knowledge of the CeCILL license and that you accept its terms.

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

	Registered.AddModule(name, &types.Module{
		Path: name,
		Vcs:  v,
		Repo: repo,
		Dir:  dir,
		File: file,
	})

	return starlark.None, nil
}
