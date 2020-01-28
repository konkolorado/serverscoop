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

- Functions and methods can return multiple values. A common idiom is to
  return a result as well as the error which the calling function can check
  if is nil or not.

- Functions and methods can name their return values. This lets the values
  be used as regular variables (like incoming parameters) -- at function start,
  they are initialized to zero values for their type. If the function or
  method returns with no args, the current values of the return params are
  used as the returned values.

- Go has a defer keyword which schedules a function call to be run immediately
  (right before) before the calling function returns. Arguments to the deferred
  function are evaluated when the defer executes, not when the call executes.

- new() is a way of allocating memory. new(T) zeroes out the new storage
  according to what the zero value of T's type is. It returns a pointer to the
  space.

- make() can only create slices, maps and channels and returns an uninitialized
  value of type T.

- Passing an array to a function will cause the function to receive a copy of
  the array, not a pointer to it. Also, the size of the array is part of its
  type. Passing pointers to an array to a function, however, isn't idiomatic
  and this functionality should instead be provided by using Slices

- Slices wrap arrays and are more convenient when dealing with sequences of
  data. They hold a reference to the underlying array and if you assign one
  slice to another, they both refer to the same array.  Changes that one slice
  makes to the array will be visible to other slices. The length of a slice
  may be changed as long as it's still in the limits of the underlying array

- Maps are dicts. Maps hold references to the underlying data structure so if
  a function modifies that contents of a map, changes will be available outside
  the function. Lookups for a key that doesn't exist will return the zero value
  of the map's key type. To distinguish between zero values and missing keys,
  check the error return via the "comma ok" idiom.

- Variables can be initialized like constants but the initializer can be an
  expression evaluated at run time i.e. home = os.Getenv("HOME")

- Each source file can have its own [multiple] init() to set up required state
  or verify correctness of program state before real execution begins. init()
  is called after the variable declarations in a package and those are evaluated
  after the imported packages have been initialized.

- It can be useful to import a package only for its side effects. To do this,
  rename the package to the blank identifier i.e. import _ "net/http/pprof"

- A goroutine has a simple model: it is a function executing concurrently with
  other goroutines in the same address space. Goroutines are multiplexed onto
  multiple OS threads so if one should block, such as while waiting for I/O,
  others continue to run.

- There is a built-in panic() function panic that in effect creates a run-time
  error that will stop the program

- You can use recover() to regain control of a goroutine that is unwinding due
  to a panic

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
