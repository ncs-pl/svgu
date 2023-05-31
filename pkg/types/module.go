package types

import (
	"go.nc0.fr/svgu/pkg/templates"
	"os"
	"path"
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
	lock sync.Mutex
}

// GenerateFile generates the index file.
func (m *Module) GenerateFile(out string, domain string) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	f := path.Join(out, m.Path+".html")

	// Create the file.
	fd, err := os.Create(f)
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
	if err := templates.ExecModule(fd, domain+m.Path, string(m.Vcs),
		m.Repo, m.Dir, m.File); err != nil {
		return err
	}

	return nil
}
