package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	_ "go/Config"
)

//配置结构体，用户连接数据库
type Config struct {
	Host        string
	UserName    string
	Password    string
	Port        int
	Database    string
	MaxIdleCons int
	MaxOpenCons int
	Debug       bool
}

var DB *gorm.DB
type DbStruct struct {
	enterprise *gorm.DB
}

func (c *Config) initDB() {
	username := c.UserName       //账号
	password := c.Password       //密码
	host := c.Host               //数据库地址，可以是Ip或者域名
	port := c.Port               //数据库端口
	Dbname := c.Database         //数据库名
	timeout := "10s"             //连接超时，10秒
	MaxIdleCons := c.MaxIdleCons //连接池中的最大闲置连接数
	MaxOpenCons := c.MaxOpenCons //设置数据库的最大连接数量
	Debug := c.Debug             //是否打印sql

	var err error
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	fmt.Println(dsn)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	// 连接数配置也可以写入配置，在此读取
	DB.DB().SetMaxIdleConns(MaxIdleCons)
	DB.DB().SetMaxOpenConns(MaxOpenCons)
	//调试模式
	DB.LogMode(Debug)      //打印sql
	DB.SingularTable(true) //全局禁用表名复数
	fmt.Println("初始化数据库成功......")
}

func init() {
	c := &Config{}
	c.UserName = viper.GetString("database.UserName")
	c.Password = viper.GetString("database.Password")
	c.Host = viper.GetString("database.Host")
	c.Port = viper.GetInt("database.port")
	c.Database = viper.GetString("database.Database")
	c.MaxIdleCons = viper.GetInt("database.MaxIdleCons")
	c.MaxOpenCons = viper.GetInt("database.MaxOpenCons")
	c.Debug = viper.GetBool("database.Debug")
	c.initDB()
}

func (c *Config) ConnectDb() {
	c.initDB()
}
