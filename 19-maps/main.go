package main

import "fmt"

// slice, map , chan --> make to instantiate it
// map is a reference type
// key of a map can be any type; on which ever the type == operator is used.
// if map is not instantiated it is nil
func main() {

	// key can be of any thing where you can directly use == operator
	// string , int, bool ,float64
	var mymap map[string]any // only declaration of map not instatiated.

	if mymap == nil {
		fmt.Println("mymap is nil")
	}

	mymap = make(map[string]any)

	mymap["trivandrum-1"] = 590034
	mymap["bangalore-1"] = 560086
	mymap["hyderabad-2"] = 543234
	mymap["bangalore-2"] = 560032
	//["bangalore-1"] = 560076

	for key, val := range mymap {
		fmt.Println("Key:", key, "Value:", val)
	}

	val, ok := mymap["bangalore-1"]
	//val1 := mymap["bangalore-2"]
	if ok {
		fmt.Println("Key already there,so the value", val)
	}

	val, ok = mymap["bangalore-3"]
	//val1 := mymap["bangalore-2"]
	if ok {
		fmt.Println("Key already there,so the value", val)
	} else {
		fmt.Println("Key not there")
	}

	delete(mymap, "bangalore-1")

	// mymap1 := mymap

}

// perform deep copy of a map
// update an element from the map, return whether it has updated or not. There should
// be an error parameter to tell that if the key does not exist shoudl return "unable to update due to no key"

// delete an element from the map , it should tell whether deleted or not.
// check nil conditions for all functions.. if nil then return as an error
