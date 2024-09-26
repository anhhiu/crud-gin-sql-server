package controllers

import (
	"bai3/database"
	"bai3/models"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Lấy danh sách sinh viên
func GetStudents(c *gin.Context) {
	var students []models.Student
	db := database.GetDB()

	rows, err := db.Query("SELECT Id, Ten, Tuoi, Lop, DiemTrungBinh FROM SinhVien")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var student models.Student
		if err := rows.Scan(&student.Id, &student.Ten, &student.Tuoi, &student.Lop, &student.DiemTrungBinh); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		students = append(students, student)
	}

	c.JSON(http.StatusOK, students)
}

// Thêm sinh viên mới
func CreateStudent(c *gin.Context) {
	var newStudent models.Student
	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB()
	_, err := db.Exec("INSERT INTO SinhVien (Ten, Tuoi, Lop, DiemTrungBinh) VALUES (@p1, @p2, @p3, @p4)",
		newStudent.Ten, newStudent.Tuoi, newStudent.Lop, newStudent.DiemTrungBinh)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Sinh viên đã được thêm"})
}

// Cập nhật sinh viên
func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB()
	_, err := db.Exec("UPDATE SinhVien SET Ten=@p1, Tuoi=@p2, Lop=@p3, DiemTrungBinh=@p4 WHERE Id=@p5",
		student.Ten, student.Tuoi, student.Lop, student.DiemTrungBinh, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sinh viên đã được cập nhật"})
}

// Xóa sinh viên
func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	db := database.GetDB()

	_, err := db.Exec("DELETE FROM SinhVien WHERE Id=@p1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sinh viên đã bị xóa"})
}

func SearchStudents(c *gin.Context) {
	name := c.Query("name")
	class := c.Query("class")
	ageStr := c.Query("age")
	minMarkStr := c.Query("min_avgmark")
	maxMarkStr := c.Query("max_avgmark")

	var students []models.Student
	db := database.GetDB()

	query := "SELECT Id, Ten, Tuoi, Lop, DiemTrungBinh FROM SinhVien WHERE 1=1"
	args := []interface{}{}

	if name != "" {
		query += " AND Ten LIKE @p1"
		args = append(args, "%"+name+"%")
	}
	if class != "" {
		query += " AND Lop = @p2"
		args = append(args, class)
	}
	if ageStr != "" {
		age, _ := strconv.Atoi(ageStr)
		query += " AND Tuoi = @p3"
		args = append(args, age)
	}
	if minMarkStr != "" {
		minMark, _ := strconv.ParseFloat(minMarkStr, 64)
		query += " AND DiemTrungBinh >= @p4"
		args = append(args, minMark)
	}
	if maxMarkStr != "" {
		maxMark, _ := strconv.ParseFloat(maxMarkStr, 64)
		query += " AND DiemTrungBinh <= @p5"
		args = append(args, maxMark)
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var student models.Student
		if err := rows.Scan(&student.Id, &student.Ten, &student.Tuoi, &student.Lop, &student.DiemTrungBinh); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		students = append(students, student)
	}

	c.JSON(http.StatusOK, students)
}

// phân trang

func GetStudentsWithPagination(c *gin.Context) {
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)

	var students []models.Student
	db := database.GetDB()

	query := "SELECT Id, Ten, Tuoi, Lop, DiemTrungBinh FROM SinhVien ORDER BY Id OFFSET @p1 ROWS FETCH NEXT @p2 ROWS ONLY"
	rows, err := db.Query(query, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var student models.Student
		if err := rows.Scan(&student.Id, &student.Ten, &student.Tuoi, &student.Lop, &student.DiemTrungBinh); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		students = append(students, student)
	}

	c.JSON(http.StatusOK, students)
}

// XEM CHI TIET ID
func GetStudentById(c *gin.Context) {
	id := c.Param("id")

	var student models.Student
	db := database.GetDB()

	err := db.QueryRow("SELECT Id, Ten, Tuoi, Lop, DiemTrungBinh FROM SinhVien WHERE Id = @p1", id).Scan(&student.Id, &student.Ten, &student.Tuoi, &student.Lop, &student.DiemTrungBinh)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sinh viên không tồn tại"})
		return
	}

	c.JSON(http.StatusOK, student)
}

// THỐNG KÊ SINH VIÊN

func GetStudentStatistics(c *gin.Context) {
	var total int
	var maxMark, minMark, avgMark float64
	db := database.GetDB()

	err := db.QueryRow("SELECT COUNT(*), MAX(DiemTrungBinh), MIN(DiemTrungBinh), AVG(DiemTrungBinh) FROM SinhVien").Scan(&total, &maxMark, &minMark, &avgMark)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_students": total,
		"max_mark":       maxMark,
		"min_mark":       minMark,
		"avg_mark":       avgMark,
	})
}

// JWTAuthMiddleware kiểm tra tính hợp lệ của JWT
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy token từ header Authorization
		token := c.GetHeader("Authorization")

		// Kiểm tra xem token có tồn tại hay không
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token không được cung cấp"})
			c.Abort() // Ngăn chặn việc tiếp tục xử lý request
			return
		}

		// Kiểm tra token (giả sử token phải bắt đầu bằng "Bearer ")
		if !strings.HasPrefix(token, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token không hợp lệ"})
			c.Abort()
			return
		}

		// Lấy token thực tế (bỏ "Bearer " ra)
		token = strings.TrimPrefix(token, "Bearer ")

		// TODO: Kiểm tra tính hợp lệ của token ở đây (giả định sử dụng một hàm kiểm tra token)
		if !isValidToken(token) { // Giả sử là hàm kiểm tra token
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token không hợp lệ"})
			c.Abort()
			return
		}

		// Nếu token hợp lệ, tiếp tục xử lý request
		c.Next()
	}
}

// isValidToken là hàm kiểm tra tính hợp lệ của token (bạn cần định nghĩa hàm này)
func isValidToken(token string) bool {
	fmt.Println(token)
	// Thực hiện kiểm tra token ở đây (có thể là giải mã token và kiểm tra thông tin)
	// Trả về true nếu token hợp lệ, false nếu không
	return true // Ví dụ: luôn trả về true (thay thế bằng logic kiểm tra thực tế)
}
