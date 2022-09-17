package usercase

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/adityaerlangga/golang-auth/app/config"
	"github.com/adityaerlangga/golang-auth/entities/userentity"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *gin.Context) (users []userentity.User) {
	config.DB.Find(&users)
	return
}

func GetUserById(c *gin.Context) (user userentity.User, err error) {
	id := c.Param("id")
	if err := config.DB.First(&user, id).Error; err != nil {
		return user, err
	}
	return
}

func RegisterUser(c *gin.Context) (user userentity.User, err error) {

	fmt.Println(c)
	if err = c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	password := string(user.Password)
	hashedPassword, _ := HashPassword(password)
	user.Password = hashedPassword

	config.DB.Create(&user)
	return user, err
}

func LoginUser(c *gin.Context) (user userentity.User, err error) {
	if err = c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	passInput := user.Password // user input (passInput)
	email := user.Email
	// Find By Email
	config.DB.Where("email = ?", email).Find(&user) // Kekurangan Fungsi Find tidak return error type yang bener.
	if user.ID == 0 {
		err = errors.New("Data tidak ditemukan...")
		return user, err
	}
	// Checking Password Input & Password DB
	checkPass := CheckPasswordHash(passInput, user.Password)
	if !checkPass {
		err = errors.New("Data tidak ditemukan...")
		return user, err
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Berhasil login..."})
	return user, err
}

func ResetPassword(c *gin.Context) (user userentity.User) {

	return
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
