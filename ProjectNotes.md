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


