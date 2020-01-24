#### serverscoop

##### General
- A package is a directory of related go files that all declare the
  same "package xxx" top level directive. Functions, types, vars,
  and constants declared in one source file are visible to all
  other source files within the same package

- A module is a collection of related Go packages. A Go repo
  usually contains only one module located at the root of the
  repo.

- mod.go declares the module path which is the import path prefix
  for all packages within the module

- The import path is a module's path joined with the packages's
  subdirectory within the module. Packages in the std lib do not
  have a module path prefix

- If a function inside a package begins with an upper-case letter
  it is exported and can be used in other packages that import
  the package

- Use gofmt [file, directory] to format your code. Use the -w flag
  to write the changes that gofmt would make directly back into the
  go source files. Use the -s flag to also attempt to simplify code

- The variable naming convention in go is MixedCaps or mixedCaps


- Check your current go environment
  ```go env [ENV_VAR]```

GOPATH
  - specifies directories outside of GOROOT that contain
  the source code for Go projects and their binaries
  - can be a list of directories, but generally you'll
  want to use a single directory for all your go code
  - for convenience, add the $GOPATH/bin directory to
  your path to have easy access to your binaries
  - when you use import, go will look inside the GOPATH's
  src/ directory
  - ex) GOPATH="/Users/utm1/Desktop/projects/go"
  - ex) ls $GOPATH will usually contain pkg, bin, src directories

GOROOT
  - defines where the Go compiler/tools are installed
  - it's possible to maintain distinct Go distributions in different
  locations, you just need to update your GOROOT each time
  - usually get's set automatically
  - ex) GOROOT="/usr/local/Cellar/go/1.13.6/libexec"

- Build and install a go program
  ```go install```

  - this command should be run inside the directory that contains your
  go.mod file
  - builds the code in the directory, producing an executable binary
  that then gets installed into $GOPATH/bin
  - you can control where the compiled binary gets installed by
  setting the GOBIN environment variable
  - in the scenario where the CWD is
  $GOPATH/src/github.com/konkolorado/serverscoop, ```go install```,
  ```go install .``` and
  ```go install github.com/konkolorado/serverscoop``` all do the same
  thing


##### Windows
Compile for Windows
$ GOOS=windows GOARCH=386 go build -o serverscoop.exe .

From the docs [https://github.com/golang/go/wiki/WindowsCrossCompiling],
targeting a different GOOS and GOARCH "will silently rebuild most of the
standard library, and for this reason will be quite slow. To speed up the
process, you can install all the windows-amd64 standard packages on your
system with ```$ GOOS=windows GOARCH=amd64 go install```"
