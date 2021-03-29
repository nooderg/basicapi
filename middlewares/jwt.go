package middlewares

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var signingKey = []byte("this is a key")

func myLookupKey(token []byte) ([]byte, error) {
	if true {
		return token, nil
	} else {
		return []byte{}, nil
	}
}

func JWTVerify(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenToTest := r.Header.Get("jwt")

		token, err := jwt.Parse(tokenToTest, func(token *jwt.Token) (interface{}, error) {
			// logique pour vérifier
			// return myLookupKey(token.Header["kid"])
			log.Println(token)
			return nil, nil
		})

		if err == nil && token.Valid {
			// token is ok
		} else {
			// token is not ok
		}

		next(w, r)
	}
}

// TODO: transform to a normal function, not a middleware
func JWTCreate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Create the token
		token := jwt.New(jwt.GetSigningMethod("HS256"))

		claims := make(jwt.MapClaims)
		claims["username"] = "username test"
		claims["date"] = time.Now().Unix()
		token.Claims = claims

		tokenString, err := token.SignedString(mySigningKey)

		next(w, r)
	}
}
