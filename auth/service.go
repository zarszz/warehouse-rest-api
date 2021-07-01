package auth

import (
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

type Service interface {
	GenerateToken(userId string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type JwtService struct {
	service Service
}

var signedKey = []byte(viper.GetString("secret"))

func NewService() *JwtService {
	return &JwtService{}
}


func (s *JwtService) GenerateToken(userId string) (string, error){
	claim := jwt.MapClaims{}
	claim["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(signedKey)
	if err != nil {
		return signedToken, err
	}
	return signedToken, err
}

func (s *JwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("INVALID TOKEN")
		}
		return []byte(signedKey), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
