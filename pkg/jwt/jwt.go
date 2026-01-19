package jwt

import (
	"errors"
	"fmt"
	"go-admin/api/entity"
	"go-admin/common/constant"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type userStdClaims struct {
	entity.JwtAdmin
	jwt.StandardClaims
}

// 过期 1天
const TokenExpiresDuration = time.Hour * 24

// token密钥
var Secret = []byte("admin-go-api")
var (
	ErrAbsent  = "token absent"  // 不存在
	ErrInvalid = "token invalid" // 令牌无效
)

// 根据用户信息生成token
func GenerateToken(admin entity.SysAdmin) (string, error) {
	var jwtAdmin = entity.JwtAdmin{
		ID:       admin.ID,
		Username: admin.Username,
		Nickname: admin.Nickname,
		Email:    admin.Email,
		Icon:     admin.Icon,
		Phone:    admin.Phone,
		Note:     admin.Note,
	}
	c := userStdClaims{
		jwtAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpiresDuration).Unix(),
			Issuer:    "admin", // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(Secret)
}

// 解析
func ValidateToken(tokenString string) (*entity.JwtAdmin, error) {
	if tokenString == "" {
		return nil, errors.New(ErrAbsent)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if token == nil {
		return nil, errors.New(ErrInvalid)
	}
	claims := userStdClaims{}
	_, err = jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	return &claims.JwtAdmin, err
}

func GetAdminId(c *gin.Context) (uint, error) {
	u, exist := c.Get(constant.ContextKeyUserObj)
	if !exist {
		return 0, errors.New("not found user")
	}
	admin, ok := u.(*entity.JwtAdmin)
	if ok {
		return admin.ID, nil
	}
	return 0, errors.New("can't convert to id struct")
}

func GetAdminName(c *gin.Context) (string, error) {
	u, exist := c.Get(constant.ContextKeyUserObj)
	if !exist {
		return string(string(0)), errors.New("not found user")
	}
	admin, ok := u.(*entity.JwtAdmin)
	if ok {
		return admin.Username, nil
	}
	return string(string(0)), errors.New("can't convert to api name")
}

func GetAdmin(c *gin.Context) (*entity.JwtAdmin, error) {
	u, exist := c.Get(constant.ContextKeyUserObj)
	if !exist {
		return nil, errors.New("not found user")
	}
	admin, ok := u.(*entity.JwtAdmin)
	if ok {
		return admin, nil
	}
	return nil, errors.New("can't convert to api struct")
}
