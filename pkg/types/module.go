package types

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
}
