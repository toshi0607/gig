gig
====

## Description
generate (or output) .gitignore using 

## Demo
![](https://user-images.githubusercontent.com/7035446/39394981-84cef0aa-4b13-11e8-86b9-7af4f979efa3.gif)

## VS. 

### gibo
[simonwhitaker/gio](https://github.com/simonwhitaker/gibo) is useful tool for .gitinore.
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

## Install

```sh
$ go get -u github.com/toshi0607/gig
```

## Contribution

Coming soon!

## Licence
[MIT](LICENSE) file for details.

## Author

[toshi0607](https://github.com/toshi0607)