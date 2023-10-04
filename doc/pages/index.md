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

title: Introduction
description: |
    SVGU is a utility tool allowing the sharing and publication of Go modules
    on personal domains easily, in a declarative manner.
---

# Welcome to SVGU
 
**SVGU** (short for *shared vanity Go URLs*) is a neat command-line utility
that allows anyone to publish their Go modules on their own domain, to
obtain names such as `example.com/foo` instead of
`github.com/example/foo`.

Having a Go module with a custom domain name avoid being trapped to the code
hosting service (such as GitHub). Indeed, you could move to another host or
change the origin without requiring your users to update all their code
to the new path.

For illustration, imagine you are using `example.com/foo` as a dependency,
whose main origin is on GitHub. If suddenly the origin is moved to GitLab
(for some particular reason), it will be transparent to you.

## How Does it Work?

SVGU works by generating a set of HTML files containing the
required meta tags (incl.
[`go-source`](https://github.com/golang/gddo/wiki/Source-Code-Links) and
[`go-import`](https://go.dev/blog/publishing-go-modules)) by the standard Go
toolchain. These documents also redirects users to the
[Go documentation service](https://pkg.go.dev) for the requested module.

The resulting output directory can be hosted on any Web server, with the only
required configuration to rewrite URLs to remove the `.html` prefix (some
call this behavior "Pretty URLs").

> Here are links to various popular Web servers enabling URL rewriting:
> [Apache HTTPD](https://httpd.apache.org/docs/current/rewrite/remapping.html),
> [Caddy](https://caddyserver.com/docs/caddyfile/directives/rewrite),
> [NGINX](https://www.nginx.com/blog/creating-nginx-rewrite-rules/),
> [Cloudflare Pages](https://pages.cloudflare.com) and
> [GitHub Pages](https://pages.github.com/) do this by default,
> ...

## Installation

If you do not see your operating system, use the [Go](#go) or
[compiling from source](#from-source) methods. Also consider contributing
to add your own package manager installation process!

### Homebrew

Nicolas Paul maintains an external [Homebrew](https://brew.sh) repository
allowing the installation of SVGU on macOS and some GNU/Linux systems.

```bash
$ brew install nc0fr/nc0/svgu
```

### Go

The standard Go toolchain can download and install executables directly.
You only need [Go 1.16+](https://go.dev).

```bash
$ # You can replace latest with any Git ref needed.
$ go install go.nc0.fr/svgu@latest
```

### From Source

Compiling SVGU from source requires [Git](https://git-scm.com) and
[Go](https://go.dev) (1.16 or more).

First, clone the repository from our GitHub origin:

```bash
$ git clone https://github.com/nc0fr/svgu.git
$ cd svgu
```

Then, build the Go software using `make`:

```bash
$ make
$ ./svgu -h
```

## Licensing

SVGU is a free software, available under the
[CeCILL 2.1 license](https://cecill.info) contract.
This documentation is available under the
[Creative Commons Attribution-ShareAlike 4.0 license](https://creativecommons.org/licenses/by-sa/4.0/).

Please see the [repository](https://github.com/nc0fr/svgu) for complete
details.
