load("@svgu/git.star", "git")
load("@svgu/svn.star", "svn")
load("@svgu/hg.star", "hg")
load("@svgu/bzr.star", "bzr")
load("@svgu/fossil.star", "fossil")

index(domain = "example.com")

module(
    name = "foo",
    vcs = git.GIT,
    repo = "https://example.com",
    dir = "https://example.com{/dir}",
    file = "https://example.com{/dir}/{file}",
)

git.github(
    name = "github/bar",
    user = "example",
    repo = "bar.git",
)

git.gitlab(
    name = "gitlab/bar",
    user = "example",
    repo = "bar.git",
)

git.sourcehut(
    name = "sourcehut/bar",
    user = "example",
    repo = "bar.git",
    ref = "trunk",
)

git.bitbucket(
    name = "bitbucket/bar",
    workspace = "example",
    repo = "bar.git",
    ref = "default",
    instance = "https://root@bitbucket.org",
)

git.gitiles(
    name = "gitiles/bar",
    repo = "example/bar.git",
    ref = "master",
    instance = "https://gerrit.googlesource.com",
)
