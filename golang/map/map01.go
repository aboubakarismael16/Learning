package main

import "fmt"

func main() {
	dict := make(map[string]int)

	dict["banana"]= 10


	fmt.Println(dict)
	fmt.Printf("type %T\n", dict)

	// a nil map can not be used to store key/value pairs.
	var color map[string]string

	value, exists := color["blue"]

	if exists {
		fmt.Println(value)
	}

	shop := map[string]int{
		"banana":23,
		"apple": 12,
		"egg" : 4,
		"orange": 33,
	}

	for key, value := range shop {
		fmt.Printf("key = %s, value = %v\n", key, value)
	}

	delete(shop, "egg")
	removeShop(shop, "apple")
	fmt.Println("after delete....")

	for key, value := range shop {
		fmt.Printf("key = %s, value = %v\n", key, value)
	}
}

func removeShop(shop map[string]int,key string)  {
	delete(shop,  key)
}

//there is nothing stopping you from using slice as a map value.