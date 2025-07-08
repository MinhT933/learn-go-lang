package handlers

import (
	"fiber-todo-api/database"
	"fiber-todo-api/models"

	"github.com/gofiber/fiber/v2"
)

//hàm này nhận về một c*fiber.Ctx ==> c*fiber.Ctx là cái gì ?????? 🤔
// nó là context của Fiber dùng để xử lý request/response
// giả sử có một cửa hàn bán bánh onl ai muốn mua gì cũng phải đặt hàng qua tờ thông tin ==> tờ thông tin ấy đây chính là c*fibe.Ctx
// nó bao gồm những thông tin cần thiết như định danh, loại bánh mì, địa chỉ,.....

// c.Body()	Đọc nội dung thư họ gửi
// c.Params("id")	Họ ghi mã đơn hàng trong phong bì
// c.Query("search")	Tìm kiếm họ muốn gì trong mẫu đơn
// c.FormFile("file")	Họ đính kèm một tấm ảnh trong thư
// c.JSON(...)	Gửi lại cho họ một lá thư phản hồi (trả kết quả)

// GetAllTasks godoc
// @Summary Lấy danh sách tất cả công việc
// @Description Trả về danh sách task hiện có
// @Tags tasks
// @Produce json
// @Success 200 {array} models.Task
// @Router /tasks [get]
func GetAllTasks(c *fiber.Ctx) error {
  var tasks []models.Task
  database.DB.Find(&tasks)
  return c.JSON(tasks)
}

// CreateTask godoc
// @Summary Tạo công việc mới
// @Description Tạo task mới với JSON
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body models.Task true "Task"
// @Success 201 {object} models.Task
// @Failure 400 {object} map[string]interface{}
// @Router /tasks [post]
func CreateTask(c *fiber.Ctx) error {
	var task models.Task
	if err:= c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{"error":"Dữ liệu không hợp lệ"})
	}
	database.DB.Create(&task)
	return c.Status(201).JSON(task)
}


// @Sumary
// @Tags tasks
// @Description  Trả về một task duy nhất dựa trên ID
// @Accept json
// @Produce json
// @Param id path int true "ID của task"
// @Success 200 {object} models.Task
// @Failure 404 {object} models.ErrorResponse
// @Router /Tasks/{id} [get]
func GetTaskByID(c *fiber.Ctx ) error {
  id := c.Params("id")

  var task models.Task
  result := database.DB.First(&task, id)

  if result.Error != nil {
	return c.Status(404).JSON(models.ErrorResponse {
		Error: "Task không tồn tại",
	})
  }

  return  c.JSON(task)
}


// @Summary
// @Tags tasks
// @Description
// @Accept json
// @Produce json
// @Param id path int true "ID của task"
// @Success 200 {object} models.Task
// @Failure 404 {object} models.ErrorResponse
// @Router /Tasks/{id} [delete]
func DeleteTaskByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var task models.Task

	// ✅ Bước 1: kiểm tra xem có tồn tại không
	if err := database.DB.First(&task, id).Error; err != nil {
		return c.Status(404).JSON(models.ErrorResponse{
			Error: "Task không tồn tại",
		})
	}

	// ✅ Bước 2: nếu tồn tại, xoá nó
	database.DB.Delete(&task)

	// ✅ Trả về task vừa xoá
	return c.JSON(task)
}

// UpdateTaskByID godoc
// @Summary Cập nhật công việc theo ID
// @Description Nhận JSON và cập nhật nội dung task
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "ID của task"
// @Param task body models.Task true "Task cần cập nhật"
// @Success 200 {object} models.Task
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /tasks/{id} [put]
func UpdateTaskByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var task models.Task
	if err := database.DB.First(&task, id).Error; err != nil {
		return c.Status(404).JSON(models.ErrorResponse{
			Error: "Task không tồn tại",
		})
	}

	// ✅ Dùng biến tạm để parse dữ liệu từ body
	var updatedData struct {
		Title string `json:"title"`
		Done  bool   `json:"done"`
	}

	if err := c.BodyParser(&updatedData); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{
			Error: "Dữ liệu không hợp lệ",
		})
	}

	// ✅ Cập nhật từng field (không đụng vào ID)
	task.Title = updatedData.Title
	task.Done = updatedData.Done

	database.DB.Save(&task)

	return c.JSON(task)
}
