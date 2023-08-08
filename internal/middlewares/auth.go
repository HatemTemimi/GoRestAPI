package middleware

import (
	"apigo/internal/user/service"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AuthMiddleware struct {
	userService service.UserService
}

func MakeAuthMiddleware(service service.UserService) AuthMiddleware {
	return AuthMiddleware{userService: service}
}

func (m *AuthMiddleware) RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"info": "you're not logged in!",
		})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "could not parse auth token",
		})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//auth success
		//check if the cookie is expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "cookie expired",
			})
			return

		}
		//notorize user with id on cookie
		user, err := m.userService.FindByID(uint(claims["sub"].(float64)))
		if err != nil || user.ID == 0 {

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "cookie does not belong to user",
			})
			return
		}
		//continue
		c.Set("user", user)
		c.Next()
	} else {
		//abort
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token is invalid",
		})
		return
	}
}
