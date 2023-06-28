# SVGU

SVGU (short for *Static, Vanity, Go URL*) is a flexible and easy to use
tool for creating vanity URLs for your Go projects.
It allows publishing a Go module under a custom domain name, without the need
to use your code forge's domain name.

For example, if you have a project hosted on GitHub, you can use SVGU to
publish it under a custom domain name, such as `myproject.com/foo` instead of
`github.com/myproject/foo`.

## How it works

SVGU requires a configuration file, usually named `DOMAINS.star`, which
describes the modules to export.

The configuration file is a [Starlark](https://starlark.net) script, which
allows for a lot of flexibility.
Starlark is a subset of Python for configuration files, and is used by the
[Bazel](https://bazel.build) build system and others.
It is a simple language, and you don't need to know Python to use it.

See the [reference documentation](doc/references.md) for more information.

Once the configuration file is ready, you can run SVGU to generate the
necessary files, and then publish them on your web server.

```shell
$ svgu
```

This will generate a `dst` directory containing the files to publish.

## Documentation

- [Getting started](doc/getting-started.md)
- [Reference documentation](doc/references.md)
- [Bug tracker](https://todo.sr.ht/~n1c00o/svgu)
- [Mailing list](https://lists.sr.ht/~n1c00o/svgu)
- [Source code](https://git.sr.ht/~n1c00o/svgu)
- [Project page](https://sr.ht/~n1c00o/svgu)

## License

The SVGU project is governed by a [BSD-style license](LICENSE).
The documentation is licensed under the [Creative Commons Attribution 4.0
International License](https://creativecommons.org/licenses/by/4.0/).
