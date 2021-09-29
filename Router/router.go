package Router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	product "go/App/Controllers/product/route"
	"go/Middlewares/corss"
	"net/http"
)

type Option func(*gin.Engine)

var options = []Option{}

// 注册app的路由配置
func include(opts ...Option) {
	options = append(options, opts...)
}

// 初始化
func Init() *gin.Engine {
	//组合路由
	include(product.LoadTestRouters)

	r := gin.Default()

	// 404处理
	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("%s %s not found", method, path),
			"code":  404,
			"data":    "",
		})
	})

	//跨域中间件
	r.Use(corss.Cors())

	//加载各模块路由
	for _, opt := range options {
		opt(r)
	}
	return r
}