package Tool

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

// 一些常量
var (
	TokenExpired     string = "token过期"
	TokenMalformed   string = "token解析错误" //非token数据
	TokenInvalid     string = "无法解析token"
	TokenNotValidYet string = "token签名错误"
	TokenEmpty       string = "token解析错误"
	SignKey          string = "123456"
)

type JWT struct {
	SigningKey []byte
}

type User struct {
	AuthId              string   `json:"auth_id"`                //用户权限
	UserId              string   `json:"user_id"`                //用户角色_id
	StoreIds            string   `json:"store_ids"`              //库房_id
	StationIds          string   `json:"station_ids"`            //用户站点
	EnterpriseId        string   `json:"enterprise_id"`          //用户企业
	VenderIds           []string `json:"vender_ids"`             //用户店铺
	UserLoginIp         string   `json:"user_login_ip"`          //用户登录ip
	UniqueDeviceCodeTb  string   `json:"unique_device_code_tb"`  //用户登录设备id
	UniqueDeviceCodeTm  string   `json:"unique_device_code_tm"`  //用户登录设备id
	UniqueDeviceCodePdd string   `json:"unique_device_code_pdd"` //用户登录设备id
}

// 载荷，access-token结构体
type JwtAccessData struct {
	jwt.StandardClaims
}

//载荷，authorization结构体
type JwtRefreshData struct {
	User
	jwt.StandardClaims
}

// 获取signKey
func GetSignKey() string {
	return SignKey
}

// 生成access-token
func CreateAccessToken() string {
	j := new(JWT)
	j.SigningKey = []byte(SignKey)

	data := &JwtAccessData{
		StandardClaims: jwt.StandardClaims{
			Audience:  "user",                          //面象的用户，可以为空
			ExpiresAt: int64(time.Now().Unix() + 7200), //过期时间
			IssuedAt:  int64(time.Now().Unix()),        //签发时间
			Issuer:    "system",                        //签发者 可以为空
			NotBefore: int64(time.Now().Unix()),        //在什么时候jwt开始生效
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	result, err := token.SignedString(j.SigningKey)
	if err != nil {
		panic(err.Error())
	}
	return result
}

// 生成Refresh-token
func (u User) CreateRefreshToken() string {
	j := new(JWT)
	j.SigningKey = []byte(SignKey)

	data := &JwtRefreshData{
		User: u,
		StandardClaims: jwt.StandardClaims{
			Audience:  "user",                           //面象的用户，可以为空
			ExpiresAt: int64(time.Now().Unix() + 86400), //过期时间
			IssuedAt:  int64(time.Now().Unix()),         //签发时间
			Issuer:    "system",                         //签发者 可以为空
			NotBefore: int64(time.Now().Unix()),         //在什么时候jwt开始生效
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	result, err := token.SignedString(j.SigningKey)
	if err != nil {
		panic(err.Error())
	}
	return result
}

func ParseToken(tokenString string) map[string]string {
	//定义返回值
	result := &person{
		name: "pprof.cn",
		city: "北京",
		age:  18,
	}

	j := new(JWT)
	token, err := jwt.ParseWithClaims(tokenString, &JwtRefreshData{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if token == nil {
		result["msg"] = TokenEmpty
		return result
	}
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				result["code"] 	= "0"
				result["msg"] = TokenMalformed
				return result
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				result["code"] = "-1"
				result["msg"] = TokenExpired
				return result
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				result["code"] = "0"
				result["msg"] = TokenNotValidYet
				return result
			} else {
				result["code"] = "0"
				result["msg"] = TokenInvalid
				return result
			}
		}
	}
	if result, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("nihao")
		fmt.Println(result)
	}
	result["code"] = "0"
	result["msg"] = TokenInvalid
	return result
}
