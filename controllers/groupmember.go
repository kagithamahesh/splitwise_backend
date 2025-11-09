package controllers

import (
	"net/http"

	"example.com/splitwise_backend/models"
	"github.com/gin-gonic/gin"
)

type GroupMemberInfo struct {
	GroupId string `gorm:"column:group_id"`
	UserId  string `gorm:"column:user_id"`
}

func CreateGroupMember(c *gin.Context) {

	user_id := c.Query("id")
	groupId := c.Query("groupId")
	var input GroupMemberInfo

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := models.GroupMembers{}

	u.UserId = user_id
	u.GroupId = groupId
	_, err := u.SaveGroupMember()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Group Member Success"})

}
