package middlewares

import (
	"log"
	"net/http"
)

func generateToken(userid int) string {
	jwtWrapper := JwtWrapper{
		SecretKey: getSecretKey(),
		Issuer: "AuthService",
		ExpirationHours: 24,
	}

	generatedToken, err := jwtWrapper.GenerateToken(userid)

	if err != nil {
		log.Println(err)
	}

	return generatedToken
}

func JWTVerify(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jwtWrapper := JwtWrapper{
			SecretKey: getSecretKey(),
			Issuer: "AuthService",
		}

		claims, err := jwtWrapper.ValidateToken(r.Header.Get("Authorization"))

		if err != nil {
			log.Println("jwt not ok")
			log.Println(err)
			log.Println(claims)
		} else {
			log.Println("jwt ok")
			log.Println(err)
			log.Println(claims)
		}

		next(w, r)
	}
}