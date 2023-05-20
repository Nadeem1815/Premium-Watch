package middleware

import (
	"fmt"

	"time"

	"github.com/golang-jwt/jwt"
)

func ValidateToken(cookie string) (string, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("secret"), nil
	})
	if err != nil {
		return "", err
	}
	// if token == nil || !token.valid {
	// 	return 0, fmt.Errorf("invalid token")
	// }
	var parsedID interface{}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		parsedID = claims["id"]
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return "", fmt.Errorf("token expired")
		}
	}
	// Type Assertion
	value, ok := parsedID.(string)
	if !ok {
		return "", fmt.Errorf("expected an int value ,but got %T", parsedID)

	}

	return value, nil
}
