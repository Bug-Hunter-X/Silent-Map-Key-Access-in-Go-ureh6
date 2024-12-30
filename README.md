# Silent Map Key Access in Go

This repository demonstrates a common, yet subtle, issue in Go: accessing non-existent keys in maps.  Go will return the zero value for the map's value type instead of raising an error.  This can lead to hard-to-debug issues where code assumes the absence of a key is explicitly signaled.

The `bug.go` file showcases the problem, while `bugSolution.go` offers a safe way to handle such access.

**Bug:**
When you try to access a key that does not exist in a Go map, the program will not crash; instead, it will return the zero value for the type of the map's value.  This might lead to unexpected behavior if you rely on the zero value to indicate the key's absence.

**Solution:**
The best practice is to explicitly check if a key exists in a map using the `map[key]` syntax within an `if` statement or the `ok` parameter of the `_, ok := m[key]` expression.