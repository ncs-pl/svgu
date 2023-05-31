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
