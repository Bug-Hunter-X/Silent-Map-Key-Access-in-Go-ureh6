```go
package main

import (


    "fmt"
)

func main() {
    var m map[string]int
    m = map[string]int{"a": 1}
    value, ok := m["a"]
    if ok {
        fmt.Println("Value of a:", value)
    } else {
        fmt.Println("Key 'a' not found")
    }
    value, ok = m["b"]
    if ok {
        fmt.Println("Value of b:", value)
    } else {
        fmt.Println("Key 'b' not found")
    }
}
```