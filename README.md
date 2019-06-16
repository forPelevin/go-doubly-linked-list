# Go doubly linked list
The Golang module that have implementation of the doubly linked list structure.

## Getting Started
1) Download the project from github to your host machine.
2) Go to the folder with project

## Prerequisites
For the successful using you must have:
```
go >= 1.12
```
## Running the tests
It's a good practice to run the tests before using the module to make sure everything is OK.
```
go test -v --bench . --benchmem
```
You should have something like that output:
```
=== RUN   TestDoublyLinkedList
--- PASS: TestDoublyLinkedList (0.00s)
goos: darwin
goarch: amd64
pkg: github.com/forPelevin/go-doubly-linked-list
BenchmarkDoublyLinkedList-8   	 5000000	       290 ns/op	      96 B/op	       2 allocs/op
```
## Sample of using
```go
dll := DoublyLinkedList{}

// Inserts a new item at the front of list.
dll.PushFront(1)

// Inserts a new item at the back of list.
dll.PushLast(1)

// Returns the list size. In our example it equals 2.
dll.Len()

// Removes the first item.
dll.First().Remove()

// Removes the last item.
dll.Last().Remove()
```
## License
This project is licensed under the MIT License.