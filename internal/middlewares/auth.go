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
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		}
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//auth success
		//check if the cookie is expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		user, err := m.userService.FindByID(uint(claims["sub"].(uint)))
		if err != nil || user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		fmt.Println("logged user", user)
		c.Next()

	} else {
		//abort
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
