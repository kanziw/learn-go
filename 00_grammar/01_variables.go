package main

import "fmt"

// Named results
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // return required
}

func main() {
	fmt.Println(split(17))
}
