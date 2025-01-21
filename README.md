# Go Set Implementation

This repository provides a simple and efficient implementation of a set data structure in Go. It uses a map under the hood for fast lookups and insertions.

## Installation
```bash
go get github.com/speculum-factorem/go-set/set
```
## Usage
```go
package main

import (
 "fmt"
 "github.com/speculum-factorem/go-set/set"
)

func main() {
 s := set.New()

 s.Add(1)
 s.Add("hello")
 s.Add(3.14)

 fmt.Println(s.Contains(1))    // Output: true
 fmt.Println(s.Contains(2))    // Output: false
 fmt.Println(s.Size())       // Output: 3

 s.Remove("hello")
 fmt.Println(s.Size())       // Output: 2

 elements := s.Elements()
 fmt.Println(elements)      // Output: [1 3.14] (order not guaranteed)

 s.Clear()
 fmt.Println(s.Size())       // Output: 0
}
```

## API Reference

* *`New()`*: Creates a new empty set. Returns a pointer to the `Set` struct.

* *`Add(element interface{})`*: Adds an element to the set. If the element already exists, it has no effect.

* *`Remove(element interface{})`*: Removes an element from the set. If the element doesn't exist, it has no effect.

* *`Contains(element interface{}) bool`*: Checks if an element is present in the set. Returns `true` if the element is found, `false` otherwise.

* *`Size() int`*: Returns the number of elements in the set.

* *`Clear()`*: Removes all elements from the set.

* *`Elements() []interface{}`*: Returns a slice containing all elements in the set. The order of elements is not guaranteed.


## Notes

* The set can store elements of any type using `interface{}`.

* The implementation uses a map with empty structs as values to minimize memory overhead.

* Since `Elements()` returns a slice of `interface{}`, you will likely need to use type assertions to work with the elements.


## Contributing

Contributions are welcome! Please feel free to submit pull requests for bug fixes, new features, or improvements to the documentation.