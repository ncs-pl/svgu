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

// HTML templates to generate for each module registered in the index.

package main

import (
	"fmt"
	"html/template"
	"io"
)

// HTML document templates
var (
	// indexTmpl is the HTML template for the root of the static site.
	// It redirects to a site configured by the user.
	//
	// {{.Count}} corresponds to the timer—in seconds— before redirecting.
	// {{.Redirect}} is the URL to redirect visitors to.
	indexTmpl = template.Must(
		template.New("index").Parse(`<!DOCTYPE html>
<meta charset="UTF-8">
<meta http-equiv="refresh" content="{{.Count}}; url='{{.Redirect}}'">
<meta name="viewport" content="width=device-width,initial-scale=1">
<p>Nothing to see, redirecting <a href="{{.Redirect}}">here</a>.`))

	// moduleTmpl is the HTML template for module-specific documents.
	// They include information for the Go toolchain to find and download
	// the module source code. It also redirects users to the module's
	// documentation on pkg.go.dev.
	//
	// {{.Prefix}} is the module's prefix, a.k.a. import path,
	// e.g. example.com/foo
	// {{.Vcs}} is the version control system used in the codebase
	// {{.Home}} is the repository's home.
	// {{.Dir}} is a URL template to a page listing the files inside a package
	// {{.File}} is a URL template listing the lines of a file
	// {{.Doc}} is the URL of the module's documentation on pkg.go.dev
	//
	// Templates support a specific set of substitutions which are documented
	// here: https://github.com/golang/gddo/wiki/Source-Code-Links
	moduleTmpl = template.Must(
		template.New("module").Parse(`<!DOCTYPE html>
<meta charset="UTF-8">
<meta name="go-import" content="{{.Prefix}} {{.Vcs}} {{.Home}}">
<meta name="go-source" content="{{.Prefix}} {{.Home}} {{.Dir}} {{.File}}">
<meta name="viewport" content="width=device-width,initial-scale=1">
<p>There is nothing to see, redirecting <a href="{{.Doc}}">here</a>.`))
)

// ExecIndex constructs an HTML document for the index of the generated site
// in the given writer "w."
// The document redirects visitor to the specified "url" after "count" seconds.
func ExecIndex(w io.Writer, url string, count uint8) error {
	return indexTmpl.Execute(w, struct {
		Redirect string
		Count    uint8
	}{
		Redirect: url,
		Count:    count,
	})
}

// ExecModule constructs an HTML document for a module indexed on the domain
// in the given writer "w." The "prefix" corresponds to the import path of the
// module, "vcs" to the version control system used — git, bazaar..., "home"
// is the repository's home and "dir"/"file" are URL templates as documented
// by GoDoc: https://github.com/golang/gddo/wiki/Source-Code-Links.
func ExecModule(w io.Writer, prefix, vcs,
	home, dir, file string) error {
	return moduleTmpl.Execute(w, struct {
		Prefix string
		Vcs    string
		Home   string
		Dir    string
		File   string
		Doc    string
	}{
		Prefix: prefix,
		Vcs:    vcs,
		Home:   home,
		Dir:    dir,
		File:   file,
		Doc:    fmt.Sprintf("https://pkg.go.dev/%s", prefix),
	})
}
