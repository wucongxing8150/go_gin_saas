package product

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go/Tool"
	"net/http"
	"time"
)

func GetTest(c *gin.Context)  {
	//生成token
	/*user := &Tool.User{
		AuthId: "1",
		UserId: "2",
	}

	token := user.CreateRefreshToken()
	fmt.Println(user)
	fmt.Println(token)*/

	//解析token
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoX2lkIjoiMSIsInVzZXJfaWQiOiIyIiwic3RvcmVfaWRzIjoiIiwic3RhdGlvbl9pZHMiOiIiLCJlbnRlcnByaXNlX2lkIjoiIiwidmVuZGVyX2lkcyI6bnVsbCwidXNlcl9sb2dpbl9pcCI6IiIsInVuaXF1ZV9kZXZpY2VfY29kZV90YiI6IiIsInVuaXF1ZV9kZXZpY2VfY29kZV90bSI6IiIsInVuaXF1ZV9kZXZpY2VfY29kZV9wZGQiOiIiLCJhdWQiOiJ1c2VyIiwiZXhwIjoxNjMwNTYwNzE2LCJpYXQiOjE2MzA0NzQzMTYsImlzcyI6InN5c3RlbSIsIm5iZiI6MTYzMDQ3NDMxNn0.0y0cwT3bN1uD0nhMfhccBe3x3pW08RzZl76zPvAx2OU"

	var result = Tool.ParseToken(token)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"status":  1,
		"error":   "",
		"data":    result,
	})

	/*config := &mysql.Config{}
	config.Host = "123.57.91.25"
	config.UserName = "root"
	config.Password = "Wucongxing1.."
	config.Port = 3306
	config.Database = "YhEShopPlatform"
	config.MaxIdleCons = 5
	config.MaxOpenCons = 10
	config.Debug = true
	//fmt.Printf("p1=%v\n", config)
	config.ConnectDb()*/


	date := time.Now()
	fmt.Println(date)
	//查询
	/*maps := make(map[string]interface{})
	maps["user_name"] = "你好"
	result := Models.Get(maps)
	fmt.Println(result)*/
	//修改
	/*maps := make(map[string]interface{})
	maps["user_name"] = "你好"
	maps["update_time"] = date
	result := Models.EditId(1,maps)
	fmt.Println(result)*/
	//新增
	/*data := make(map[string]interface{})
	data["user_name"] = "name"
	data["user_phone"] = "phone"
	result := Models.Add(data)*/
	//fmt.Println(result)

	/*c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"status":  1,
		"error":   "",
		"data":    result,
	})*/
}