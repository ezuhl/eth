package accounttoken

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"time"
)

const (
	key = "totalPrivateKey"
)

type SystemClaims struct {
	EthTestID string `json:"eth_test_id"`
	jwt.StandardClaims
}

type Token interface {
	MakeToken(userId int64) (string, error)
	ValidateToken(token string) (*SystemClaims, error)
}

type tokenAuth struct {
}

func NewToken() Token {
	t := &tokenAuth{}
	return t
}

func (t *tokenAuth) MakeToken(userId int64) (string, error) {

	claims := SystemClaims{
		fmt.Sprintf("%d", userId),
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Second * 15000).Unix(),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))

}

func (t *tokenAuth) ValidateToken(tokenString string) (*SystemClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &SystemClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "could not parse token with claim")
	}

	if claims, ok := token.Claims.(*SystemClaims); ok && token.Valid {
		return claims, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, errors.Wrap(err, "not a real token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, errors.Wrap(err, "token timed out")
		} else {
			return nil, errors.Wrap(err, "something is wrong with the token")
		}
	}
	return nil, errors.New("could not validate token")
}
