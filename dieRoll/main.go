package main

import (
	"fmt"
	"math/rand"
	"time"
)

func dieRoll(size int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(size) + 1
}
func timesFive(nums ...int) {
	for _, num := range nums {
		fmt.Println(num * 5)
	}

}
func timesSix(number, limit int) {

	for i := 0; i < limit; i++ {
		x := number * 6
		number++
		fmt.Println(x)
	}
}
func timesRandom(number, limit int) {

	for i := 0; i < limit; i++ {
		rand.Seed(time.Now().UnixNano())
		x := number*rand.Intn(number) + 1
		number++
		fmt.Println(x)
	}
}

func main() {
	fmt.Println("***********timesFive****************")
	timesFive(7, 8, 9, 5, 11, 33, 555, 21212)
	fmt.Println("**********timesSix*****************")
	timesSix(2, 10)
	fmt.Println("***********timesRandom****************")
	timesRandom(2, 5)
	fmt.Println("***********test****************")
	fmt.Printf("Rolled a die of size %d, result: %d\n", 5, dieRoll(5))
}
