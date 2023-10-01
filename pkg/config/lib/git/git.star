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

# Utilities to index Go modules hosted on Git repositories.

_GIT = "git"

# https://github.com/github/renaming
_GITHUB_DEFAULT_INSTANCE = "https://github.com"
_GITHUB_DEFAULT_REF = "main"

def _github(
        name,
        user,
        repo,
        ref = _GITHUB_DEFAULT_REF,
        instance = _GITHUB_DEFAULT_INSTANCE):
    """Register a module hosted on GitHub.

    Args:
        name (str): The name of the module.
        user (str): The name of the user or organization.
        repo (str): The name of the repository.
        ref (str): The name of the ref.
            Defaults to `git.GITHUB_DEFAULT_REF`.
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

_SOURCEHUT_DEFAULT_INSTANCE = "https://git.sr.ht"
_SOURCEHUT_DEFAULT_REF = "master"

# TODO(nc0): See the status for organizations and groups, as they are expected
#   to use another symbol than `~`.
def _sourcehut(
        name,
        user,
        repo,
        ref = _SOURCEHUT_DEFAULT_REF,
        instance = _SOURCEHUT_DEFAULT_INSTANCE):
    """Register a module hosted on Source Hut's Git hosting.

    Args:
        name (str): The name of the module.
        user (str): The name of the user or organization.
        repo (str): The name of the repository.
        ref (str): The name of the ref.
            Defaults to `git.SOURCEHUT_DEFAULT_REF`.
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
_GITLAB_DEFAULT_REF = "main"

def _gitlab(
        name,
        user,
        repo,
        ref = _GITLAB_DEFAULT_REF,
        instance = _GITLAB_DEFAULT_INSTANCE):
    """Register a module hosted on GitLab.

    Args:
        name (str): The name of the module.
        user (str): The name of the user or organization.
        repo (str): The name of the repository.
        ref (str): The name of the ref.
            Defaults to `git.GITLAB_DEFAULT_REF`.
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
               (instance, workspace, repo, ref),
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
