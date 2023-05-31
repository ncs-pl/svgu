# Utilities to index Go modules hosted on Git repositories.

_GIT = "git"

# https://github.com/github/renaming
_GITHUB_DEFAULT_INSTANCE = "https://github.com"
_GITHUB_DEFAULT_MASTER = "main"

def _github(
        name,
        user,
        repo,
        ref = _GITHUB_DEFAULT_MASTER,
        instance = _GITHUB_DEFAULT_INSTANCE):
    """Register a module hosted on GitHub.

    Args:
        name (str): The name of the module.
        user (str): The name of the user or organization.
        repo (str): The name of the repository.
        ref (str): The name of the ref.
            Defaults to `git.GITHUB_DEFAULT_MASTER`.
        instance (str): The name of the instance.
            Defaults to `git.GITHUB_DEFAULT_INSTANCE`.
    """

    return module(
        name = name,
        vcs = _GIT,
        repo = "%s/%s/%s" % (instance, user, repo),
        dir = "%s/%s/%s/tree/%s{/dir}" % (instance, user, repo, ref),
        file = "%s/%s/%s/blob/%s{/dir}/{file}#L{line}" %
               (instance, user, repo, ref),
    )

_SOURCEHUT_DEFAULT_INSTANCE = "git.sr.ht"
_SOURCEHUT_DEFAULT_MASTER = "master"

# TODO(nc0): See the status for organizations and groups, as they are expected
#   to use another symbol than `~`.
def _sourcehut(
        name,
        user,
        repo,
        ref = _SOURCEHUT_DEFAULT_MASTER,
        instance = _SOURCEHUT_DEFAULT_INSTANCE):
    """Register a module hosted on Source Hut's Git hosting.

    Args:
        name (str): The name of the module.
        user (str): The name of the user or organization.
        repo (str): The name of the repository.
        ref (str): The name of the ref.
            Defaults to `git.SOURCEHUT_DEFAULT_MASTER`.
        instance (str): The name of the instance.
            Defaults to `git.SOURCEHUT_DEFAULT_INSTANCE`.
    """

    return module(
        name = name,
        vcs = _GIT,
        repo = "%s/~%s/%s" % (instance, user, repo),
        dir = "%s/~%s/%s/tree/%s{/dir}" %
              (instance, user, repo, ref),
        file = "%s/~%s/%s/tree/%s/item{/dir}/{file}#L{line}" %
               (instance, user, repo, ref),
    )

# https://about.gitlab.com/blog/2021/03/10/new-git-default-branch-name/
_GITLAB_DEFAULT_INSTANCE = "https://gitlab.com"
_GITLAB_DEFAULT_MASTER = "main"

def _gitlab(
        name,
        user,
        repo,
        ref = _GITLAB_DEFAULT_MASTER,
        instance = _GITLAB_DEFAULT_INSTANCE):
    """Register a module hosted on GitLab.

    Args:
        name (str): The name of the module.
        user (str): The name of the user or organization.
        repo (str): The name of the repository.
        ref (str): The name of the ref.
            Defaults to `git.GITLAB_DEFAULT_MASTER`.
        instance (str): The name of the instance.
            Defaults to `git.GITLAB_DEFAULT_INSTANCE`.
    """

    return module(
        name = name,
        vcs = _GIT,
        repo = "%s/%s/%s" % (instance, user, repo),
        dir = "%s/%s/%s/-/tree/%s{/dir}" % (instance, user, repo, ref),
        file = "%s/%s/%s/-/blob/%s{/dir}/{file}#L{line}" %
               (instance, user, repo, ref),
    )

# https://confluence.atlassian.com/bitbucketserver/setting-a-system-wide-default-branch-name-1021220665.html
_BITBUCKET_DEFAULT_INSTANCE = "https://bitbucket.org"
_BITBUCKET_DEFAULT_REF = "master"

def _bitbucket(
        name,
        workspace,
        repo,
        ref = _BITBUCKET_DEFAULT_REF,
        instance = _BITBUCKET_DEFAULT_INSTANCE):
    """Register a module hosted on Bitbucket.

    By default, we use https://bitbucket.org, however you usually need to
    change it with your username and workspace, following:

        https://<username>@bitbucket.org/<workspace>

    Args:
        name (str): The name of the module.
        workspace (str): The ID of the workspace.
        repo (str): The name of the repository.
        ref (str): The name of the ref.
            Defaults to `git.BITBUCKET_DEFAULT_REF`.
        instance (str): The name of the instance.
            Defaults to `git.BITBUCKET_DEFAULT_INSTANCE`.
    """

    return module(
        name = name,
        vcs = _GIT,
        repo = "%s/%s/%s" % (instance, workspace, repo),
        dir = "%s/%s/%s/src/%s{/dir}" % (instance, workspace, repo, ref),
        file = "%s/%s/%s/src/%s{/dir}/{file}#{file}-{line}" %
               (instance, user, repo, ref),
    )

_GITILES_DEFAULT_REF = "master"

def _gitiles(
        name,
        instance,
        repo,
        ref = _GITILES_DEFAULT_REF):
    """Register a module hosted a Gitiles (Gerrit) installation.

    Args:
        name (str): The name of the module.
        instance (str): The Gitiles instance URL.
        repo (str): The repository path.
        ref (str):  The name of the ref.
            Defaults to `git.GITILES_DEFAULT_REF`.
    """

    return module(
        name = name,
        vcs = _GIT,
        repo = "%s/%s" % (instance, repo),
        dir = "%s/%s/+/refs/heads/%s{/dir}" % (instance, repo, ref),
        file = "%s/%s/+/refs/heads/%s{/dir}/{file}#{line}" %
               (instance, repo, ref),
    )

git = make_module(
    "git",
    GIT = _GIT,
    GITHUB_DEFAULT_INSTANCE = _GITHUB_DEFAULT_INSTANCE,
    GITHUB_DEFAULT_REF = _GITHUB_DEFAULT_REF,
    SOURCEHUT_DEFAULT_INSTANCE = _SOURCEHUT_DEFAULT_INSTANCE,
    SOURCEHUT_DEFAULT_REF = _SOURCEHUT_DEFAULT_REF,
    GITLAB_DEFAULT_INSTANCE = _GITLAB_DEFAULT_INSTANCE,
    GITLAB_DEFAULT_REF = _GITLAB_DEFAULT_REF,
    BITBUCKET_DEFAULT_INSTANCE = _BITBUCKET_DEFAULT_INSTANCE,
    BITBUCKET_DEFAULT_REF = _BITBUCKET_DEFAULT_REF,
    GITILES_DEFAULT_REF = _GITILES_DEFAULT_REF,
    github = _github,
    sourcehut = _sourcehut,
    gitlab = _gitlab,
    bitbucket = _bitbucket,
    gitiles = _gitiles,
)
