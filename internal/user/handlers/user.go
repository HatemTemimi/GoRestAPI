package handlers

import (
	"apigo/internal/user/models"
	"apigo/internal/user/repository"
	"apigo/internal/user/service"
	"net/http"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserApi struct {
	UserService service.UserService
}

func MakeUserApi(DB *gorm.DB) UserApi {
	return UserApi{
		UserService: service.MakeUserService(repository.MakeUserRepository(DB)),
	}
}

func ProvideUserApi(service service.UserService) UserApi {
	return UserApi{UserService: service}
}

var req struct {
	Email    string
	Password string
}

func (u *UserApi) Signup(c *gin.Context) {
	//check request
	if c.Bind(&req) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
	}

	//if ok create user
	usr := models.User{Email: req.Email, Password: string(hash)}
	user, err := u.UserService.Create(usr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})

}

func (u *UserApi) Login(c *gin.Context) {
	//check request
	if c.Bind(&req) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
	}

	user, err := u.UserService.FindByEmail(req.Email)
	if err != nil {
		if c.Bind(&req) != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err,
			})
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "could not create auth token",
		})
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.Set("user", user)
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})

}

func (m *UserApi) Validate(c *gin.Context) {
	token, err := c.Cookie("Authorization")
	if err != nil || token == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": "Not logged in",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "logged in",
			"token":  token,
		})
	}
}
