package main

import (
	"fmt"
	"math"
)

func equal(a, b float64) bool {
	return math.Nextafter(a, b) == b
}

func main() {

	var x int8 = -1
	var y int8 = -128
	fmt.Printf("x: %08b\n", x)
	fmt.Printf("y: %08b\n", y)

	/*
		1의 보수 11111111 - x = -x
		2의 보수 100000000 - x = -x : 1의 보수에서 +1을 한것과 같은 수치이다.
	*/

	/*
		int8 : -128 ~ 127
		uint8 : 0 ~ 255
	*/

	/*
		실수오차 : 완벽하게 극복하는 방법은 없다.
		실수 : 부호 | 지수 | 소수
	*/

	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%f + %f == %f : %v\n", a, b, c, a+b == c)
	fmt.Println(a + b)

	fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, equal(a+b, c))

}
