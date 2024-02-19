package routes

import(
	"context"
	"github.com/DilipR14/Productom.git/controllers"
	"github.com/DilipR14/Productom.git/routes"
	"github.com/gin-gonice/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/user/signup", controllers.SignUp())
	incomingRoutes.POST("/user/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/user/productview", controllers.SearchProduct())
	incomingRoutes.GET("/user/search", controllers.SearchProductByQuery())
} 	
	