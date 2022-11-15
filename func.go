package main

import "fmt"

func emote(ratings float64) string {
	if ratings < 5.0 {
		return "Disappointed 😞"
	} else if ratings >= 5.0 && ratings < 7.0 {
		return "Normal 😐"
	} else if ratings >= 7.0 && ratings < 10.0 {
		return "Good 🥰"
	} else {
		return "🤷🤷🤷🤷"
	}
}
func main() {
	// Workshop: function
	// กำหนด: 1.สร้างฟังก์ชันชื่อ emote และรับพารามิเตอร์หนึ่งตัวชื่อ ratings เป็นตัวแปรชนิด float64 และคืนค่าเป็น string
	// 	   ถ้า ratings มีค่าน้อยกว่า 5.0 จะทำการคืนค่าคำว่า "Disappointed 😞"
	// 	   ถ้า ratings มีค่ามากกว่าหรือเท่ากับ 5.0 และน้อยกว่า 7.0 จะทำการคืนค่าคำว่า "Normal 😐"
	// 	   ถ้า ratings มีค่ามากกว่าหรือเท่ากับ 7.0 และน้อยกว่า 10.0 จะทำการคืนค่าคำว่า "Good 🥰"
	//  	   กรณีอื่นๆ ให้ทำการคืนค่าคำว่า "🤷🤷🤷🤷"
	fmt.Println("\n\tEXAMPLE: Function")
	fmt.Println(emote(4.9))
	fmt.Println(emote(5.0))
	fmt.Println(emote(7.0))
	fmt.Println(emote(15.0) + "\n")
	// Output:
	// Disappointed 😞
	// Normal 😐
	// Good 🥰
	// 🤷🤷🤷🤷

	// Workshop: array
	// กำหนด: ให้ประกาศตัวแปรอาร์เรย์ประเภทหนัง(genres)ที่เก็บค่าเป็น "Action", "Adventure" และ "Fantasy" ตามลำดับ
	// 	  ให้อ่านค่าของอาร์เรย์แต่ละช่องเพื่อแสดงผล ให้แสดงผลทั้งหมดตามตัวอย่าง Output ด้านล่าง
	// 	  หลังจากนั้นเปลี่ยนแปลงค่าในอาร์เรย์ index ที่ 1 ให้เป็น Sci-Fi พร้อมทั้งแสดงผล เพื่อยืนยันว่าค่าเปลี่ยนจริง
	fmt.Println("\tEXAMPLE: Array")
	genres := [3]string{"Action", "Adventure", "Fantasy"}
	fmt.Println("genres[0]:", genres[0])
	fmt.Println("genres[1]:", genres[1])
	fmt.Println("genres[2]:", genres[2])
	genres[1] = "Sci-Fi"
	fmt.Println("genres[1]:", genres[1])
	// Output:
	// "genres[0]: Action"
	// "genres[1]: Adventure"
	// "genres[2]: Fantasy"
	// "genres[1]: Sci-Fi"

}
