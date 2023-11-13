package router

import (
	"net/http"

	"github.com/sulerasyid/go-crud/controller"

	"github.com/gin-gonic/gin"
)

func DummyMiddleware() gin.HandlerFunc {
	// Do some initialization logic here
	// Foo()
	return func(c *gin.Context) {

		c.Next()
	}
}

func NewRouter(tagController *controller.TagController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	service.Use(DummyMiddleware())

	router := service.Group("/api")
	service.Use()
	tagRouter := router.Group("/tag")
	tagRouter.GET("", tagController.FindAll)
	tagRouter.GET("/:tagId", tagController.FindById)
	tagRouter.POST("", tagController.Create)
	tagRouter.PATCH("/:tagId", tagController.Update)
	tagRouter.DELETE("/:tagId", tagController.Delete)

	return service
}
