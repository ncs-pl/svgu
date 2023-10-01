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

// Loader for the Starlark execution system with a pre-defined environment.

package main

import (
	"go.starlark.net/starlark"
)

// ExecConfig configures the Starlark environment and executes the given
// configuration file "fl".
// The function returns a list of registered modules, or an error if something
// went wrong.
func ExecConfig(fl string) (*Index, error) {
	th := &starlark.Thread{
		Name: "exec " + fl,
	}

	env := starlark.StringDict{
		"index":  starlark.NewBuiltin("index", InternIndex),
		"module": starlark.NewBuiltin("module", InternModule),
	}

	Registered = Index{
		Domain:  "",
		Modules: make(map[string]*Module),
	}
	if _, err := starlark.ExecFile(th, fl, nil, env); err != nil {
		return &Index{}, err
	}

	return &Registered, nil
}
