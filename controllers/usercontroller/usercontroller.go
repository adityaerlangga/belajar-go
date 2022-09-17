package usercontroller

import (
	"net/http"

	"github.com/adityaerlangga/golang-auth/usecase/usercase"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	users := usercase.GetUsers(c)
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func Show(c *gin.Context) {
	user, err := usercase.GetUserById(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}

func Register(c *gin.Context) {
	user, err := usercase.RegisterUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data gagal ditambahkan", "error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Anda berhasil terdaftar...", "user": user})
	}
}

func Login(c *gin.Context) {
	user, err := usercase.LoginUser(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Login berhasil", "user": user})
	}
}

func ResetPassword(c *gin.Context) {

}
