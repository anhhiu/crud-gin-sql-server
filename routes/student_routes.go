package routes

import (
	"bai3/controllers"

	"github.com/gin-gonic/gin"
)

// Đăng ký route cho sinh viên
func RegisterStudentRoutes(router *gin.Engine) {
	api := router.Group("/api/students")
	{
		api.Use(controllers.JWTAuthMiddleware())
		api.GET("/phantrang", controllers.GetStudentsWithPagination)
		api.GET("/thongke", controllers.GetStudentStatistics)
		api.GET("/search", controllers.SearchStudents)
		api.GET("/", controllers.GetStudents)
		api.GET("/:id", controllers.GetStudentById)
		api.POST("/", controllers.CreateStudent)
		api.PUT("/:id", controllers.UpdateStudent)
		api.DELETE("/:id", controllers.DeleteStudent)
	}
}
