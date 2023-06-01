# Utilities to index Go modules hosted on Bazaar repositories.

_BAZAAR = "bzr"

_LAUNCHPAD_DEFAULT_INSTANCE = "https://bazaar.launchpad.net"
_LAUNCHPAD_DEFAULT_REV = "head:"
_LAUNCHPAD_DEFAULT_BRANCH = "trunk"

def _launchpad(
        name,
        user,
        repo,
        branch = _LAUNCHPAD_DEFAULT_BRANCH,
        rev = _LAUNCHPAD_DEFAULT_REV,
        instance = _LAUNCHPAD_DEFAULT_INSTANCE):
    """Register a module hosted on Launchpad.

    Args:
        name (str): The name of the module.
        user (str): The ID of the user.
        repo (str): The name of the repository.
        branch (str): The name of the branch.
            Defaults to `bzr.LAUNCHPAD_DEFAULT_BRANCH`.
        rev (str): The revision number.
            Defaults to `bzr.LAUNCHPAD_DEFAULT_REV`.
        instance (str): The name of the instance.
            Defaults to `bzr.LAUNCHPAD_DEFAULT_INSTANCE`.
    """

    return module(
        name = name,
        vcs = _BAZAAR,
        repo = "%s/~%s/%s/%s" % (instance, user, repo, branch),
        dir = "%s/~%s/%s/%s/files/%s{/dir}" %
              (instance, user, repo, branch, rev),
        file = "%s/~%s/%s/%s/view/%s{/dir}/{file}#L{line}" %
               (instance, user, repo, branch, rev),
    )

bzr = make_module(
    "bzr",
    BAZAR = _BAZAAR,
    LAUNCHPAD_DEFAULT_INSTANCE = _LAUNCHPAD_DEFAULT_INSTANCE,
    LAUNCHPAD_DEFAULT_REV = _LAUNCHPAD_DEFAULT_REV,
    LAUNCHPAD_DEFAULT_BRANCH = _LAUNCHPAD_DEFAULT_BRANCH,
    launchpad = _launchpad,
)
