package app

//根据特定的鉴权场景对jwt-go库进行设计,组合其提供的 API
import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenOptions struct {
	AccessSecret string
	AccessExpire int64 `json:"access_expire"`
	UserID       int64 `json:"user_id"`
	RoleLevel    int64 `json:"role_level"`
}

// GenerateToken 生成 JWT Token
func GenerateToken(tokenOptions TokenOptions) (string, error) {
	//claims是jwt的Payload，存储有效字段数据
	claims := make(jwt.MapClaims)
	claims["user_id"] = tokenOptions.UserID
	claims["role_level"] = tokenOptions.RoleLevel
	claims["access_expire"] = time.Now().Unix() + tokenOptions.AccessExpire

	//jwt.NewWithClaims根据 Claims 结构体创建 Token 实例
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//SignedString根据所传入 Secret 不同，进行签名并返回标准的 Token。
	token, err := tokenClaims.SignedString([]byte(tokenOptions.AccessSecret))
	return token, err
}

//// ParseToken 解析和校验 Token，返回 *Claims有效字段负载
//func ParseToken(token, secret string) (*Claims, error) {
//	//解析鉴权的声明，方法内部是具体的解码和校验过程，返回 *Token。
//	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
//		return GetJWTSecret(secret), nil
//	})
//	if err != nil {
//		return nil, err
//	}
//	//拿到token之后，验证是否有效
//	if tokenClaims != nil {
//		claims, ok := tokenClaims.Claims.(*Claims)
//		//验证基于时间的声明
//		if ok && tokenClaims.Valid {
//			return claims, nil
//		}
//	}
//	return nil, err
//}
