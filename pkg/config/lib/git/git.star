# Utilities to index Go modules hosted on Git repositories.

_GIT = "git"

# https://github.com/github/renaming
_GITHUB_DEFAULT_INSTANCE = "https://github.com"
_GITHUB_DEFAULT_MASTER = "main"

def _github(
        name,
        user,
        repo,
        branch = _GITHUB_DEFAULT_MASTER,
        instance = _GITHUB_DEFAULT_INSTANCE):
    """Register a module hosted on GitHub.

    Args:
        name (str): The name of the module.
        user (str): The name of the user or organization.
        repo (str): The name of the repository.
        branch (str): The name of the branch.
            Defaults to `git.GITHUB_DEFAULT_MASTER`.
        instance (str): The name of the instance.
            Defaults to `git.GITHUB_DEFAULT_INSTANCE`.
    """

    return module(
        name = name,
        vcs = _GIT,
        repo = "%s/%s/%s" % (instance, user, repo),
        dir = "%s/%s/%s/tree/%s{/dir}" % (instance, user, repo, branch),
        file = "%s/%s/%s/blob/%s{/dir}/{file}#L{line}" %
               (instance, user, repo, branch),
    )

_SOURCEHUT_DEFAULT_INSTANCE = "git.sr.ht"
_SOURCEHUT_DEFAULT_MASTER = "master"

# TODO(nc0): See the status for organizations and groups, as they are expected
#   to use another symbol than `~`.
def _sourcehut(
        name,
        user,
        repo,
        branch = _SOURCEHUT_DEFAULT_MASTER,
        instance = _SOURCEHUT_DEFAULT_INSTANCE):
    """Register a module hosted on Source Hut's Git hosting.

    Args:
        name (str): The name of the module.
        user (str): The name of the user or organization.
        repo (str): The name of the repository.
        branch (str): The name of the branch.
            Defaults to `git.SOURCEHUT_DEFAULT_MASTER`.
        instance (str): The name of the instance.
            Defaults to `git.SOURCEHUT_DEFAULT_INSTANCE`.
    """

    return module(
        name = name,
        vcs = _GIT,
        repo = "%s/~%s/%s" % (instance, user, repo),
        dir = "%s/~%s/%s/tree/%s{/dir}" %
              (instance, user, repo, branch),
        file = "%s/~%s/%s/tree/%s/item{/dir}/{file}#L{line}" %
               (instance, user, repo, branch),
    )

# https://about.gitlab.com/blog/2021/03/10/new-git-default-branch-name/
_GITLAB_DEFAULT_INSTANCE = "https://gitlab.com"
_GITLAB_DEFAULT_MASTER = "main"

def _gitlab(
        name,
        user,
        repo,
        branch = _GITLAB_DEFAULT_MASTER,
        instance = _GITLAB_DEFAULT_INSTANCE):
    """Register a module hosted on GitLab.

    Args:
        name (str): The name of the module.
        user (str): The name of the user or organization.
        repo (str): The name of the repository.
        branch (str): The name of the branch.
            Defaults to `git.GITLAB_DEFAULT_MASTER`.
        instance (str): The name of the instance.
            Defaults to `git.GITLAB_DEFAULT_INSTANCE`.
    """

    return module(
        name = name,
        vcs = _GIT,
        repo = "%s/%s/%s" % (instance, user, repo),
        dir = "%s/%s/%s/-/tree/%s{/dir}" % (instance, user, repo, branch),
        file = "%s/%s/%s/-/blob/%s{/dir}/{file}#L{line}" %
               (instance, user, repo, branch),
    )

# https://confluence.atlassian.com/bitbucketserver/setting-a-system-wide-default-branch-name-1021220665.html
_BITBUCKET_DEFAULT_INSTANCE = "https://bitbucket.org"
_BITBUCKET_DEFAULT_MASTER = "master"

def _bitbucket(
        name,
        workspace,
        repo,
        branch = _BITBUCKET_DEFAULT_MASTER,
        instance = _BITBUCKET_DEFAULT_INSTANCE):
    """Register a module hosted on Bitbucket.

    By default, we use https://bitbucket.org, however you usually need to
    change it with your username and workspace, following:

        https://<username>@bitbucket.org/<workspace>

    Args:
        name (str): The name of the module.
        workspace (str): The ID of the workspace.
        repo (str): The name of the repository.
        branch (str): The name of the branch.
            Defaults to `git.BITBUCKET_DEFAULT_MASTER`.
        instance (str): The name of the instance.
            Defaults to `git.BITBUCKET_DEFAULT_INSTANCE`.
    """

    return module(
        name = name,
        vcs = _GIT,
        repo = "%s/%s/%s" % (instance, workspace, repo),
        dir = "%s/%s/%s/src/%s{/dir}" % (instance, workspace, repo, branch),
        file = "%s/%s/%s/src/%s{/dir}/{file}#{file}-{line}" %
               (instance, user, repo, branch),
    )

_GITILES_DEFAULT_MASTER = "master"

def _gitiles(
        name,
        instance,
        repo,
        branch = _GITILES_DEFAULT_MASTER):
    """Register a module hosted a Gitiles (Gerrit) installation.

    Args:
        name (str): The name of the module.
        instance (str): The Gitiles instance URL.
        repo (str): The repository path.
        branch (str):  The name of the branch.
            Defaults to `git.GITILES_DEFAULT_MASTER`.
    """

    return module(
        name = name,
        vcs = _GIT,
        repo = "%s/%s" % (instance, repo),
        dir = "%s/%s/+/refs/heads/%s{/dir}" % (instance, repo, branch),
        file = "%s/%s/+/refs/heads/%s{/dir}/{file}#{line}" %
               (instance, repo, branch),
    )

git = make_module(
    "git",
    GIT = _GIT,
    GITHUB_DEFAULT_INSTANCE = _GITHUB_DEFAULT_INSTANCE,
    GITHUB_DEFAULT_MASTER = _GITHUB_DEFAULT_MASTER,
    SOURCEHUT_DEFAULT_INSTANCE = _SOURCEHUT_DEFAULT_INSTANCE,
    SOURCEHUT_DEFAULT_MASTER = _SOURCEHUT_DEFAULT_MASTER,
    GITLAB_DEFAULT_INSTANCE = _GITLAB_DEFAULT_INSTANCE,
    GITLAB_DEFAULT_MASTER = _GITLAB_DEFAULT_MASTER,
    BITBUCKET_DEFAULT_INSTANCE = _BITBUCKET_DEFAULT_INSTANCE,
    BITBUCKET_DEFAULT_MASTER = _BITBUCKET_DEFAULT_MASTER,
    GITILES_DEFAULT_MASTER = _GITILES_DEFAULT_MASTER,
    github = _github,
    sourcehut = _sourcehut,
    gitlab = _gitlab,
    bitbucket = _bitbucket,
    gitiles = _gitiles,
)
