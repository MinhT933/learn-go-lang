package main

import "fmt"

func main() {
	name:= "Toàn"
	age:= 23
	hobby:= "học Go  mỗi sáng"
    score := 8.75
	//Println log xuống dòng, Print không xuống dòng
	// Printf là format string có các place holder %
	// %s string
	// %d int
	// %f float64
	// %v in ra bất kì giá trị nào
	fmt.Println("Xin chào", name)
	fmt.Println("Mình", age, "tuổi")
    fmt.Println("Tôi thích", hobby)
	
    fmt.Printf("👋 Xin chào %s\n", name)
    fmt.Printf("🎂 Tuổi của bạn là: %d\n", age)
    fmt.Printf("📊 Điểm lập trình Go: %.2f\n", score)
}
