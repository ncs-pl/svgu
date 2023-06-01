# Utilities to index Go modules hosted on Mercurial repositories.

_MERCURIAL = "hg"

_SOURCEHUT_DEFAULT_INSTANCE = "https://hg.sr.ht"
_SOURCEHUT_DEFAULT_REV = "tip"

def _sourcehut(
        name,
        user,
        repo,
        rev = _SOURCEHUT_DEFAULT_REV,
        instance = _SOURCEHUT_DEFAULT_INSTANCE):
    """Register a module hosted on a Mercurial repository on Source Hut.

    Args:
        name (str): The name of the module.
        user (str): The name of the user or organization that owns the
            repository.
        repo (str): The name of the repository.
        rev (str): The revision to use. Defaults to `hg.SOURCEHUT_DEFAULT_REV`.
        instance (str): The instance of Source Hut to use.
            Defaults to `hg.SOURCEHUT_DEFAULT_INSTANCE`.
    """

    module(
        name = name,
        vcs = _MERCURIAL,
        repo = "%s/~%s/%s" % (instance, user, repo),
        dir = "%s/~%s/%s/browse{/dir}?rev=%s" %
              (instance, user, repo, rev),
        file = "%s/~%s/%s/browse{/dir}/{file}?rev=%s#L{line}" %
               (instance, user, repo, rev),
    )

hg = make_module(
    "hg",
    MERCURIAL = _MERCURIAL,
    SOURCEHUT_DEFAULT_INSTANCE = _SOURCEHUT_DEFAULT_INSTANCE,
    SOURCEHUT_DEFAULT_REV = _SOURCEHUT_DEFAULT_REV,
    sourcehut = _sourcehut,
)
