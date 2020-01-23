#### serverscoop

##### Windows
Compile for Windows
$ GOOS=windows GOARCH=386 go build -o serverscoop.exe .

From the docs [https://github.com/golang/go/wiki/WindowsCrossCompiling],
targeting a different GOOS and GOARCH "will silently rebuild most of the
standard library, and for this reason will be quite slow. To speed up the
process, you can install all the windows-amd64 standard packages on your
system with ```$ GOOS=windows GOARCH=amd64 go install```"
