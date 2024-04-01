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
 
