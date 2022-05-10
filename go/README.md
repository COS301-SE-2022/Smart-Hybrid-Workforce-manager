# Testing in Go  
## File Naming
Files that are used for testing in Go should be named in the following way:  
> the name of the file or module being tested, followed by a _test.go suffix, e.g. A test for a file, funcs.go, would be named, funcs_test.go.  


## File structure
The test files should be kept in the same directory as the code that they are testing, hence the following example structure must be followed
```bash
.
├── dir1
│   ├── file1.go
│   └── file1_test.go
├── dir2
│   ├── file2.go
│   ├── file2_test.go
│   ├── file3.go
│   └── file3_test.go
```
Further, the testing file should be declared as being in the same package as it is testing.
