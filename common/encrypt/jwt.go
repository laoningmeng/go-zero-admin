package encrypt

import "github.com/golang-jwt/jwt/v4"

type Jwt struct {
	Field map[string]string
	jwt.Claims
}

func (j *Jwt) Create(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["user_id"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
