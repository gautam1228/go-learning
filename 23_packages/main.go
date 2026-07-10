package main

import "github.com/gautam1228/go-learning/auth"

func main() {
	auth.LoginWithCredentials("codedbygautam", "gautam123")
}

// command to initialize a package: go mod init <package_name>
// Note: it's a convention to use the git repo url as the package name

// command to get a 3rd party package: go get <package_name>
// where to get a package (in-general): pkg.go.dev

// When not in use, inside the go.mod file, the package
// will stay as an indirect dependency
// Once that package is used, it becomes a direct dependency
// changes into require syntax, if it doesn't then run: go mod tidy
// It will change it to a direct dependency

// Alternatively we can also paste the link inside the import in a file
// and then run : go mod tidy
// it automatically adds the package imports to go.mod

// The go.sum file is like package-lock.json in NodeJS
// It stores the specific versions that are required
