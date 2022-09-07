package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kritsanaphat/book-ecm-golang/controllers"
)

func UserRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.POST("/user/signup", controllers.SignUp())
	incommingRoutes.POST("/user/login", controllers.Login())
	incommingRoutes.POST("/add/addproduct", controllers.ProductViewAdmin())
	incommingRoutes.GET("/user/productview", controllers.SearchProduct())
	incommingRoutes.GET("/user/search", controllers.SearchProductByQuery())

}
