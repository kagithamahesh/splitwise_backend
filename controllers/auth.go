package controllers

import (
	"fmt"
	"net/http"

	"example.com/splitwise_backend/models"
	"example.com/splitwise_backend/utils/token"
	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Username string `gorm:"column:username"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email"`
}

func Register(c *gin.Context) {

	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := models.Userlist{}
	u.UserName = input.Username
	u.Password = input.Password
	u.Email = input.Email

	_, err := u.SaveUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "registration success"})

}

type LoginInput struct {
	Username string `gorm:"column:user_name"`
	Password string `gorm:"column:password"`
	Email    string `json:"email"`
}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	u := models.Userlist{}

	u.UserName = input.Username
	u.Password = input.Password
	u.Email = input.Email

	token, err := models.LoginCheck(u.UserName, u.Password)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "status": "Success", "profile": gin.H{
		"name":  u.UserName,
		"email": u.Email,
	}})

}

func CurrentUser(c *gin.Context) {

	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}
