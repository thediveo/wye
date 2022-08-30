# Why `wye`?

[![PkgGoDev](https://pkg.go.dev/badge/github.com/thediveo/wye)](https://pkg.go.dev/github.com/thediveo/wye)
[![GitHub](https://img.shields.io/github/license/thediveo/wye)](https://img.shields.io/github/license/thediveo/wye)
![build and test](https://github.com/thediveo/wye/workflows/build%20and%20test/badge.svg?branch=master)
![goroutines](https://img.shields.io/badge/go%20routines-not%20leaking-success)
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)
[![Go Report Card](https://goreportcard.com/badge/github.com/thediveo/wye)](https://goreportcard.com/report/github.com/thediveo/wye)

When you need to mix in one
[context.Context](https://pkg.go.dev/context#Context) into another Context.

Why would you ever want to do this?

Because “someone” terribly messed up an API, grossly misusing contexts on them
design spree.

Such as [Podman](https://github.com/containers/podman)'s REST API client
bindings.

([_obligatory Captain Picard meme reference_](https://knowyourmeme.com/memes/facepalm))

## Fun Fact

The module should have been named “waɪ”, but waɪ adding insult to injury by
using a Unicode import path?

## Copyright and License

Copyright 2022 Harald Albrecht, licensed under the Apache License, Version 2.0.
