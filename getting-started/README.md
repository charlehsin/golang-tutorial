# Tutorial and sample codes for Getting Started guide

## Basic

- https://tour.golang.org/welcome/1
- https://golang.org/doc/tutorial/getting-started
- This is at hello folder.
- Commands at hello folder
   - (init dependency tracking) go mod init example/hello
   - (add new module and sums, and synchronize dependencies) go mod tidy
   - (run) go run .

## Creating a module

- https://golang.org/doc/tutorial/create-module
- This is at greetings folder.
- Commands at greetings folder
   - (init dependency tracking) go mod init example.com/greetings
- Commands at hello folder
   - (run) go run .

## Using a module

- https://golang.org/doc/tutorial/call-module-code
- This is at hello folder.
- Commands at hello folder
   - (dependency tracking: use the local module) go mod edit -replace example.com/greetings=../greetings
   - (add new module and sums, and synchronize dependencies) go mod tidy
   - (run) go run .

## Raising and handling errors

- https://golang.org/doc/tutorial/handle-errors
- This is at greetings and hello folders.
- Commands at hello folder
   - (run) go run .

## Random number and slice

- https://golang.org/doc/tutorial/random-greeting
- This is at greetings folder.
- Commands at hello folder
   - (run) go run .

## Function with slice input and map output

- https://golang.org/doc/tutorial/greetings-multiple-people
- This is at greetings and hello folders.
- Commands at hello folder
   - (run) go run .

## Unit testing

- https://golang.org/doc/tutorial/add-a-test
- This is at greetings folder.
- Commands at greetings folder
   - (run unit testing) go test
   - (run unit testing with verbose flag) go test -v

## Compile and install the application

- https://golang.org/doc/tutorial/compile-install
- Commands at hello folder
   - (build an exe) go build
   - (discover the go install path) go list -f '{{.Target}}'
   - (build and put the exe in the install path) go install
