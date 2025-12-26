package controllers

import (
	"encoding/base64"
	"fmt"
	"io"
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
	user_id := c.Request.FormValue("id")
	group_name := c.Request.FormValue("groupName")
	description := c.Request.FormValue("description")
	file, _, err := c.Request.FormFile("image")
	var imageBytes []byte
	if err == nil {
		imageBytes, err = io.ReadAll(file)
		fmt.Println(imageBytes, "imageBytes")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image stream"})
			return
		}
		file.Close()
	} else {
		// If there's an error, it means no image was sent or the field name "image" didn't match
		fmt.Println("No image uploaded or error encountered:", err)
	}
	u := models.Groupslist{}
	u.CreatedByUserId = user_id
	u.GroupName = group_name
	u.Description = description
	u.ImageData = imageBytes

	_, err = u.SaveGroup()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Group created with byte data"})

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
	for i := range groupsList {
		if len(groupsList[i].ImageData) > 0 {
			groupsList[i].ImageDataBase64 =
				base64.StdEncoding.EncodeToString(groupsList[i].ImageData)
		}
	}
	c.JSON(http.StatusOK, gin.H{"groups": groupsList})

}
