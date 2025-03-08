# Dining Philosophers

## Required Concepts

* Golang
  * programming fundamentals: variables, functions, for loops, structs
  * modules
    - how to build and run multi-file go programs
    - `go mod init example.com/dinner` : creates a go.mod file for your program
    - `go build` : builds the binary
    - `go run .` : runs the binary
  * concurrency
    - channels 
      - threadsafe data structure
      - it's like Array, List, and Queue
      - used to keep a list of messages
      - messages can be any data type: boolean, integer, or object pointer
    - wait groups
      - threadsafe data structure
      - used to keep track of all your threads
      - you increment a counter in each thread: `wg.Add(1)`
      - when your thread finishes, you decrement the counter: `defer wg.Done()`
    - goroutines (threads)
      - `go` keyword followed by a function call
      - creates a thread that is only executing that function
      - eg `go philosopher.Eat()`

* Operating Systems Concurrency
  * shared resources
  * synchronization
  * mutual exclusion
    - lock a resource, no one else can use it except me
    - unlock the resource, ok now other processes/threads can use it
  * deadlock
 
