package main

import "fmt"

// Slices are like arrays but the length is changable.
func main() {
	xi := []int {1, 2, 3}

	fmt.Println(xi)

// adding to a slice:
	xi = append(xi, 5, 6, 7)
	fmt.Println(xi)

// slicing a slice
	zi := []int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("zi - %#v\n", zi)
	// [inclusive:exclusive]
	fmt.Printf("zi - %#v\n", zi[0:4])

	// [:exclusive] - from beginning
	fmt.Printf("zi - %#v\n", zi[:7])

	// [inclusive:] - until the end
	fmt.Printf("zi - %#v\n", zi[4:])

	//[:] - everything
	fmt.Printf("zi - %#v\n", zi[:])

// DELETING FORM SLICE
	zi = append(zi[:4], zi[5:]...)

	fmt.Println(zi)

// SLICE - Make
	si := []string{"a", "b", "c"}
	fmt.Println(si)
	
	yi := make([]int, 0, 10)
	fmt.Println(yi) // []
	fmt.Println(len(yi)) // 0
	fmt.Println(cap(yi)) // 10 (capacity)

// MULTIDIMENSIONAL SLICE
	jb := []string {"James", "Bond", "Martini", "Chocolate"}
	jm := []string {"Jenny", "Moneypenny", "Guiness", "Wolverine Tracks"}

	fmt.Println(jb)
	fmt.Println(jm)

	xxs := [][]string{jb, jm}

	fmt.Println(xxs)

	


}