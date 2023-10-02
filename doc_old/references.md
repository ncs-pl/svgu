# References

This document contains a list of references and resources related to SVGU's
Starlark environment.

The sources for the environment are located inside the
[pkg/config](../pkg/config) directory, especially the
[lib](../pkg/config/lib) package.

## Starlark

[Starlark][starlark-link] is a dialect of Python designed for use as a
configuration language. It is simple, safe, and expressive.
Originally called Skylark, Starlark is commonly used with
[Bazel][bazel-link]-like build systems, but it is not tied to any
particular tool or project.

SVGU's Starlark environment is based on the
[Starlark Go][starlark-go] implementation open sourced by Google.

## Table of Contents

- [Prelude](#prelude)
    + [index](#index)
    + [module](#module)
- [Bazaar](#bazaar)
  + [bzr.BAZAAR](#bzrbazaar)
  + [bzr.LAUNCHPAD_DEFAULT_INSTANCE](#bzrlaunchpad_default_instance)
  + [bzr.LAUNCHPAD_DEFAULT_REV](#bzrlaunchpad_default_rev)
  + [brz.LAUNCHPAD_DEFAULT_BRANCH](#brzlaunchpad_default_branch)
  + [bzr.launchpad](#bzrlaunchpad)
- [Fossil](#fossil)
    + [fossil.FOSSIL](#fossilfossil)
- [Git](#git)
  + [git.GIT](#gitgit)
  + [git.GITHUB_DEFAULT_INSTANCE](#gitgithub_default_instance)
  + [git.GITHUB_DEFAULT_REF](#gitgithub_default_ref)
  + [git.SOURCEHUT_DEFAULT_INSTANCE](#gitsourcehut_default_instance)
  + [git.SOURCEHUT_DEFAULT_REF](#gitsourcehut_default_ref)
  + [git.GITLAB_DEFAULT_INSTANCE](#gitgitlab_default_instance)
  + [git.GITLAB_DEFAULT_REF](#gitgitlab_default_ref)
  + [git.BITBUCKET_DEFAULT_INSTANCE](#gitbitbucket_default_instance)
  + [git.BITBUCKET_DEFAULT_REF](#gitbitbucket_default_ref)
  + [git.GITILES_DEFAULT_REF](#gitgitiles_default_ref)
  + [git.github](#gitgithub)
  + [git.sourcehut](#gitsourcehut)
  + [git.gitlab](#gitgitlab)
  + [git.bitbucket](#gitbitbucket)
  + [git.gitiles](#gitgitiles)
- [Mercurial](#mercurial)
  + [hg.MERCURIAL](#hgmercurial)
  + [hg.SOURCEHUT_DEFAULT_INSTANCE](#hgsourcehut_default_instance)
  + [hg.SOURCEHUT_DEFAULT_REV](#hgsourcehut_default_rev)
  + [hg.sourcehut](#hgsourcehut)
- [Subversion](#subversion)
    + [svn.SUBVERSION](#svnsubversion)

## Prelude

The [prelude](../pkg/config/lib/prelude/prelude.go) module contains a set of
built-in functions.

These functions are the only required symbols to use SVGU in the configuration
file.

Additional functions are helpers that can be used to simplify the configuration
file. For instance, [git.github](#gitgithub) will automatically construct the
template URLs for `dir` and `file` based on the repository's home URL.

### index

Initialize the index of module.
Should be called before any [module](#module) operations.

`index(domain)`

#### Parameters

| Name     | Type     | Description                                               |
|----------|----------|-----------------------------------------------------------|
| `domain` | `string` | The domain name of the repository, e.g. `go.example.com`. |

#### Example

```starlark
index(domain = "go.example.com")

# ...
```

### module

Register a module in the index.
Note that you cannot register multiple modules with the same name.

> A module name cannot be empty, `.`, `..`, `\`, `/`, `:`, `*`, `?`, `"`, `<`,
> `>`, `|`, or any ASCII control character (i.e. `\x00` through `\x1f`).
> This is required to avoid problems with file systems.
>
> The module name cannot also be simply `index`. This is a limitation of the
> tool since it generates a file named `index.html` in the root of the output
> directory.
>
> To see the full list of reserved names, see the
> [invalidName](https://git.sr.ht/~n1c00o/svgu/tree/master/item/pkg/config/lib/prelude/prelude.go#L11)
> variable.

#### Parameters

| Name   | Type     | Description                                                                                                                                                                                                                                      |
|--------|----------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `name` | `string` | The name of the module.                                                                                                                                                                                                                          |
| `vcs`  | `string` | The name of the Version Control System (VCS) used by the module. It should be one of `"bzr"`, `"fossil"`, `"git"`, `"hg"`, or `"svn"`.                                                                                                           |
| `repo` | `string` | The URL of the repository's home, e.g. `https://src.example.com/foo`.                                                                                                                                                                            |
| `dir`  | `string` | A template URL to a page listing the various files inside a directory, e.g. `https://src.example.com/foo/tree/master{/dir}`. See the [go-source][go-source-tag] meta tag's documentation for more information about substitution variables.      |
| `file` | `string` | A template URL to a page listing the lines inside a Go file, e.g. `https://src.example.com/foo/tree/master{/dir}/{file}#L{line}`. See the [go-source][go-source-tag] meta tag's documentation for more information about substitution variables. |

#### Example

```starlark
index(domain = "go.example.com")

# Will be available as `go.example.com/foo`.
module(
    name = "foo",
    vcs = "git",
    repo = "https://src.example.com/foo",
    dir = "https://src.example.com/foo/tree/master{/dir}",
    file = "https://src.example.com/foo/tree/master{/dir}/{file}#L{line}",
)

# Will be available as `go.example.com/bar/baz`.
module(
    name = "bar/baz",
    vcs = "svn",
    repo = "https://src.example.com/bar/baz",
    dir = "https://src.example.com/bar/baz{/dir}",
    file = "https://src.example.com/bar/baz{/dir}/{file}#L{line}",
)
```

## Bazaar

The [bzr](../pkg/config/lib/bzr/bzr.star) (Bazaar) module contains a set
of utilities to work with modules hosted on
[GNU Bazaar][bzr-link] repositories.

The module can be imported inside your configuration file by adding the
following code at the beginning of the file:

```starlark
load("@svgu/bzr", "bzr")
```

### bzr.BAZAAR

`"bzr"` \
A constant containing the name of the VCS as required by the
[`module`](#module) function.

#### Example

```starlark
module(
    ...
    vcs = bzr.BAZAAR,
)
```

### bzr.LAUNCHPAD_DEFAULT_INSTANCE

`"https://bazaar.launchpad.net` \
A constant containing the default instance of the
[Launchpad Bazaar hosting][launchpad-bzr-link] service.

### bzr.LAUNCHPAD_DEFAULT_REV

`"head:"` \
A constant containing the default revision number to use when the repository
is hosted on [Launchpad Bazaar hosting][launchpad-bzr-link] service.

### bzr.LAUNCHPAD_DEFAULT_BRANCH

`"trunk"` \
A constant containing the default branch name to use when the repository
is hosted on [Launchpad Bazaar hosting][launchpad-bzr-link] service.

### bzr.launchpad

A macro that registers a module hosted on
[Launchpad Bazaar hosting][launchpad-bzr-link] service.

#### Parameters

| Name       | Type     | Description                                                                                                       |
|------------|----------|-------------------------------------------------------------------------------------------------------------------|
| `name`     | `string` | The name of the module.                                                                                           |
| `user`     | `string` | The name of the user or organization that owns the repository.                                                    |
| `repo`     | `string` | The name of the repository.                                                                                       |
| `branch`   | `string` | The name of the branch. Defaults to [`bzr.LAUNCHPAD_DEFAULT_BRANCH`](#bzrlaunchpad_default_branch)                |
| `rev`      | `string` | The revision number. Defaults to [`bzr.LAUNCHPAD_DEFAULT_REV`](#bzrlaunchpad_default_rev)                         |
| `instance` | `string` | The URL of the Launchpad instance. Defaults to [`bzr.LAUNCHPAD_DEFAULT_INSTANCE`](#bzrlaunchpad_default_instance) |

#### Example

```starlark
load("@svgu/bzr", "bzr")

index(domain = "go.example.com")

# Will be available as `go.example.com/foo`.

bzr.launchpad(
    name = "foo",
    user = "bar",
    repo = "foo",
)
```

## Fossil

The [fossil](../pkg/config/lib/fossil/fossil.star) module contains a set
of utilities to work with modules hosted on
[Fossil][fossil-link] repositories.

The module can be imported inside your configuration file by adding the
following code at the beginning of the file:

```starlark
load("@svgu/fossil", "fossil")
```

### fossil.FOSSIL

`"fossil"` \
A constant containing the name of the VCS as required by the
[`module`](#module) function.

#### Example

```starlark
module(
    ...
    vcs = fossil.FOSSIL,
)
```

## Git

The [git](../pkg/config/lib/git/git.star) module contains a set of utilities
to work with modules hosted on [Git][git-link] repositories.

The module can be imported inside your configuration file by adding the
following code at the beginning of the file:

```starlark
load("@svgu/git", "git")
```

### git.GIT

`"git"` \
A constant containing the name of the VCS as required by the
[`module`](#module) function.

#### Example

```starlark
module(
    ...
    vcs = git.GIT,
)
```

### git.GITHUB_DEFAULT_INSTANCE

`"https://github.com"` \
A constant containing the default GitHub instance to use when the repository
is hosted on [GitHub][github-link].

### git.GITHUB_DEFAULT_REF

`"main"` \
A constant containing the default Git reference to use when the repository is
hosted on [GitHub][github-link].

See [GitHub's decision to rename master to main](https://github.com/github/renaming).

### git.SOURCEHUT_DEFAULT_INSTANCE

`"https://git.sr.ht"` \
A constant containing the default SourceHut instance to use when the repository
is hosted on [Source Hut's Git hosting][sourcehut-git-link].

### git.SOURCEHUT_DEFAULT_REF

`"master"` \
A constant containing the default Git reference to use when the repository is
hosted on [Source Hut's Git hosting][sourcehut-git-link].

### git.GITLAB_DEFAULT_INSTANCE

`"https://gitlab.com"` \
A constant containing the default GitLab instance to use when the repository
is hosted on [GitLab][gitlab-link].

### git.GITLAB_DEFAULT_REF

`"main"` \
A constant containing the default Git reference to use when the repository is
hosted on [GitLab][gitlab-link].

See [GitLab's renaming from master to main announcement](https://about.gitlab.com/blog/2021/03/10/new-git-default-branch-name/).

### git.BITBUCKET_DEFAULT_INSTANCE

`"https://bitbucket.org"` \
A constant containing the default Bitbucket instance to use when the repository
is hosted on [Bitbucket][bitbucket-link].

### git.BITBUCKET_DEFAULT_REF

`"master"` \
A constant containing the default Git reference to use when the repository is
hosted on [Bitbucket][bitbucket-link].

### git.GITILES_DEFAULT_REF

`"master"` \
A constant containing the default Git reference to use when the repository is
hosted on [Gitiles][gitiles-link] ([Gerrit][gerrit-link] with a Gitiles
front-end).

### git.github

A macro that registers a module hosted on [GitHub][github-link].

#### Parameters

| Name       | Type     | Description                                                                                          |
|------------|----------|------------------------------------------------------------------------------------------------------|
| `name`     | `string` | The name of the module.                                                                              |
| `user`     | `string` | The GitHub user or organization name.                                                                |
| `repo`     | `string` | The GitHub repository name.                                                                          |
| `ref`      | `string` | The Git reference to use. Default to [`git.GITHUB_DEFAULT_REF`](#gitgithub_default_ref).             |
| `instance` | `string` | The GitHub instance to use. Default to [`git.GITHUB_DEFAULT_INSTANCE`](#gitgithub_default_instance). |

#### Example

```starlark
load("@svgu/git", "git")

index(domain = "go.example.com")

# By default, the function assumes that the repository is hosted on
# https://github.com and the reference is `main`.
git.github(
    name = "foo",
    user = "example",
    repo = "foo",
)

# You can override the default reference and instance.
git.github(
    name = "bar",
    user = "example",
    repo = "bar",
    ref = "trunk",
    instance = "https://github.example.com",
)
```

### git.gitlab

A macro that registers a module hosted on [GitLab][gitlab-link].

#### Parameters

| Name       | Type     | Description                                                                                          |
|------------|----------|------------------------------------------------------------------------------------------------------|
| `name`     | `string` | The name of the module.                                                                              |
| `user`     | `string` | The GitLab user or organization name.                                                                |
| `repo`     | `string` | The GitLab repository name.                                                                          |
| `ref`      | `string` | The Git reference to use. Default to [`git.GITLAB_DEFAULT_REF`](#gitgitlab_default_ref).             |
| `instance` | `string` | The GitLab instance to use. Default to [`git.GITLAB_DEFAULT_INSTANCE`](#gitgitlab_default_instance). |

#### Example

```starlark
load("@svgu/git", "git")

index(domain = "go.example.com")

# By default, the function assumes that the repository is hosted on
# https://gitlab.com and the reference is `main`.
git.gitlab(
    name = "foo",
    user = "example",
    repo = "foo",
)

# You can override the default reference and instance.
git.gitlab(
    name = "bar",
    user = "example",
    repo = "bar",
    ref = "master",
    instance = "https://gitlab.example.com",
)
```

### git.sourcehut

A macro that registers a module hosted on
[Source Hut's Git hosting][sourcehut-git-link].

> Note: Source Hut's Git hosting is still in beta. Organization support is
> not yet available.

#### Parameters

| Name       | Type     | Description                                                                                                    |
|------------|----------|----------------------------------------------------------------------------------------------------------------|
| `name`     | `string` | The name of the module.                                                                                        |
| `user`     | `string` | The Source Hut user name.                                                                                      |
| `repo`     | `string` | The Source Hut repository name.                                                                                |
| `ref`      | `string` | The Git reference to use. Default to [`git.SOURCEHUT_DEFAULT_REF`](#gitsourcehut_default_ref).                 |
| `instance` | `string` | The Source Hut instance to use. Default to [`git.SOURCEHUT_DEFAULT_INSTANCE`](#gitsourcehut_default_instance). |

#### Example

```starlark
load("@svgu/git", "git")

index(domain = "go.example.com")

# By default, the function assumes that the repository is hosted on
# https://git.sr.ht and the reference is `master`.

git.sourcehut(
    name = "foo",
    user = "example",
    repo = "foo",
)

# You can override the default reference and instance.
git.sourcehut(
    name = "bar",
    user = "example",
    repo = "bar",
    ref = "main",
    instance = "https://git.srht.example.com",
)
```

### git.bitbucket

A macro that registers a module hosted on [Bitbucket][bitbucket-link].

It is likely that you want to override the default instance to include your
user in the URL.
To do so, you can use the `instance` parameter:

```starlark
git.bitbucket(
    ...
    instance = "https://user@bitbucket.org/example",
)
```

As always, your mileage may vary.

#### Parameters

| Name        | Type     | Description                                                                                                   |
|-------------|----------|---------------------------------------------------------------------------------------------------------------|
| `name`      | `string` | The name of the module.                                                                                       |
| `workspace` | `string` | The Bitbucket workspace ID.                                                                                   |
| `repo`      | `string` | The Bitbucket repository name.                                                                                |
| `ref`       | `string` | The Git reference to use. Default to [`git.BITBUCKET_DEFAULT_REF`](#gitbitbucket_default_ref).                |
| `instance`  | `string` | The Bitbucket instance to use. Default to [`git.BITBUCKET_DEFAULT_INSTANCE`](#gitbitbucket_default_instance). |

#### Example

```starlark
load("@svgu/git", "git")

index(domain = "go.example.com")

# By default, the function assumes that the repository is hosted on
# https://bitbucket.org and the reference is `master`.

git.bitbucket(
    name = "foo",
    workspace = "example",
    repo = "foo",
)

# You can override the default reference and instance.
git.bitbucket(
    name = "bar",
    workspace = "example",
    repo = "bar",
    ref = "main",
    instance = "https://bitbucket.example.com",
)
```

### git.gitiles

A macro that registers a module hosted on [Gitiles][gitiles-link]
([Gerrit][gerrit-link] with Gitiles front-end).

> Note: This is not a complete macro as I lack experience with Gerrit.
> Feel free to open a pull request if you want to improve it.

#### Parameters

| Name       | Type     | Description                                                                                 |
|------------|----------|---------------------------------------------------------------------------------------------|
| `name`     | `string` | The name of the module.                                                                     |
| `instance` | `string` | The Gitiles instance to use.                                                                |
| `repo`     | `string` | The Gitiles repository path.                                                                |
| `ref`      | `string` | The Git reference to use. Defaults to [`git.GITILES_DEFAULT_REF`](#gitgitiles_default_ref). |

#### Example

```starlark
load("@svgu/git", "git")

index(domain = "go.example.com")

# By default, the function assumes that the used reference is `master`.
git.gitiles(
    name = "go",
    instance = "https://go.googlesource.com",
    repo = "go",
)

# You can override the default reference.
git.gitiles(
    name = "go",
    instance = "https://go.googlesource.com",
    repo = "go",
    ref = "release-branch.go1.17",
)
```

## Mercurial

The [hg](../pkg/config/lib/hg/hg.star) (Mercurial) module contains a set
of utilities to work with modules hosted on
[Mercurial][hg-link] repositories.

The module can be imported inside your configuration file by adding the
following code at the beginning of the file:

```starlark
load("@svgu/hg", "hg")
```

### hg.MERCURIAL

`"hg"` \
A constant containing the name of the VCS as required by the
[`module`](#module) function.

#### Example

```starlark
module(
    ...
    vcs = hg.MERCURIAL,
)
```

### hg.SOURCEHUT_DEFAULT_INSTANCE

`"https://hg.sr.ht"` \
A constant containing the default instance to use when the repository is hosted
on [Source Hut Mercurial][sourcehut-hg-link] hosting.

### hg.SOURCEHUT_DEFAULT_REV

`"tip"` \
A constant containing the default revision (branch, ...) to use when the
repository is hosted on [Source Hut Mercurial][sourcehut-hg-link] hosting.

### hg.sourcehut

A macro that registers a module hosted on
[Source Hut Mercurial][sourcehut-hg-link] hosting.

> Note: Source Hut's Mercurial hosting is still in beta. Organization support
> is not yet available.

#### Parameters

| Name       | Type     | Description                                                                                                  |
|------------|----------|--------------------------------------------------------------------------------------------------------------|
| `name`     | `string` | The name of the module.                                                                                      |
| `user`     | `string` | The Source Hut user.                                                                                         |
| `repo`     | `string` | The Source Hut repository name.                                                                              |
| `instance` | `string` | The Source Hut instance to use. Default to [`hg.SOURCEHUT_DEFAULT_INSTANCE`](#hgsourcehut_default_instance). |
| `rev`      | `string` | The revision (branch, ...) to use. Default to [`hg.SOURCEHUT_DEFAULT_REV`](#hgsourcehut_default_rev).        |

#### Example

```starlark
load("@svgu/hg", "hg")

index(domain = "go.example.com")

# By default, the function assumes that the repository is hosted on
# https://hg.sr.ht and the revision is `tip`.

hg.sourcehut(
    name = "foo",
    user = "example",
    repo = "foo",
)

# You can override the default revision and instance.
hg.sourcehut(
    name = "bar",
    user = "example",
    repo = "bar",
    rev = "default",
    instance = "https://hg.example.com",
)
```

## Subversion

The [svn](../pkg/config/lib/svn/svn.star) (Subversion) module contains
a set of utilities to work with modules hosted on
[Apache Subversion][svn-link] repositories.

The module can be imported inside your configuration file by adding the
following code at the beginning of the file:

```starlark
load("@svgu/svn", "svn")
```

### svn.SUBVERSION

`"svn"` \
A constant containing the name of the VCS as required by the
[`module`](#module) function.

#### Example

```starlark
module(
    ...
    vcs = svn.SUBVERSION,
)
```

[svn-link]: https://subversion.apache.org/

[hg-link]: https://www.mercurial-scm.org/

[git-link]: https://git-scm.com/

[fossil-link]: https://www.fossil-scm.org/

[bzr-link]: https://bazaar.canonical.com/

[bitbucket-link]: https://bitbucket.org

[github-link]: https://github.com/about

[sourcehut-git-link]: https://man.sr.ht/git.sr.ht/

[gitlab-link]: https://about.gitlab.com

[gitiles-link]: https://gerrit.googlesource.com/gitiles/

[gerrit-link]: https://www.gerritcodereview.com/

[go-source-tag]: https://github.com/golang/gddo/wiki/Source-Code-Links

[starlark-link]: https://github.com/bazelbuild/starlark

[starlark-go]: https://github.com/google/starlark-go

[bazel-link]: https://bazel.build/

[launchpad-bzr-link]: https://launchpad.net/bzr

[sourcehut-hg-link]: https://hg.sr.ht/
