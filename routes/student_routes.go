package routes

import (
	"github.com/gin-gonic/gin"
	"go-gin/data"
	"net/http"
)

func (r Routes) setStudentRoutes(rg *gin.RouterGroup) {
	studentRoutes := rg.Group("students")

	studentRoutes.GET("/", func(ctx *gin.Context) {
		res := data.AllStudents
		ctx.JSON(http.StatusOK, map[string]any{"data": res})
	})

	studentRoutes.GET("/:studentId", func(ctx *gin.Context) {
		studentId := ctx.Param("studentId")

		var res data.Student
		for _, student := range data.AllStudents {
			if student.StudentId == studentId {
				res = student
				break
			}
		}

		if res.StudentId == "" {
			ctx.JSON(http.StatusNotFound, map[string]any{"message": "NOT FOUND"})
			return
		}
		ctx.JSON(http.StatusOK, map[string]any{"data": res})
	})
}
