package encrypt

import (
	"fmt"
	"testing"
	"time"
)

func TestJwt(t *testing.T) {
	type User struct {
		Username string
	}
	jwt := NewJwt(JwtBaseInfo{
		ExpiresAt: time.Now().Add(3 * time.Hour),
		IssuedAt:  time.Now(),
		Secret:    "aaaaa",
	}, User{
		Username: "notojbk",
	})
	token, err := jwt.Generate()
	if err != nil {
		t.Error(err)
	}
	data, b, err := ParseToken(token, jwt.Base.Secret)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(data, b, err)
	d := data.(map[string]interface{})
	fmt.Println(d["Username"])
}
