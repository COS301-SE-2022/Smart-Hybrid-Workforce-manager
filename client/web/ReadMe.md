# Example

The example app was generated using `npm install -g create-react-app@3.4.1` followed by `npm init react-app example --use-npm`

More information can be found at: https://mherman.org/blog/dockerizing-a-react-app/

TODO [KP]: The example app should have hot reloading but it doesn't seem to work please check `./docker-compose.yml` for all the attempted fixes also see: https://github.com/facebook/create-react-app/issues/9904 to track this issue.

# React Testing
## File Naming
Files that are used for testing in Go should be named in the following way:  
> the name of the file or module being tested, followed by a .test.js suffix, e.g. A test for a file, component.js, would be named, component.test.js  


## File structure
The test files should be kept in the same directory as the code that they are testing, hence the following example structure must be followed
```bash
.
├── dir1
│   ├── file1.js
│   └── file1.test.js
├── dir2
│   ├── file2.js
│   ├── file2.test.js
│   ├── file3.js
│   └── file3.test.js
```
React comes with Jest preinstalled, however, more will be added later on the consideration of different testing libraries for improved testing. 

