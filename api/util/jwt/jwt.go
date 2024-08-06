package jwt

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"hbbapi/app/enum/error_enum"
	"io/ioutil"
	"os"
	"time"
)

type CustomClaims struct {
	Uid int `json:"uid"`
	jwt.StandardClaims
}

// Sign 定义一个 JWT验签 结构体
type Sign struct {
	SigningKey *rsa.PublicKey
}

var SignObj Sign

func Init() {
	SignObj = Sign{}
	env := os.Getenv("GIN_ENV")
	fileObj, _ := ioutil.ReadFile(fmt.Sprintf("./util/jwt/%s/jwt_public_key.pem", env))
	SignObj.SigningKey, _ = jwt.ParseRSAPublicKeyFromPEM(fileObj)
}

// CreateToken 生成一个token
func (j *Sign) CreateToken(claims CustomClaims) (string, error) {
	// 生成jwt格式的header、claims 部分
	tokenPartA := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 继续添加秘钥值，生成最后一部分
	return tokenPartA.SignedString(j.SigningKey)
}

// ParseToken 解析Token
func (j *Sign) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return j.SigningKey, nil
		})
	if token == nil {
		return nil, errors.New(error_enum.ErrorsTokenInvalid)
	}
	if err != nil {
		fmt.Println(err.Error())
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New(error_enum.ErrorsTokenMalFormed)
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New(error_enum.ErrorsTokenNotActiveYet)
			} else {
				return nil, errors.New(error_enum.ErrorsTokenInvalid)
			}
		}
	}
	claims, _ := token.Claims.(*CustomClaims)
	return claims, nil
}

// RefreshToken 更新token
func (j *Sign) RefreshToken(tokenString string, extraAddSeconds int64) (string, error) {
	if CustomClaims, err := j.ParseToken(tokenString); err == nil {
		CustomClaims.ExpiresAt = time.Now().Unix() + extraAddSeconds
		return j.CreateToken(*CustomClaims)
	} else {
		return "", err
	}
}
