package main

import "math/rand"

// break is automatically applied in each case
// if you want to acheive same thing is you dont give break in other programming language
// fallthrough
func main() {
	//var day uint8 = 5
	switch day := 4; day {
	case 1:
		println("Sunday")
	case 2:
		println("Monday")
	case 3:
		println("Tuesday")
	case 4:
		println("Wednesday")
	case 5:
		println("Thursday")
	case 6:
		println("Friday")
	case 7:
		println("Saturday")
	default:
		println("no day")
	}

	for i := 1; i <= 10; i++ {
		num := rand.Intn(2000)
		switch {
		case num < 50:
			println(num, "is less than 50")
		case num >= 50 && num < 100:
			println(num, " is greater than or equal to 50 and less than 100")
		case num >= 100 && num < 200:
			println(num, " is greater than or equal to 100 and less than 200")
		default:
			println(num, "is greater than or equal to 200")
		}
	}
	str := "Hello World"
	count := 0
	//outer:
	for _, char := range str {
		switch char {
		case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
			count++
			// if count == 2 {
			// 	break outer
			// }
			println(string(char), "is an oval")
		default:
			println(string(char), "is either a consonent or a special char")
		}
	}
	println("Number of ovals in a given string is ", count)
	// var ch rune = 65
	// if ch == 'A' {
	// 	println("Yes")
	// }

}
