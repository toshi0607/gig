gig
====

[![Go Report Card](https://goreportcard.com/badge/github.com/toshi0607/gig)](https://goreportcard.com/report/github.com/toshi0607/gig)
[![MIT License](http://img.shields.io/badge/license-MIT-lightgrey.svg)](https://github.com/toshi0607/gig/blob/master/LICENSE)
[![Codecov](https://codecov.io/github/toshi0607/gig/coverage.svg?branch=master)](https://codecov.io/github/toshi0607/gig?branch=master)

## Description
generate (or output) .gitignore using [github/gitignore](https://github.com/github/gitignore)

## Demo
![](https://user-images.githubusercontent.com/7035446/39394981-84cef0aa-4b13-11e8-86b9-7af4f979efa3.gif)

## VS. 

### gibo
[simonwhitaker/gibo](https://github.com/simonwhitaker/gibo) is useful tool for .gitinore.
It does `git clone` the templates from [github/gitignore](https://github.com/github/gitignore) and it uses local files.
So the gibo is fast, but you have to update local files to use tha latest template.

### gig
[toshi0607/gig](https://github.com/toshi0607/gig) is also a tool for .gitinore.
The gig always use the latest template by accessing github each time.
So you don't have to update something manually.

## Requirement
if you build gig loccally, please exec this command first.

```sh
$ dep ensure
```

## Usage

```
Usage:
  gig [OPTIONS] [Language]

Application Options:
  -l, --list      Show list of available language
  -f, --File      Output .gitignore file
  -q, --quiet     Hide stdout
  -v, --version   Show version

Help Options:
  -h, --help      Show this help message
```

### Example

```sh
# show available languages
$ gig -l
Actionscript
Ada
Agda
Android
...


# search available languages like go
$ gig -l | grep -i go
Go
Godot
IGORPro


# output to the .gitignore file
$ gig Ruby -f
$ cat .gitignore
*.gem
*.rbc
/.config
/coverage/
...


# add to the existing .gitignore file
$ gig Go >> .gitignore
$ cat .gitignore
...
# Binaries for programs and plugins
*.exe
*.exe~
...

```

### Tips

[peco](https://github.com/peco/peco) 's incremental search helps gig a lot.

```sh
$ gig $(gig -l | peco)
```

![](https://user-images.githubusercontent.com/7035446/39398424-86087f74-4b48-11e8-9428-6f771ac8074b.gif)

Setting alias like blow to your dotfile (.bashrc, .zshrc, etc) is also useful.

```sh
alias pgig='gig $(gig -l | peco)'
```

## Install

### for Homebrew (macOS, linux)

```sh
$ brew tap toshi0607/homebrew-gig
$ brew install gig
```

### for Go environment

```sh
$ go get -u github.com/toshi0607/gig
```

### for others
You can download the binary directly from [latest release](https://github.com/toshi0607/gig/releases/latest)

* gig_darwin_386.zip
* gig_darwin_amd64.zip
* gig_linux_386.zip
* gig_linux_amd64.zip
* gig_windows_386.zip
* gig_windows_amd64.zip

## Contribution

1. Fork ([https://github.com/toshi0607/gig/fork](https://github.com/toshi0607/gig/fork))
1. Create a feature branch
1. Commit your changes
1. Run test suite with the `make test` command and confirm that it passes
1. Run `gofmt -s`
1. Create new Pull Request

## Licence
[MIT](LICENSE) file for details.

## Author

* [GitHub](https://github.com/toshi0607)
* [twitter](https://twitter.com/toshi0607)
