## Go Test Methods 

This is a list of some Go Test methods used in this training project so far. 
The complete list of Go Test methods is available from [go.dev/testing](https://pkg.go.dev/testing#pkg-index)

| Method                               | Purpose                                                             |
| ------------------------------------ | ------------------------------------------------------------------- |
| t.Log(args ...any)                   | Logs a message to the test output                                   |
| t.Logf(format string, args ...any)   | Logs a formatted log message to the test output                     |
| t.Fail()                             | Marks the test as a failed test and keeps running the test function |
| t.FailNow()                          | Marks the test as failed test and stop running the test funtion     |
| t.Error(args ...any)                 | It is equivalent to calling Log and Fail methods                    |
| t.Errorf(format string, args ...any) | It is equivalent to calling Logf and Failf methods                  |
| t.Fatal(args ...any)                 | It is equivalent to calling Log and FailNow methods                 |
| t.Fatalf(format string, args ...any) | It is equivalent to calling Logf and FailNow methods                |
| t.Helper()                           | It tells the testing package that this is a test helper             |

### Running Subtest
Useful commands for running tests  

| Command                                  | Purpose                                                          |
| ---------------------------------------- | ---------------------------------------------------------------- |
| go test -v -run=TestParse                | Run a any subtest start with TestParse                           |
| go test -v -run=TestParse$               | Run only specific subtest                                        |
| go test -run=TestURLPort/^host_with_port | Run subtest begin with ^host_with_port                           |
| go test -run=TestURLPort/with_port       | Run any subtest name contain with_port                           |
| go test -v -run=URL/with_port            | Run any top level contains URL, ans subtest start with with_port |

### Measuring Test Coverage
Useful test coverage commands 

| Command                                        | Purpose                                       |
| ---------------------------------------------- | --------------------------------------------- |
| go test -cover                                 | Run all test and show coverage                |
| go test -coverprofile cover.out                | Run test and get test coverage output         |
| go tool cover -html=cover.out                  | Show test coverage output from cover.out file |
| go tool cover -html=cover.out -o coverage.html | Generate an html test report                  |
| go tool cover -func=cover.out                  | Print test coverage result to console         |
 
## Cross-platform compiling 

### Go build 

Useful Go build commands

| Command                                           | Purpose                                                          |
| ------------------------------------------------- | ---------------------------------------------------------------- |
| go build                                          | build executable binary in current directory                     |
| go env GOOS GOARCH                                | return the os platform and cpu architecture                      |
| go tool dist list -json                           | return a json list of supported OS and CPU                       |
| GOOS=windows GOARCH=amd64 go build                | compile a windows execuable with default directory and file name |
| GOOS=windows GOARCH=amd64 go build -o bin/hit.ext | -o direct executable file to bin directory and name it hit.exe   |

### Using a make file 
Using a make file to automate compile tool for multiple operating systems. 

The content of the make file below allows the go tool to compile binary for Linux, MacOS intel, MacOS M1, M2, M3, and Windows...  

```make
compile:
			# compile it for linux
			GOOS=linux GOARCH=amd64 go build -o ./bin/hit_linux_amd64 ./cmd/hit
			# compile it for macOS
			GOOS=darwin GOARCH=amd64 go build -o ./bin/hit_darwin_amd64 ./cmd/hit
			# compile it for Apple M1
			GOOS=darwin GOARCH=arm64 go build -o ./bin/hit_darwin_arm64 ./cmd/hit
			# compile it for Windows
			GOOS=windows GOARCH=amd64 go build -o ./bin/hit_win_amd64.exe ./cmd/hit
```
