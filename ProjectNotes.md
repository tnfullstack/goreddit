## Go Test Methods 

The list of some Go Test methods used in this training project. 
The complete list of Go Test methods is available from [go.dev/testing](https://pkg.go.dev/testing#pkg-index)

| Method                             | Purpose                                                             |
| ---------------------------------- | ------------------------------------------------------------------- |
| t.Log(args ...any)                 | Logs a message to the test output                                   |
| t.Logf(format string, args ...any) | Logs a formatted log message to the test output                     |
| t.Fail()                           | Marks the test as a failed test and keeps running the test function |
| t.FailNow()                        | Marks the test as failed test and stop running the test funtion     |
|                                    |