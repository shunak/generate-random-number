package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

//show random number from 0 to argument value (default (no argument case) from 0 to 2)
func main() {
	rand.Seed(time.Now().UnixNano())
	var intOpt = flag.Int("n", 3, "Specify the number")
	flag.Parse()
	fmt.Println(rand.Intn(*intOpt))
	// show random number from 0 to 99
	// fmt.Println(rand.Intn(100))
}
