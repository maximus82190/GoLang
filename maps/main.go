package main

import "fmt"

func main() {
	am := map[string]int {
		"Todd": 42,
		"Henry": 16,
		"Padget": 14,
	}

	fmt.Println("The age of Henry was ", am["Henry"])

	fmt.Println(am)
	fmt.Printf("%#v\n", am)
	
	an := make(map[string]int)
	an["Lucas"] = 28
	an["Steph"] = 27
	
	fmt.Println(an)
	fmt.Printf("%#v\n", an)

	fmt.Println(len(an))

// FOR RANGE LOOP OVER A MAP
	for k, v := range an { // // k = key, v = value
		fmt.Println(k, v)
	}

	for _, v := range an { // // k = key, v = value
		fmt.Println(v)
	// PRINTS OUT ONLY THE VALUES
	}

	for k := range an { // // k = key, v = value
		fmt.Println(k)
	// PRINTS OUT ONLY THE KEYS
	} 

// DELETE ELEMENT FROM MAP
	delete(an, "George")

// COMMA OKAY IDIOM
	v, ok := an["Georgey"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key didn't exist")
	}
	//  CAN BE CONSOLIDATED TO:
	if v, ok := an["Georgey"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key didn't exist")
	}
	// ALSO CAN BE EXPRESSED AS:
	if v, ok := an["Georgey"]; !ok {
		fmt.Println("Key didn't exist")
	} else {
		fmt.Println(v)
	}


}