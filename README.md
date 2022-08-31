# Why `wye`?

[![PkgGoDev](https://pkg.go.dev/badge/github.com/thediveo/wye)](https://pkg.go.dev/github.com/thediveo/wye)
[![GitHub](https://img.shields.io/github/license/thediveo/wye)](https://img.shields.io/github/license/thediveo/wye)
![build and test](https://github.com/thediveo/wye/workflows/build%20and%20test/badge.svg?branch=master)
![goroutines](https://img.shields.io/badge/go%20routines-not%20leaking-success)
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)
[![Go Report Card](https://goreportcard.com/badge/github.com/thediveo/wye)](https://goreportcard.com/report/github.com/thediveo/wye)

When you need to mix in one (shorter-lived service)
[context.Context](https://pkg.go.dev/context#Context) into another long-living
context: kind of a “Y” joint. The opposite of Go's “⅄” pattern of deriving new
contexts from existing contexts.

Why would you ever want to do this?

Because “_someone_” terribly messed up an API, grossly misusing contexts on them
REST API client design spree.

Such as [Podman](https://github.com/containers/podman)'s REST API client
bindings.

([_obligatory Captain Picard meme reference_](https://knowyourmeme.com/memes/facepalm))

## Installation

```bash
go get github.com/thediveo/wye
```

## Supported Go Versions

`wye` supports versions of Go that are noted by the [Go release
policy](https://golang.org/doc/devel/release.html#policy), that is, _N_ and
_N_-1 major versions.

## Miscellaneous

- to view the package documentation _locally_:
  - either: `make pkgsite`,
  - or, in VSCode (using the VSCode-integrated simple browser): “Tasks: Run
    Task” ⇢ “View Go module documentation”.
- `make` shows the available make targets.

## Fun Fact

The module should have been named “waɪ”, but waɪ adding insult to injury by
using a Unicode import path?

## Copyright and License

Copyright 2022 Harald Albrecht, licensed under the Apache License, Version 2.0.
