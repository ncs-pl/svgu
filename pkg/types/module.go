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

package types

import (
	"fmt"
	"go.nc0.fr/svgu/pkg/templates"
	"os"
	"path"
	"strings"
	"sync"
)

// Vcs is an enum for version control systems supported by the standard Go
// toolchain.
//
// See https://pkg.go.dev/cmd/go#hdr-Module_configuration_for_non_public_modules
type Vcs string

// Vcs enum.
const (
	VcsBazaar     Vcs = "bzr"
	VcsFossil     Vcs = "fossil"
	VcsGit        Vcs = "git"
	VcsMercurial  Vcs = "hg"
	VcsSubversion Vcs = "svn"
)

// Module represents a Go module to index.
type Module struct {
	Path string // module path (without domain)
	Vcs  Vcs    // vcs system
	Repo string // repository's home
	Dir  string // url template
	File string // url template

	// internal
	mu sync.Mutex
}

// GenerateFile generates the index file.
func (m *Module) GenerateFile(out string, domain string) error {
	m.mu.Lock()
	p := m.Path
	v := m.Vcs
	r := m.Repo
	d := m.Dir
	f := m.File
	m.mu.Unlock()

	outf := path.Join(out, p+".html")

	// Create the file.
	if strings.Contains(p, "/") {
		if err := os.MkdirAll(path.Dir(outf), 0755); err != nil {
			return err
		}
	}

	fd, err := os.Create(outf)
	if err != nil {
		return err
	}
	defer func(fd *os.File) {
		err := fd.Close()
		if err != nil {
			panic(err)
		}
	}(fd)

	// Execute the template and write the output to the file.
	if err := templates.ExecModule(fd,
		fmt.Sprintf("%s/%s", domain, p), string(v), r,
		d, f); err != nil {
		return err
	}

	return nil
}
