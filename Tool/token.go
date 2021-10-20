package Tool

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"time"
)

// 一些常量
var (
	TokenExpired     = "token过期"
	TokenMalformed   = "token解析错误" //非token数据
	TokenInvalid     = "无法解析token"
	TokenNotValidYet = "token签名错误"
	TokenEmpty       = "token解析错误"
	SignKey          = viper.GetString("tokenKey")
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

//解析token
func ParseToken(tokenString string) (message string, code int) {

	//解析、验证并返回令牌
	j := new(JWT)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// alg:加密方式
		// HS256：一种对称加密算法，使用同一个秘钥对signature加密解密
		// RS256：一种非对称加密算法，使用私钥加密，公钥解密
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return j.SigningKey, nil
	})

	/*token, err := jwt.ParseWithClaims(tokenString, &JwtRefreshData{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})*/

	if token == nil {
		return TokenEmpty, 0
	}

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return TokenMalformed, 0

			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return TokenExpired, -1

			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return TokenNotValidYet, 0

			} else {
				return TokenInvalid, 0
			}
		}
	}
	if result, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("nihao")
		fmt.Println(result)
	}
	return TokenInvalid, 0
}
