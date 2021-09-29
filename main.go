package main

import (
	"fmt"
	"github.com/spf13/viper"
	_ "go/Config" //初始化读取配置文件
	_ "go/Cores/mysql"
	"go/Router"
	"strconv"
)
func main() {
	// 初始化路由
	r := Router.Init()
	if err := r.Run(":" + strconv.Itoa(viper.GetInt("port"))); err != nil {
		fmt.Println("启动失败:%v\n", err)
	}
}
