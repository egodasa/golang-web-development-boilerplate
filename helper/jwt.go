// Packacge middleware untuk menampung fungsi yang berhubungan dengan middleware
package helper

import (
	fmt "fmt"
  
	jwt "github.com/dgrijalva/jwt-go"
)

// GenJwt
// Untuk generate token dengan jwt
// @data : Data yang ingin dimasukan ke token, berupa map[string]interface{}
// @kunci : Kunci untuk generate token
func GenerateJwtToken(data map[string]interface{}, kunci string) string {
	var sign *jwt.Token = jwt.New(jwt.GetSigningMethod("HS256"))
	var claims map[string]interface{} = sign.Claims.(jwt.MapClaims)

	for key, value := range data {
		claims[key] = value
	}
	var token, _ = sign.SignedString([]byte(kunci))
	return token
}

func ValidateJwtToken(tokenString string, kunci string) bool {
	var token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Token not valid!")
		}
		return []byte(kunci), nil
	})

	if token != nil && err == nil {
		return true
	} else {
		return false
	}

}
