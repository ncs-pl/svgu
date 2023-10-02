# Getting started

This document will help you get started with SVGU.

## Table of contents

- [Installation](#installation)
  + [From source](#compiling-from-source)
  + [Pre-built binaries](#pre-built-binaries)
  + [Docker](#docker)
- [Usage](#usage)
  + [Configuration](#configuration)
  + [Generating the HTML files](#generating-the-html-files)
  + [Hosting the HTML files](#hosting-the-html-files)
- [Advanced configuration](#advanced-configuration)

## Installation

First, you need to install SVGU. There are several ways to do this.

### Compiling from source

SVGU is a Go program, and can be installed quickly using `go install`:

```shell
# Change "latest" to the Git ref you want to install.
$ go install go.nc0.fr/svgu/cmd/svgu@latest
```

> Note: SVGU requires Go 1.16 or later.

### Pre-built binaries

TODO(nc0): Add pre-built binaries and pkg for various platforms.

### Docker

TODO(nc0): Add Docker image.

## Usage

SVGU is a command-line tool that, given a configuration file, will generate
a set of HTML 5 documents ready to be served by any web server.

### Configuration

The configuration file is a [Starlark][starlark] script that registers the
various modules to publish under a domain.

Starlark is a dialect of Python, and is very easy to learn. The reference
documentation is available [here][starlark].

In addition to the standard Starlark specification, SVGU provides a set of
utility functions to make it easier to write configuration files.
The documentation for these functions is available [here](references.md).

To get started, you need to create a file named `DOMAINS.star` in the directory
of your choice. This file will contain the configuration for your website.

> Note: The name of the file is not important, by default SVGU will look for
> a file named `DOMAINS.star` in the current directory. You can specify a
> different file using the `-c` flag.

In this file, you need to create an index (a set of modules on a domain) and
registers various modules to it. For example, let's say you want to publish
the module `foo` on the domain `example.com`. You would write the following:

```starlark
# DOMAINS.star
index(domain = "example.com")

module(
    name = "foo",
    vcs = "git",
    repo = "https://github.com/example/foo.git",
    dir = "https://github.com/example/foo.git/tree/master{/dir}",
    file = "https://github.com/example/foo.git/blog/master{/dir}/{file}#L{line}",
)
```

This may look a bit complicated, but it's actually quite simple. Let's break
it down:

- `index(domain = "example.com")` creates an index—a set of modules—on the
  domain `example.com`. The domain is used to generate the URLs for the
  modules, e.g. `example.com/foo`.
- `module(name = "foo", ...)` registers a module named `foo` on the index.
  The `name` argument is the name of the module, and is used to generate
  the URLs for the module, e.g. `example.com/foo`.
- `vcs = "git"` tells SVGU that the module is hosted on a Git repository.
- `repo = "..."` is the URL of the Git repository.
- `dir = "..."` and `file = "..."` are URL templates for
  [pkg.go.dev](https://pkg.go.dev). They are used to allow automatic
  documentation. You can read more about them in the
  [reference](references.md#module).

You can add as many modules as you want to an index. For example, let's say
you want to add the module `bar` hosted on a Mercurial repository:

```starlark
module(
    name = "bar",
    vcs = "hg",
    repo = "https://example.com/bar",
    dir = "https://example.com/bar/file/tip{/dir}",
    file = "https://example.com/bar/file/tip{/dir}/{file}#L{line}",
)
```

### Generating the HTML files

When you're done, you can run SVGU to generate the HTML files:

```shell
$ svgu
```

This will generate a directory named `out` containing the HTML files. You can
then serve this directory using any web server.
If you used another name for the configuration file, you can specify it using
the `-c` flag.
You can also change the output directory using the `-o` flag.
And lastly, the `-v` flag can be used to print more information.

### Hosting the HTML files

Once you have the generated HTML documents inside the output directory
*(by default it is `dst`)*, you can upload them to your web server.

The target web server does not need to be configured in any special way,
we only require that it supports URL rewriting (as most web servers do)
to remove the `.html` extension from the URLs.
This is sometimes called [clean URLs](https://en.wikipedia.org/wiki/Clean_URL)
or "Pretty URLs".

Here is a non-exhaustive list of web servers that support URL rewriting:

- [Apache HTTPD](https://httpd.apache.org/docs/current/rewrite/remapping.html)
- [Caddy](https://caddyserver.com/docs/caddyfile/directives/rewrite)
- Cloudflare Pages does it automatically
- [NGINX](https://www.nginx.com/blog/creating-nginx-rewrite-rules/)
- [GitHub Pages](https://pages.github.com/) does it by default
- ...

## Advanced configuration

As you can see, the configuration file is very simple. However, it is also
a bit boilerplate-heavy, especially with the template URLs for `dir` and
`file`.

To avoid this, SVGU provides a set of utility libraries that can be imported
via the `load()` function. For example, the previous configuration can be
shortened to:

```starlark
load("@svgu/git.star", "git")

index(domain = "example.com")

git.github(
  name = "foo",
  user = "example",
  repo = "foo",
  branch = "master",
)

...
```

You can read more about the utility libraries in the [reference](references.md).
Those can help when writing long configuration files.

[starlark]: https://github.com/bazelbuild/starlark
