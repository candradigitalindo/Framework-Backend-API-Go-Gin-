package routes

import (
	"candra/backend-api/controllers"
	"candra/backend-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.POST("/login", controllers.Login)
		api.POST("/register", controllers.Register)
		role := api.Group("/role", middlewares.AuthMiddleware())
		{
			role.GET("/", controllers.GetAllRoles)
			role.GET("/:id", controllers.GetRoleByID)
			role.POST("/", controllers.CreateRole)
			role.PUT("/:id", controllers.UpdateRole)
			role.DELETE("/:id", controllers.DeleteRole)
		}
		user := api.Group("/user", middlewares.AuthMiddleware())
		{
			user.GET("/", controllers.GetAllUsers)
			user.GET("/:id", controllers.GetUserByID)
			user.PUT("/:id", controllers.UpdateUser)
			user.DELETE("/:id", controllers.DeleteUser)
		}

		// Tambahkan endpoint lain di sini
	}
	return r
}
