# Go Project-abc

[![Go Report Card](https://goreportcard.com/badge/github.com/datewu/project-abc?style=flat-square)](https://goreportcard.com/report/github.com/datewu/project-abc)
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

### ps

Remember to run `go mod init` after git clone:

```bash
git clone github.com/datewu/project-abc
## change Makefile #L59: @go mod init github.com/YOUR_REPO/${APP}
make custom

## configure github secrets settings
``` 
