package rect

import (
	"fmt"
	"math/rand"
	"time"
)

func Greet() string {
	return "Hello World"
}

func ShowTime() {
	fmt.Println(time.Now())
}

// This is not exporeted function becase it starts with lowercase
func sayHey() {
	fmt.Println("Hey how are yhou doing!-> This is your token number-->", rand.Intn(2011))
}
