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

title: Starlark
description: |
    Starlark is a dialect of Python designed for use as an efficient and
    powerful configuration language with safety and simplicity in mind.
---

# Starlark Configuration Language

**Starlark**, formerly known as *Skylark*, is a configuration language designed
at Google for simplicity, safety and expressivity. Starlark is a dialect of
Python, removed of some Turing-completness for safety.

Starlark is commonly used within [Bazel](https://bazel.build) and
[Drone](https://drone.io/).
The design choices behind Starlark ensures multiple neat properties to the
language and, by extension, to the configurations you write: determinism,
hermeticity, parallelism, and simplicity.

Here is an exemple of a Starlark code (courtesy of [Bazel's Starlark repository](https://github.com/bazelbuild/starlark?tab=readme-ov-file#tour)):

```python
# Define a number
number = 18

# Define a dictionary
people = {
    "Alice": 22,
    "Bob": 40,
    "Charlie": 55,
    "Dave": 14,
}

names = ", ".join(people.keys())  # Alice, Bob, Charlie, Dave

# Define a function
def greet(name):
    """Return a greeting."""
    return "Hello {}!".format(name)

greeting = greet(names)

above30 = [name for name, age in people.items() if age >= 30]
print("{} people are above 30.".format(len(above30)))

def fizz_buzz(n):
    """Print Fizz Buzz numbers from 1 to n."""
    for i in range(1, n + 1):
        s = ""
        if i % 3 == 0:
            s += "Fizz"
            if i % 5 == 0:
                s += "Buzz"
                print(s if s else i)

fizz_buzz(20)
```

## Implementations

SVGU is build upon the Go implementation of Starlark, but other exists.
We encourage the adoption of Starlark everywhere.

- [Go](https://github.com/google/starlark-go/)
- [Java (Bazel)](https://github.com/bazelbuild/bazel/tree/master/src/main/java/net/starlark/java)
- [Rust](https://github.com/facebookexperimental/starlark-rust)

See the [Starlark repository](https://github.com/bazelbuild/starlark/blob/master/users.md)
for a list of Bazel tools, users, and more.
