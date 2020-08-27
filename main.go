package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	router = gin.Default()
)

func main() {
	router.POST("/login", login)
	log.Fatal(router.Run(":8080"))
}

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var user = User{
	ID:       1,
	Username: "username",
	Password: "password",
}

func createToken(userId int64) (string, error) {
	// Create Access Token
	os.Setenv("ACCESS_SECRET", "gheuireuihewjkbhjewk")
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	if err != nil {
		return "", err
	}
	return token, nil
}

func login(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	if user.Username != u.Username || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Pleae provide a valid login")
		return
	}

	token, err := createToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, token)
}
