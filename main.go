package main

import (
	"bai3/database"
	"bai3/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Kết nối đến cơ sở dữ liệu
	database.ConnectDB()

	// Tạo router
	r := gin.Default()

	// Đăng ký các route
	routes.RegisterStudentRoutes(r)

	// Chạy server
	r.Run(":8080")
}
