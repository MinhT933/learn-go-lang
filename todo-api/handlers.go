package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

//trong khi học viết api này tôi cảm thấy thắc mắc khá nhiều thứ và sau khi tìm hiểu và tự trả lời cho thắc mắc cùa mình
// w và r là gì? =>  đơn giản nó là khai báo tên biến có nghĩa là writer, reader
// tại sao phải đặt dấu * ở trước http.Rquest ==> * là con trỏ giống trong C. Con trỏ đúng như trên của nó dùng để chỉ trỏ 😊. tại sao phải để con trỏ ở chỗ này
// vì làm vậy để truy cập và sửa đổi nội dung trực tiếp của request

//name string	Sao chép nội dung → sửa trong hàm KHÔNG ảnh hưởng biến gốc
//name *string	Truyền địa chỉ → sửa trong hàm LÀM THAY ĐỔI biến gốc


func GetTodos(w http.ResponseWriter, r *http.Request) {
    // Đọc file JSON
    data, err := os.ReadFile("tasks.json")
    if err != nil {
        http.Error(w, "❌ Không thể đọc file", http.StatusInternalServerError)
        fmt.Println("Lỗi khi đọc file:", err)
        return
    }

    // (Tuỳ chọn) Parse JSON thành []Task để xử lý nếu muốn
    // var tasks []Task
    // _ = json.Unmarshal(data, &tasks)

    // Trả JSON thô về luôn
    w.Header().Set("Content-Type", "application/json")
    w.Write(data)
}

func CreateTodo(w http.ResponseWriter, r *http.Request){
	var newTask Task
	err := json.NewDecoder(r.Body).Decode(&newTask)

	if err != nil {
		http.Error(w, "❌ Dữ liệu không hợp lệ", http.StatusBadRequest)
		return
	}
	Tasks = append(Tasks, newTask)
	SaveTasksToFile()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to my API"))
}


func SaveTasksToFile() {
	data, err := json.MarshalIndent(Tasks, "", "  ")
	if err != nil {
		fmt.Println("❌ Không thể mã hóa JSON:", err)
		return
	}

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		fmt.Println("❌ Không thể ghi file:", err)
	}
}
