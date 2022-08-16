package main

import "fmt"

func main() {
	// for(int i=0;)
	for i := 0; i < 10; i++ {
		print(i, " ")
	}
	println()
	j := 0
	for j < 10 { // similar to while loop
		print(j, " ")
		j++ // j=j+1
	}
	println()
	k := 0
	for {
		if k == 10 {
			break
		}
		print(k, " ")
		k++
	}

	println()
	for i := 2; i <= 20; i += 2 {
		if i%2 != 0 {
			continue
		}
		print(i, " ")
	}

	println()
	for i := 1; i <= 20; i++ {
		if i%2 != 0 {
			continue
		}
		print(i, " ")
	}

	// for i := 0; i < 5; i++ {
	// 	for j := 0; j < 5; j++ {
	// 		println("i:", i, " j:", j)
	// 	}
	// }
	println()
	// using break in nested loops
outer:
	for i := 0; i <= 5; i++ {
		for j := 2; j <= 5; j++ {
			if i == j {
				break outer
			}
			println("i:", i, " j:", j)
		}
	}

	fmt.Println("Hello World")

}

// print 1-100 prime numbers
