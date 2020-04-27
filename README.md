# Go Project-abc

## go project template for github.

### Features
1. Makefile
2. Dockerfile
3. Version.go
4. github ACTIONS (you may need change Github secrets settings)

### Preference
1. LOG: github.com/rs/zerolog
2. TESTING: github.com/stretchr/testify

### ps

Remember to run `go mod init` after git clone:

```bash
git clone github.com/datewu/project-abc
rm -rf go.*
go mod init github.com/YOUR_NAME/YOUR_PROJECT
go build
make test
make build

## update .gitignore
## edit README ....
## edit github secrets settings
git add .
git commit -am "init commit"
git push 

``` 