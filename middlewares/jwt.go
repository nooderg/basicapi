package middlewares

import (
	"log"
	"net/http"
	"basic-api/utils"
)

func JWTVerify(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jwtWrapper := utils.JwtWrapper{
			SecretKey: utils.GetJWTSecretKey(),
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