// Copyright (c) 2023 Nicolas Paul All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// Config represents the content of the `vanity.yaml` configuration file.
//
// Hostname is the used domain, e.g. `go.example.com`.
//
// Paths describes the various import paths to generate.
type Config struct {
	Hostname string `yaml:"hostname"`
	Paths    []Path `yaml:"paths"`
}

// Path represents an import path.
//
// Repository is the repository URL as it would appear in `go-import` meta tags
// (https://golang.org/cmd/go/#hdr-Remote_import_paths).
//
// Vcs marks the version control system used on the defined repository.
//
// Dir is the URL template for a page listing the files in the package.
// Available substitutions: `{dir}`, `{/dir}`.
//
// File is the URL template for a link to a line in a source file.
// Available substitutions: `{file}`, `{line}`.
//
// Prefix represents the URL path in the total import path, without a
// leading `/`, e.g. `foo`.
//
// Read more about substitutions: https://github.com/golang/gddo/wiki/Source-Code-Links
type Path struct {
	Dir        string `yaml:"dir"`
	File       string `yaml:"file"`
	Prefix     string `yaml:"prefix"`
	Repository string `yaml:"repository"`
	Vcs        VCS    `yaml:"vcs"`
}

// VCS represents a version control system for a Go repository.
type VCS string

const (
	VcsBazaar     VCS = "bzr"
	VcsFossil     VCS = "fossil"
	VcsGit        VCS = "git"
	VcsMercurial  VCS = "hg"
	VcsSubversion VCS = "svn"
)
