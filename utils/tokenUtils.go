package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var tokenSecret string
var tokenExpirationTime int64

//TokenChecker check the token returning the token claims
func TokenChecker(token string) (interface{}, error) {
	tokenParsed, parseError := jwt.Parse(token, func(tk *jwt.Token) (interface{}, error) {
		if _, success := tk.Method.(*jwt.SigningMethodHMAC); !success {
			return nil, fmt.Errorf("Token error: invalid header algorithm")
		}
		return []byte(getTokenSecret()), nil
	})
	if parseError != nil {
		return nil, parseError
	}
	if claims, ok := tokenParsed.Claims.(jwt.MapClaims); ok {
		return claims["data"], parseError
	}
	return nil, parseError
}

//EncodeToken generate a token with claims and assigned and expiration dates
func EncodeToken(payload interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": payload,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Second * time.Duration(getTokenExpirationTime())).Unix(),
	})
	return token.SignedString([]byte(getTokenSecret()))
}

func getTokenExpirationTime() int64 {
	if tokenExpirationTime == 0 {
		tkExpTime := os.Getenv("TOKEN_EXPIRATION_TIME")
		if tkExpTime == "" {
			tkExpTime = "360"
		}
		var strConvError error
		tokenExpirationTime, strConvError = strconv.ParseInt(tkExpTime, 10, 64)
		if strConvError != nil {
			tokenExpirationTime = 360
		}
	}
	return tokenExpirationTime
}

func getTokenSecret() string {
	if tokenSecret == "" {
		tokenSecret = os.Getenv("TOKEN_SECRET")
		if tokenSecret == "" {
			tokenSecret = "secret"
		}
	}
	return tokenSecret
}
