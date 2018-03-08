package utils

import(
	"os"
	"strconv"
	"testing"
)

//TestGetTokenSecret check if the token's secret key always has a valid value
func TestGetTokenSecret(t *testing.T) {
	t.Run("EnviromentVarNotSeted", func(t *testing.T) {
		tokenSecret = ""
		tkSecret := getTokenSecret()
		if tkSecret == "" {
			t.Error("Missing token secret")
		}
		if tokenSecret == "" {
			t.Error("Token secret was not seted")
		}
	})
	t.Run("EnviromentVarSeted", func(t *testing.T) {
		tokenSecret = ""
		mySecret := "mySecRet"
		setEnvError := os.Setenv("TOKEN_SECRET", mySecret)
		if setEnvError != nil {
			t.Error("Fail to set the enviroment variable TOKEN_SECRET: " + setEnvError.Error())
		}
		tkSecret := getTokenSecret()
		if tkSecret != mySecret {
			t.Error("Token secret was not correctly seted")
		}
	})
	t.Run("SecondEnviromentVarSeted", func(t *testing.T) {
		tkSecret := getTokenSecret()
		if tkSecret == "" {
			t.Error("Token secret was not correctly seted in the previous function")
		}
	})
}

//TestGetTokenExpirationTime check if the token's expiration time always has a valid value
func TestGetTokenExpirationTime(t *testing.T) {
	t.Run("EnviromentVarNotSeted", func(t *testing.T) {
		tokenExpirationTime = 0
		tkExpirationTime := getTokenExpirationTime()
		if tokenExpirationTime == 0 || tokenExpirationTime != tkExpirationTime {
			t.Error("Missing token expiration time")
		}
	})
	t.Run("EnviromentVarSeted", func(t *testing.T) {
		tokenExpirationTime = 0
		myTime := int64(360)
		setEnvError := os.Setenv("TOKEN_EXPIRATION_TIME", strconv.FormatInt(myTime, 10))
		if setEnvError != nil {
			t.Error("Fail to set the enviroment variable TOKEN_EXPIRATION_TIME: " + setEnvError.Error())
		}
		tokenExpirationTime := getTokenExpirationTime()
		if tokenExpirationTime == 0 {
			t.Error("Token expiration time was not correctly seted")
		}
		if tokenExpirationTime != myTime {
			t.Error("Token expiration time is not using the enviroment variable TOKEN_EXPIRATION_TIME")
		}
	})
}

//TestTokenChecker check if the func 'TokenChecker' is handle correctly a token string
/*	Valid token string mock
		Header (algorithm and type)
		{
			"alg": "HS256",
			"typ": "JWT"
		}
		Payload (data)
		{
			"data": {
				"sub": "1234567890",
				"name": "John Doe",
				"iat": 1516239022
			}
		}
		Verify Signature
			HMACSHA256(base64UrlEncode(header) + "." + base64UrlEncode(payload), "secret")
*/
func TestTokenChecker(t *testing.T) {
	tokenMock := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InN1YiI6IjEyMzQ1Njc4OTAiLCJuYW1lIjoiSm9obiBEb2UiLCJpYXQiOjE1MTYyMzkwMjJ9fQ.YdUPFvca3f-HP4w0M8r80Vtr8rJH8TrAqcdbTDxk2io"
	t.Run("ValidToken", func(t *testing.T) {
		tokenSecret = "secret"
		tokenPayload, tokenCheckerError := TokenChecker(tokenMock)
		if tokenCheckerError != nil {
			t.Error("Can not verify the token mock: " + tokenCheckerError.Error())
		}
		if tokenPayload == nil {
			t.Error("Missing data")
		}
		tkPayloadMap, ok := tokenPayload.(map[string]interface{})
		if !ok {
			t.Error("Payload has an invalid format")
		}
		if tkPayloadMap["sub"] == nil || tkPayloadMap["name"] == nil || tkPayloadMap["iat"] == nil {
			t.Error("Missing params on payload")
		}
		subValue, subOk := tkPayloadMap["sub"].(string)
		if !subOk || subValue != "1234567890" {
			t.Error("Sub param corrupted")
		}
		nameValue, nameOk := tkPayloadMap["name"].(string)
		if !nameOk || nameValue != "John Doe" {
			t.Error("Name param corrupted")
		}
		iatValue, iatOk := tkPayloadMap["iat"].(float64)
		if !iatOk || iatValue != 1516239022 {
			t.Error("Iat param corrupted")
		}
	})
	t.Run("InvalidToken", func(t *testing.T) {
		tokenSecret = "secret2"
		_, tokenCheckerError := TokenChecker(tokenMock)
		if tokenCheckerError == nil {
			t.Error("Using a wrong token's secret should return an error")
		}
	})
}

//TestEncodeToken should check the methods that assign/verify the token
func TestEncodeToken(t *testing.T) {
	var tokenString string
	dataMock := map[string]interface{} {
		"sub": "1234567890",
		"name": "John Doe",
	}
	t.Run("AssignToken", func(t *testing.T) {
		tokenSecret = "secret"
		tkString, encodeTokenError := EncodeToken(dataMock)
		if encodeTokenError != nil {
			t.Error("Can not assign the token: " + encodeTokenError.Error())
		}
		tokenString = tkString
	})
	t.Run("VerifyToken", func(t *testing.T) {
		tokenPayload, tokenCheckerError := TokenChecker(tokenString)
		if tokenCheckerError != nil {
			t.Error("Can not verify the token string: " + tokenCheckerError.Error())
		}
		if tokenPayload == nil {
			t.Error("Missing data")
		}
		tkPayloadMap, ok := tokenPayload.(map[string]interface{})
		if !ok {
			t.Error("Token payload invalid format")
		}
		for key, value := range tkPayloadMap {
			if value != dataMock[key] {
				t.Error("The token payload not match the mocked data")
			}
		}
	})
	t.Run("VerifyTokenWrongSecret", func(t *testing.T) {
		tokenSecret = "wrongSecret"
		_, tokenCheckerError := TokenChecker(tokenString)
		if tokenCheckerError == nil {
			t.Error("Using a wrong token's secret should return an error")
		}
	})
}