package main

import "fmt"
import "math"

const s string = "constants"

func main() {
	fmt.Println(s)
	const n = 5000000
	fmt.Println(n)
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
