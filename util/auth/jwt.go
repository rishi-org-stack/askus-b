package auth

import (
	"askUs/v1/util/config"
	"fmt"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Auth struct {
	key  []byte
	algo jwt.SigningMethod
	// Duration for which the jwt token is valid.
	ttl time.Duration
}

func Init(env *config.Env) (*Auth, error) {
	sgMethod := jwt.GetSigningMethod(env.Algo)
	if sgMethod == nil {
		return &Auth{}, fmt.Errorf("util/auth.go -line20: unable to genrate signing method chec algo in env file")
	}
	return &Auth{
		key:  []byte(env.Key),
		algo: sgMethod,
		ttl:  time.Duration(env.JWTDurtaion) * time.Minute,
	}, nil
}
func (s *Auth) ParseToken(authHeader string) (*jwt.Token, error) {
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return nil, fmt.Errorf("util/auth.go -line33: header passed should contains Bearer and ayth token in it")
	}

	return jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
		if s.algo != token.Method {
			return nil, fmt.Errorf("util/auth.go -line38: method of header doen't matches the config")
		}
		return s.key, nil
	})

}
func (a *Auth) GenrateToken(id int, email string, typeOfClinet string) (string, error) {
	return jwt.NewWithClaims(a.algo, jwt.MapClaims{
		"id":     id,
		"email":  email,
		"exp":    time.Now().Add(a.ttl).Unix(),
		"client": typeOfClinet,
	}).SignedString(a.key)
}
