package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/splitwise_backend/models"
	"github.com/gin-gonic/gin"
)

type GroupInput struct {
	GroupName       string `gorm:"column:group_name"`
	CreatedByUserId string `gorm:"column:created_by_user_id"`
}

func CreateGroup(c *gin.Context) {
	user_id := c.Query("id")
	group_name := c.Query("groupName")

	u := models.Groupslist{}
	u.CreatedByUserId = user_id
	u.GroupName = group_name
	_, err := u.SaveGroup()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Group Success"})

}

func GetGroupByuserId(c *gin.Context) {
	//var input GroupInput
	user_id := c.Query("id")
	for k, v := range c.Request.Header {
		fmt.Printf("%s: %s\n", k, v)
	}
	userIDUint64, err := strconv.ParseUint(user_id, 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"userId": "userId missing or invalid"})
		return
	}
	groupsList, err := models.GetGroupsCreatedByUserID(uint(userIDUint64))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"groups": groupsList})

}
