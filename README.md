# Data Structures in Go

This is a reference package developed for practise purpose,and covers basic data structure implementation.
This Go package provides a collection of data structures with a focus on an array list implementation. It includes:


## Data Types
- Array List: A dynamic array that supports automatic resizing and efficient access.

## Data Structures
- Queue: A FIFO (First In, First Out) structure implemented using the array list, ideal for scenarios where elements are processed in the order they are added.
- Stack: A LIFO (Last In, First Out) structure also based on the array list, suitable for use cases where the most recently added elements are accessed first.

The package offers simple and efficient methods for managing these data structures, making it easy to incorporate them into your Go applications.

## Installation
```bash
go get github.com/vedashruta/go-ds
```

## Running unit test cases
unit_test.go in respective module covers all test cases for the implementations<br>
Run 
```bash
go test ./...
```
to see if all the test cases are covered

# Reference
https://github.com/emirpasic/gods
