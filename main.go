package main

import (
	"dev_sc/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// router.StaticFile("/", "client/public/index.html")

	v1 := router.Group("/v1")
	{
		student := new(controllers.StudentController)

		v1.POST("/create", student.Create)
		v1.GET("/fetch", student.Find)
		v1.GET("/students/:id", student.Get)
		v1.POST("/update", student.Update)
		v1.POST("/delete", student.Delete)
	}

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})

	//router.Use(logger.SetLogger())

	router.Run(":9090")
}
