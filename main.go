package main

import (
	"net/http"

	"example.com/splitwise_backend/controllers"
	"example.com/splitwise_backend/middlewares"
	"example.com/splitwise_backend/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	models.ConnectDataBase()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: false,
	}))
	public := r.Group("/api")
	{
		public.POST("/register", controllers.Register)
		public.POST("/login", controllers.Login)
	}

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	{
		protected.GET("/profile", controllers.CurrentUser)

		protected.POST("/creategroup", controllers.CreateGroup)
		protected.POST("/creategroupmember", controllers.CreateGroupMember)
		protected.POST("/expensivecreate", controllers.ExpenseCreate)
		protected.POST("/expensivelistcreate", controllers.ExpenseListCreate)

		protected.GET("/groups", controllers.GetGroupByuserId)

	}

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "server running"})
	})

	r.Run(":8080")
}
