package main

import "fmt"

func main() {
	var u1 Unique
	u1.int = 101
	u1.string = "Code-1"
	u1.bool = false
	fmt.Println(u1)
	u2 := Unique{102, "Code-2", true}
	fmt.Println(u2)

	// anonymous struct
	a1 := struct {
		ID   int
		Name string
	}{
		ID:   101,
		Name: "Jiten",
	}

	fmt.Println(a1)

	a2 := struct {
		int
		string
	}{
		int:    102,
		string: "Jiten",
	}
	fmt.Println(a2)

}

// struct with not names but fields
type Unique struct {
	int
	string
	//string // cannot decalre two string types here
	bool
}
