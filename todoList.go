package main // Khai báo package chính (chương trình bắt đầu từ đây)

import (
	"bufio"   // 📥 Dùng để đọc input từ bàn phím (đọc dòng)
	"fmt"     // 🖨 Dùng để in ra màn hình
	"os"      // 🧱 Quản lý input/output như os.Stdin
	"strconv" // 🔢 Chuyển đổi string → int (v.v.)
	"strings" // ✂️ Xử lý chuỗi: trim, split...
)

type Task struct {   // 📦 Struct định nghĩa công việc
    Title string     // 📝 Tiêu đề công việc
    Done  bool       // ✅ Trạng thái: đã hoàn thành hay chưa
}

var tasks []Task // 📋 Slice toàn cục để lưu danh sách công việc

func todoList() {   // 🚪 Hàm khởi động ứng dụng Todo List
    reader := bufio.NewReader(os.Stdin) // 🎧 Tạo trình đọc input từ bàn phím

    for { // 🔁 Vòng lặp vô hạn để hiển thị menu liên tục

        // 🧾 Hiển thị menu
        fmt.Println("\n====== MENU TODO CLI ======")
        fmt.Println("1. 📃 Xem danh sách công việc")
        fmt.Println("2. ➕ Thêm công việc mới")
        fmt.Println("3. ✅ Đánh dấu hoàn thành")
        fmt.Println("4. ❌ Thoát")
        fmt.Print("👉 Chọn chức năng (1–4): ")

        // 📥 Đọc lựa chọn từ người dùng
		// khai báo mới thú vị 
		// trong go nó cho phép chúng ta trả về nhiều giá trị
		//ví dụ ở dưới đây choiceStr, err := reader.ReadString('n') ==> trả về 2 giá trị: 1 là choiceStr, 2 là err
		//còn khai báo _ ở dưới để bỏ qua giá trị lỗi
		// tại sao dùng choiceStr, _ := reader.ReadString('\n')   mà không dùng   choiceStr:= reader.ReadString('\n') ==> reader.ReadString('\n') trả về 2 giá trị    
		
		//==========================================Xử lí lỗi with go===========================================================
		// value, err := someFunc()
		//if err != nil {
        //fmt.Println("❌ Có lỗi xảy ra:", err)
        //return // hoặc xử lý theo ý bạn
         //}

        choiceStr, _ := reader.ReadString('\n')          // Nhập chuỗi
        choiceStr = strings.TrimSpace(choiceStr)         // Cắt bỏ dấu xuống dòng
        choice, _ := strconv.Atoi(choiceStr)             // Chuyển sang số nguyên

        // 🔄 Xử lý lựa chọn
        switch choice {
        case 1:
            listTask()             // Xem danh sách
        case 2:
            addTask(reader)        // Thêm công việc
        case 3:
            markDone(reader)       // Đánh dấu hoàn thành
        case 4:
            fmt.Println("👋 Bái bai") // Thoát
            return                    // Thoát khỏi vòng lặp và hàm
        default:
            fmt.Println("⚠️ Lựa chọn không hợp lệ.")
        }
    }
}

func listTask() { // 📋 Hàm in danh sách công việc
    if len(tasks) == 0 {
        fmt.Println("📭 Chưa có công việc nào")
        return
    }

    // 👓 In danh sách
    fmt.Println("================Danh sách công việc=================")
    for i, t := range tasks { // Duyệt từng task trong danh sách
        status := "❌ Chưa hoàn thành"
        if t.Done {
            status = "✅ Đã hoàn thành"
        }
        fmt.Printf("%d. %s - %s\n", i+1, t.Title, status)
    }
}

func addTask(reader *bufio.Reader) { // ➕ Thêm công việc mới
    fmt.Println("🆕 Nhập tên công việc:")
    title, _ := reader.ReadString('\n')   // Nhập tiêu đề
    title = strings.TrimSpace(title)      // Cắt khoảng trắng/thừa

    task := Task{Title: title, Done: false} // Tạo công việc mới
    tasks = append(tasks, task)             // Thêm vào danh sách
    fmt.Println("✅ Đã thêm:", title)
}

func markDone(reader *bufio.Reader) { // ✅ Đánh dấu hoàn thành
    listTask()  // In lại danh sách trước khi chọn

    fmt.Print("🔢 Nhập số thứ tự công việc muốn đánh dấu hoàn thành: ")
    indexStr, _ := reader.ReadString('\n')
    indexStr = strings.TrimSpace(indexStr)
    index, err := strconv.Atoi(indexStr)

    // Kiểm tra hợp lệ
    if err != nil || index < 1 || index > len(tasks) {
        fmt.Println("⚠️ Số thứ tự không hợp lệ.")
        return
    }

    tasks[index-1].Done = true // ✅ Gán trạng thái hoàn thành
    fmt.Println("🎉 Đã đánh dấu hoàn thành:", tasks[index-1].Title)
}
