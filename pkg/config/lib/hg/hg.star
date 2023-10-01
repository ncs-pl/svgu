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
