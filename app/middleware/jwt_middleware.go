package middleware

import (
	"fmt"

	"agnos-assignment/app/constant"
	"agnos-assignment/app/pkg"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer pkg.PanicHandler(c)

		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			pkg.PanicException(constant.Unauthorized)
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := pkg.JWTServiceInit().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)

			c.Set("JWTService", pkg.JWTServiceInit())
			c.Set("JWTClaims", claims)

			jwtSvc := pkg.JWTServiceInit()
			payload := jwtSvc.GetPayloadInToken(c)

			c.Set("UserContext", *payload)

		} else {
			fmt.Println("testing")
			fmt.Println(err)
			pkg.PanicException(constant.Unauthorized)
		}

	}
}
