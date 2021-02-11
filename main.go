package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 0
	var b int = 1
	var res int
	var p = fmt.Println

	/* Logical operators
	1 as true
	0 as false
	*/

	//And
	res = a & b
	p(a, "&", b, ":", res)

	//Or
	res = a | b
	p(a, "|", b, ":", res)

	//XOr //Returns 1, when a or b are 1, but not both
	res = a ^ b
	p(a, "^", b, ":", res)

	a = 9
	b = 2
	/*Displacement operators
	number to be moved, and number of positions
	*/

	//Shift left
	res = a << b
	p(convert(a), "<<", convert(b), ":", convert(res))

	//Sign-propagating right shift (keep the sign)
	res = a >> b
	p(convert(a), ">>", convert(b), ":", convert(res))
}

func convert(n int) string {
	var str = make([]int, 0, 5)
	var aux string

	for n >= 1 {
		str = append(str, (n % 2))
		n = n / 2
	}

	for i := len(str) - 1; i >= 0; i-- {
		aux += strconv.Itoa(str[i])
	}

	return aux
}
