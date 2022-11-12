package main

import "fmt"

func main() {
	var aint = 128567
	var bfloat = 70.43
	var r rune = 'F'

	test := `just testing ei ei
	nahhh
	test not di farkkkk
	eieiei
	eiei`
	fmt.Println(test)

	fmt.Println(r)
	fmt.Printf("Rune: %c \t Type: %T\n", aint, aint)
	fmt.Printf("Rune: %.3f \t Type: %T\n", bfloat, bfloat)
}
