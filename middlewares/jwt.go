package middlewares

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
)

func myLookupKey(token []byte) ([]byte, error) {
	if (true) {
		return token, nil
	} else {
		return []byte{}, nil
	}
}

func JWTVerify(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenToTest := r.Header.Get("jwt")

		token, err := jwt.Parse(tokenToTest, func(token *jwt.Token) ([]byte, error) {
			// logique pour vérifier
			return myLookupKey(token.Header["kid"])
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

		// Set some claims
		token.Claims["username"] = "bar"
		token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		// Sign and get the complete encoded token as a string
		tokenString, err := token.SignedString(mySigningKey)

		next(w, r)
	}
}