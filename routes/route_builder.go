package routes

import "github.com/gin-gonic/gin"

type Routes struct {
	router *gin.Engine
}

func ProvideRoutes() *gin.Engine {
	r := Routes{router: gin.Default()}
	apiGroup := r.router.Group("/api")

	r.setStudentRoutes(apiGroup)
	return r.router
}

func (r Routes) Run(addr string) error {
	return r.router.Run(addr)
}
