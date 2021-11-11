# Tutorial and sample codes for basic getting started and using other modules

- https://golang.org/doc/tutorial/getting-started
- Commands
   - (dependency tracking) go mod init example/hello
   - (dependency tracking: use the local module) go mod edit -replace example.com/greetings=../greetings
   - (add new module and sums, and synchronize dependencies) go mod tidy
   - (run) go run .
   - (get help) go help