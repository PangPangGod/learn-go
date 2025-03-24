package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pangpanggod/learn-go/util"

	"rsc.io/quote"
)

// variable initializer
var i, j int = 1, 2
var f float64

// func add(x int, y int) int {
// 	return x + y
// }

func swap(x, y string) (string, string) {
	return y, x
}

// naked return test
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("This is New Learning-Go Project with my dirty codes")
	fmt.Println("-", quote.Go())
	
	fmt.Println("The time is", time.Now())
	fmt.Println("Random Number", rand.Intn(10))


	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

	fmt.Println(i, j)

	// short variable declaration
	k := "크아악"
	fmt.Println(k)

	// return nothing
	util.Variable_test()

	// Variables declared without an explicit initial value are given their zero value
	fmt.Printf("%v", f)
}