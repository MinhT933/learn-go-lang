package main

import (
	"fmt"
)

func main() {
    todoList()
	name:= "Toàn"
	age:= 23
	hobby:= "học Go  mỗi sáng"
    score := 8.75
	number:= 10 
	count:=5
	sum:= 0
	//Println log xuống dòng, Print không xuống dòng
	// Printf là format string có các place holder %
	// %s string
	// %d int
	// %f float64
	// %v in ra bất kì giá trị nào
	
    fmt.Printf("👋 Xin chào %s\n", name)
    fmt.Printf("🎂 Tuổi của bạn là: %d\n", age)
    fmt.Printf("📊 Điểm lập trình Go: %.2f\n", score)
	fmt.Printf("sở thích: %s\n", hobby)



	/// làm quen với if else
  

	if age >= 18 {
		fmt.Println("✅ Bạn đủ tuổi rồi")
	}else {
		fmt.Println("❌ chưa đủ tuổi uống rượi bia")
	}

	if number%2 ==0 {
      println("Số này là số chẵn")
	}else{
	  println("Số này là số lẽ")
	}
	 // for- vòng lặp
	 for i := 0; i < count; i++ {
		fmt.Println("lặp lần thứ", i)
	 }

	 // tính tổng với for
	//  for i := 0; i < count; i++ {
	// 	sum+=i
	
	//  }
	//  	fmt.Println("tổng trong", sum)

    // tính cổng số lẽ từ 0 đến `100`
	for i := 0; i < count; i++ {
		if(i%2 == 0){
			sum +=i
		}
	}
	fmt.Println("tổng số chẵn", sum)

	// in ra tam giác số
    // đếm giảm
	for i := 0;  i<= count; i++ {
		for j:= 0; j<= i; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}

	//đếm tăng
	// for i := count;  i >= 1; i-- {
	// 	for j:= 1; j <= i; j++ {
	// 		fmt.Printf("%d ", j)
	// 	}
	// fmt.Println()


	//Struct -- kiểu dữ liệu tự định nghĩa

	type User struct {
		Name string
		Age int
	}

	u := User{Name:"Tường", Age: 23}

	fmt.Printf("🫂 người dùng: %s (%d tuổi)\n", u.Name, u.Age)


	// làm quen với kiểu dữ liệu khai bảo mảng trong go

	type Task struct {
       Title string
	   Done bool
	}

	tasks:= []Task{
		{Title: "Học go", Done: false},
	    {Title: "Học go", Done: true},
	}
    for _, task := range tasks {
	    fmt.Printf("📌 Việc: %s | Trạng thái: ", task.Title)
		if task.Done {
			fmt.Println("🎉 Đã hoàn thành")
		}else {
			fmt.Println("🕧 Chưa hoàn thành")
		}
	}

	// check(duyet mảng) dựa theo key và value

	roles:= map[string]string {
		"admin": "Quản trị viên",
		"use": "Người dùng",
	}
	fmt.Println("👑 vai trò", roles["admin"])
}


