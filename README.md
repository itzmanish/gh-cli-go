# gh-cli-go

A github cli app for retrieving and downloading user info to local filesystem.
It also have a superb game you surely want to play with.
Check that out with `gh-cli-go play` or do `gh-cli-go help` for more information.

[![Go](https://github.com/itzmanish/gh-cli-go/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/itzmanish/gh-cli-go/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/itzmanish/gh-cli-go/branch/master/graph/badge.svg?token=RDWK10DQYI)](https://codecov.io/gh/itzmanish/gh-cli-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/itzmanish/gh-cli-go)](https://goreportcard.com/report/github.com/itzmanish/gh-cli-go)

Checklist

- [x] user info
- [x] repo info
- [x] organisation info
- [x] followers
- [x] following
- [x] stars
- [x] Guess the star game
- [x] Testcases
- [x] Makefile

#### Installation Instruction

```go
go build main.go -o gh-cli-go
go install
```

#### Usages

```
gh-cli-go -h
NAME:
   gh-cli-go - A Github CLI which provides downloading users data from github.

USAGE:
   gh-cli-go [global options] command [command options] [args and such]

VERSION:
   v1.0.0

AUTHOR:
   Manish <itzmanish108@gmail.com>

COMMANDS:
   init      Initialize user with github username and token
   download  Download all information available for initialized user
   show      show all information available for initialized user
   play      play guess the star game
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  Shows current cli version (default: false)

COPYRIGHT:
   (c) 2021 Manish
```
