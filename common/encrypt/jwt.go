package encrypt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mitchellh/mapstructure"
	"time"
)

const (
	SigningMethodHS256 = iota
	SigningMethodHS384
	SigningMethodHS512
)

type JwtBaseInfo struct {
	Issuer     string
	Subject    string
	NotBefore  time.Time
	ExpiresAt  time.Time
	IssuedAt   time.Time
	Secret     string
	SignMethod int32
}
type Jwt struct {
	Base  JwtBaseInfo
	Extra interface{}
	jwt.RegisteredClaims
}

func NewJwt(base JwtBaseInfo, data interface{}) *Jwt {
	return &Jwt{
		Base:  base,
		Extra: data,
	}
}

func (j *Jwt) Generate() (string, error) {

	j.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(j.Base.ExpiresAt)
	j.RegisteredClaims.Issuer = j.Base.Issuer
	j.RegisteredClaims.Subject = j.Base.Subject
	j.RegisteredClaims.IssuedAt = jwt.NewNumericDate(j.Base.IssuedAt)
	j.RegisteredClaims.NotBefore = jwt.NewNumericDate(j.Base.NotBefore)

	var m *jwt.SigningMethodHMAC
	switch j.Base.SignMethod {
	case SigningMethodHS256:
		m = jwt.SigningMethodHS256
	case SigningMethodHS384:
		m = jwt.SigningMethodHS384
	case SigningMethodHS512:
		m = jwt.SigningMethodHS512
	default:
		m = jwt.SigningMethodHS256
	}
	token := jwt.NewWithClaims(m, j)
	return token.SignedString([]byte(j.Base.Secret))
}

func ParseToken(tokenString string, secret string) (interface{}, bool, error) {

	token, err := jwt.ParseWithClaims(tokenString, &Jwt{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, false, err
	}
	if claims, ok := token.Claims.(*Jwt); ok && token.Valid {
		return claims.Extra, true, nil
	}
	return nil, false, errors.New("invalid token")
}

func GetDataFromToken(token, secret string, target interface{}) error {
	dataFromToken, isValid, err := ParseToken(token, secret)
	if err != nil {
		return err
	}
	if !isValid {
		return errors.New("invalid token")
	}
	return mapstructure.Decode(dataFromToken, &target)
}
