# Go Project-abc

[![Go Report Card](https://goreportcard.com/badge/github.com/datewu/project-abc?style=flat-square)](https://goreportcard.com/report/github.com/datewu/project-abc)

// pkg only; delete this badge on main/binary program
[![Godoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/datewu/project-abc)

## Description
A go project template on github.

### Features
1. Makefile
2. Dockerfile
3. Version.go
4. github ACTIONS (you may need to set repo secrets)

### Preference
1. log: github.com/rs/zerolog
2. test: github.com/stretchr/testify

### Usage
```bash
git clone github.com/datewu/project-abc

## *** change main.go #L3: 
## ***    import "github.com/datewu/project-abc/utils"
## *** to
## ***   import "github.com/${YOUR_REPO}/utils"
## *** AND
## *** change Readme.md accordingly.

## you may need to change Makefile #L68:
##    @go mod init github.com/datewu/${APP}
##  to
##    @go mod init github.com/${YOUR_REPO}

make custom

## configure github secrets settings

``` 
