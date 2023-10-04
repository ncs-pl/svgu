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

title: Example
description: |
    An example and guide to setup and use SVGU to host a list of Go modules
    on a given domain name.
---

# Example

This document is a guide to setting up SVGU, configuring an index and hosting
it on a Web server.

A production-ready example can be found for [go.nc0.fr](https://go.nc0.fr)
which works thanks to SVGU. The source repository is available on
[GitHub](https://github.com/nc0fr/gomods).

## Getting Started

To get started, you need to make sure you have [installed SVGU](./install.mdx).
Also, you will need a [text editing software](https://vim.org),
a web server (you can either host one yourself or use a service like
[Vercel](https://vercel.com) or
[Cloudflare Pages](https://pages.cloudflare.com), and a domain name with access
to your DNS.

## Overview

In this guide, we will use the `example.com` domain as our target.
We will assume the existence of three Go modules, `example.com/a` hosted on
GitHub (Git), `example.com/b` hosted on SourceHut (Mercurial), and
`example.com/c` hosted on a private server with Bazaar.

Remember that Go supports [Bazaar][bzr], [Subversion][svn], [Git][git],
[Fossil][fossil], and [Mercurial][hg].

[bzr]: https://www.gnu.org/software/bazaar/
[fossil]: https://www2.fossil-scm.org/home/doc/trunk/www/index.wiki
[git]: https://git-scm.com
[hg]: https://www.mercurial-scm.org
[svn]: https://subversion.apache.org

## Writing the Configuration

Let's create a file named `svgu.star` in a directory of your choice.
This file will serve as our index configuration.

### Initialization

Inside this file, we will initialize the index and register our three modules.
Let's start with the index:

```python
# svgu.star

index(domain="example.com")
```

The [`index()`](./references.mdx#index) is responsible for the initialization
of the index. It takes one argument (named `domain`) which corresponds to the
domain name the registry will be published at.
Since we have decided to use `example.com` in this guide, we set `domain` to
`example.com`.

### Modules

Next, we need to declare our three modules, `example.com/{a,b,c}`.
This can be achieved with the help of the
[`module()`](./references.mdx#module).

```python
# svgu.star

index(domain="example.com")

# Module: example.com/a
# Hosted on GitHub with Git: https://github.com/example/a
# Share the "main" branch.
module(
    name="a",
    vcs="git",
    repo="https://github.com/example/a",
    dir="https://github.com/example/a/tree/main{/dir}",
    file="https://github.com/example/a/blob/main{/dir}/{file}#{line}",
)

# Module: example.com/b
# Hosted on Source Hut with Mercurial: https://hg.sr.ht/~example/b
# Share the "master" revision.
module(
    name="b",
    vcs="hg",
    repo="https://hg.sr.ht/~example/b",
    dir="https://hg.sr.ht/~example/b/browse{/dir}?rev=master",
    file="https://hg.sr.ht/~example/b/browse{/dir}/{file}?rev=master#L{line}",
)

# Module: example.com/c
# Private Bazaar repository.
# Share the "latest" revision of the "master" branch.
module(
    name="c",
    vcs="bzr",
    repo="https://private.example.com/c",
    dir="https://private.example.com/c/master{/dir}?rev=latest",
    file="https://private.example.com/c/master{/dir}/{file}?rev=latest",
)
```

Here we can see our three Go modules declared successfully.

### Helpers

While you can write as many `module()` as required, there is a lot boilerplate
involved in declaring a module. Luckily, SVGU use the
[Starlark configuration language](./starlark.mdx) which helps with that.

Indeed, imagine you need to declare multiple repositories, all hosted on your
public GitHub profile `example`. They all publish their `main` branch.

You could write as many `module()` as needed, but this will be a long and
tedeous task, and it will be hard to maintain.
Instead, you can write a helper function:

```python
# svgu.star

index(domain="example.com")

# To avoid mistyping the VCS, we can make a constant.
GIT = "git"

def github(name):
    """
    GitHub declares and registers a Go module hosted on GitHub
    on the `example` account.
    """

    # The "%" operator is a shortcut for the ".format()" function.
    # It allows variable substitution in strings.
    repo = "https://github.com/example/%s" % name
    dir = "%s/tree/main{/dir}" % repo

    module(
        name=name,
        vcs=GIT,
        repo=repo,
        dir=dir,
        file="%s/blob/main{/dir}/{file}#{line}" % repo
    )

# example.com/{d,e,f,g,h,i} all hosted on GitHub.
[ module(repo) in ("d", "e", "f", "g", "h", "i") ]
```

See how much easier it is to write and maintain!
Starlark is a programming language, you can use its properties to assist you
in your tasks, so make sure to use it.

### Linting

You can format and lint your configuration using the
[Buildifier](https://github.com/bazelbuild/buildtools).

## Going to Production

### Compiling

Now that your configuration is ready, you may generate your index as a set
of HTML document.

The `svgu` command-line tool accepts two parameters:

- `-c=file` the configuration file to use, here it will be `svgu.star`;
- `-o=directory` the output directory.

We want to generate our registry inside a directory called `out`.
Therefore, we can run:

```bash
$ svgu -c=svgu.star -o=out
```

If everything runs well, you should end with a new directory `out/` containing
HTML documents:

```bash
$ ls out
a.html  b.html  c.html  index.html
```

### Deploying

Now that you have everything ready inside your `out` directory, the last step
is to deploy the directory's content on the Web server of your choice.

Make sure your Web server has URL rewriting enabled, such that the `.html`
extension is not required in the request URL (some services, including
[GitHub Pages](https://pages.github.io) and
[Cloudflare Pages](https://pages.cloudflare.com) have this enabled by default,
traditional servers may require configuration:
[Caddy](https://caddyserver.com/docs/caddyfile/directives/rewrite),
[NGINX](https://www.nginx.com/blog/creating-nginx-rewrite-rules)...).

Finally, make sure you have correctly configured your DNS entries, and you
should quickly be able to import your Go modules using your domain name!

## Additional Notes

By default, the standard Go toolchain uses Google's Go module proxy to retrieve
modules. The proxy ensures the availablity of modules for all users by caching
Go modules.
However, due to the cache, you may experience delays and differences
between the latest pushed version and the imported one. These issues are
related to Google's proxy and are usually not an issue as the cache is often
updated in minutes.

Additionnally, Google's proxy may cause a lot of requests on your Web server,
due to the *Go Team at Google*'s choice to update the cache quickly (and
therefore avoid stale data).
You may opt-out the proxy by
[contacting a maintainer](https://github.com/golang/go/issues/new).

> Consider using a more powerful Web server or machine if needed, as the proxy
> is a valuable tool.

In general, please read [proxy.golang.org](https://proxy.golang.org) for
everything related to the Go proxy and module index.

