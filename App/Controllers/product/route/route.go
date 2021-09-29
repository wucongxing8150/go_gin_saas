package route

import (
	"github.com/gin-gonic/gin"
	"go/App/Controllers/product"
)

func LoadTestRouters(e *gin.Engine) {
	orderGroup := e.Group("/product")
	{
		testGroup := orderGroup.Group("/test")
		{
			testGroup.GET("", product.GetTest)
		}
	}
}