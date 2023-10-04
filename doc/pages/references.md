---
# Copyright Nicolas (2023)
#
# * Nicolas Paul
#
# This software is a computer program whose purpose is to allow the hosting
# and sharing of Go modules using a personal domain.
#
# This software is governed by the CeCILL license under French law and
# abiding by the rules of distribution of free software.  You can  use, 
# modify and/ or redistribute the software under the terms of the CeCILL
# license as circulated by CEA, CNRS and INRIA at the following URL
# "http://www.cecill.info". 
#
# As a counterpart to the access to the source code and  rights to copy,
# modify and redistribute granted by the license, users are provided only
# with a limited warranty  and the software's author,  the holder of the
# economic rights,  and the successive licensors  have only  limited
# liability. 
#
# In this respect, the user's attention is drawn to the risks associated
# with loading,  using,  modifying and/or developing or reproducing the
# software by the user in light of its specific status of free software,
# that may mean  that it is complicated to manipulate,  and  that  also
# therefore means  that it is reserved for developers  and  experienced
# professionals having in-depth computer knowledge. Users are therefore
# encouraged to load and test the software's suitability as regards their
# requirements in conditions enabling the security of their systems and/or 
# data to be ensured and,  more generally, to use and operate it in the 
# same conditions as regards security.
#
# The fact that you are presently reading this means that you have had
# knowledge of the CeCILL license and that you accept its terms.

title: References
description: |
    References and documentation of the SVGU public API.
---

# SVGU References

This document presents the exposed API of SVGU for use within Starlark
configurations.

## index

```python
index(domain)
```

Index initializes the module registry to host at the given domain.
This function must be called before any [`module()`](#module) function.

Parameters:

- `domain` (str): the domain name the repository will be hosted at, e.g.
  [go.nc0.fr](https://go.nc0.fr).

```python
# Modules hosted at example.com`
index(domain="example.com")

# ...
```

## module

```python
module(name, vcs, repo, dir, file)
```

Module registers a module inside the index. A module is defined by its `name`.
Two modules with the same name cannot exist at the same time in the same index.

A module name must be a valid URL path segment, it cannot be empty, nor
contains dots (`.`), slashes (`/`), backslashes (`\`), colons (`:`),
asterisks (`*`), or any ASCII control characters (i.e. `\x00` through `\x1f`)
(required to avoid issues with filesystems and URLs).

The module name cannot also be simply `index`. This is a limitation of the
file generation system.

Parameters:

- `name` (str): the name of the module (without the domain and the prefixing
  slash), e.g. `"foo/v2"`;
- `vcs` (str): the version control system used to host the code supported by
  the standard Go toolchain. Must be one of `"bzr"` ([Bazaar][bzr]), `"fossil"`
  ([Fossil][fossil]), `"git"` ([Git][git]), `"hg"` ([Mercurial][hg]), or
  `"svn"` ([Subversion][svn]);
- `repo`: the URL of the code repository, e.g. `"https://github.com/foo/bar"`;
- `dir`: a Go template forming the URL to a page listing the various files
  inside a directory, e.g. `"https://github.com/foo/bar/tree/main{/dir}"`.
  Please refers to the [`go-source`][go-source] for more information about the
 - `file`: a Go template forming the URL to a page listing the lines of code
  inside a file, and links to a specific line if needed, e.g. 
  `"https://github.com/foo/bar/blob/main{/dir}/{file}#{line}"`. Please refers to
  the [`go-source`][go-source] for more information about the
  template template.

[go-source]: https://github.com/golang/gddo/wiki/Source-Code-Links
[bzr]: https://www.gnu.org/software/bazaar/
[fossil]: https://www2.fossil-scm.org/home/doc/trunk/www/index.wiki
[git]: https://git-scm.com
[hg]: https://www.mercurial-scm.org
[svn]: https://subversion.apache.org

Here are valid templates for a few popular hosting services:

- **GitHub (Git):** `https://<instance>/<user>/<repo>/tree/<ref>{/dir}` (dir)
  and`https://<instance>/<user>/<repo>/blob/<ref>{/dir}/{file}#{line}` (file);
- **GitLab (Git):** `https://<instance>/<user>/<repo>/-/tree/<ref>{/dir}` (dir)
  and `https://<instance>/<user>/<repo>/-/blob/<ref>{/dir}/{file}#L{line}`
  (file);
- **Source Hut (Git):** `https://<instance>/~<user>/<repo>/tree/<ref>{/dir}`
  (dir) and
  `https://<instance>/~<user>/<repo>/tree/<ref>/item{/dir}/{file}#L{line}`
  (file);
- **Bitbucket (Git):** `https://<instance>/<workspace>/<repo>/src/<ref>{/dir}`
  (dir) and
  `https://<instance/<workspace>/<repo>/src/<ref>{/dir}/{file}#{file}-{file}`
  (file);
- **Gitiles (Git):** `https://<instance>/<repo>/+/refs/heads/<ref>{/dir}` (dir)
  and `https://<instance>/<repo>/+/refs/heads/<ref>{/dir}/{file}#{line}` (file);
- **Launchpad (Bazaar):**
  `https://<instance>/~<user>/<repo>/<branch>/files/<rev>{/dir}` (dir) and
  `https://<instance>/~<user>/<repo>/<branch>/view/<rev>{/dir}/{file}#L{line}`
  (file);
- **Source Hut (Mercurial):**
  `https://<instance>/~<user>/<repo>/browse{/dir}?rev=<rev>` (dir) and
  `https://<instance>/~<user>/<repo>/browse{/dir}/{file}?rev=<rev>#L{line}`
  (file).

> Do not forget to replace `<variables>` with the expected value.

```python
index(domain="example.com")

SUBVERSION = "svn"

# Will be usable in Go as "example.com/foo".
module(name="foo", vcs=SUBVERSION, repo="https://src.example.com/foo",
        dir="https://src.example.com/tree/master{/dir}",
        file="https://src.example.com/tree/master{/dir}/{file}#{line}")
```

## Helpers

In previous versions of SVGU, a set of Starlark modules (`@svgu/bzr`...) was
exposing a set of macros to help writing configurations for modules hosted
on popular platforms, such as GitHub.

We decided to remove them due to the burden of maintaining them with the
rational that they are trivial to write. Indeed, the old `git.github()` macro
can be implemented as such:

```python
def github(name, user, repo, ref = "main", instance = "github.com"):
    repo = "https://%s/%s/%s" % (instance, user, repo)
    dir = "%s/tree/%s{/dir} % (repo, ref)

    module(
        name=name,
        vcs="git",
        repo=repo,
        dir=dir,
        file="%s/{file}#{line}" % dir,
    )
```

> The same applies to other macro inside other modules, such as
>`bzr.launchpad()` and `hg.sourcehut()`.

And constants, such as `git.GIT` and `fossil.FOSSIL` are simple variables:

```python
BAZAAR = "bzr"
FOSSIL = "fossil"
GIT = "git"
MERCURIAL = "hg"
SUBVERSION = "svn"
```

It is encouraged to write your own helper functions and constants at the
beginning of your configuration file.

