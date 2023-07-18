package main

import "fmt"

func main() {
	xi := [3]int {1, 2, 3}
	fmt.Println(xi)
}

// Arrays are rarely used because they have a set length that in unchangable. 
// This is show with the [3] in the example above. Alternatively, one could use the spread operator for Go to figure out
// the length, but it will still remain unchangable.