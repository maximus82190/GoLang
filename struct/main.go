package main

import "fmt"

	type person struct {
		first string
		last string
		age int
	}

	type secretAgent struct {
		person
		ltk bool
	}

func main() {
	sa1 := secretAgent {
		person: person {
			first: "James",
			last: "Bond",
			age: 32,
		},
		ltk: true,
	}

	p2 := person {
		first: "Jenny",
		last: "Moneypenny",
		age: 27,
	}

// anonymous struct
	p3 := struct {
		first string
		last string
		age int
	}{
		first: "James",
		last: "Bond",
		age: 32,
	}
	fmt.Println(sa1)
	fmt.Println(p2)

	fmt.Printf("%T \t %#v\n", sa1, sa1)

	fmt.Println(sa1.first, sa1.last, sa1.age)

	fmt.Println(p3)
}