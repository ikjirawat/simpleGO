package main

import "fmt"

func main() {
	var aint = 128567
	var bfloat = 70.43
	var r rune = 'F'

	fmt.Println(r)
	fmt.Printf("Rune: %c \t Type: %T\n", aint, aint)
	fmt.Printf("Rune: %.3f \t Type: %T\n", bfloat, bfloat)

	// -------------------------------------------------------------//

	test := `just testing ei ei
	nahhh
	test not di farkkkk
	eieiei
	eiei`
	fmt.Println(test)

	// -------------------------------------------------------------//

	var title, year, rate, genre, isSuperhero = "Avengers: Endgame", 2019, 8.4, "Sci-Fi", true
	var fav rune = '⭐'

	fmt.Printf(" เรื่อง: %s\n ปี: %d\n เรตติ้ง: %.1f\n ประเภท: %s\n ซุปเปอร์ฮีโร่: %t\n เรื่องโปรด: %c", title, year, rate, genre, isSuperhero, fav)

	// -------------------------------------------------------------//
}
