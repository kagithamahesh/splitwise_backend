package controllers

import (
	"net/http"

	"example.com/splitwise_backend/models"
	"github.com/gin-gonic/gin"
)

type ExpenseInput struct {
	ExpensesId   uint   `gorm:"primarykey;column:expense_id" json:"expense_id"`
	GroupId      string `gorm:"size:100;not null" json:"group_id"`
	Description  string `gorm:"size:100;not null" json:"description"`
	Amount       string `gorm:"size:100;not null" json:"amount"`
	PaidByUserId string `gorm:"size:100;not null" json:"paid_by_user_id"`
	Date         string `gorm:"size:100;not null" json:"date"`
	SplitMethod  string `gorm:"primarykey;column:split_method" json:"split_method"`
}

func ExpenseCreate(c *gin.Context) {
	user_id := c.Query("id")
	group_id := c.Query("groupId")
	description := c.Query("description")
	amount := c.Query("amount")
	date := c.Query("date")
	split_method := c.Query("split_method")

	var input ExpenseInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.Expenses{}

	u.Amount = amount
	u.Date = date
	u.SplitMethod = split_method
	u.Description = description
	u.GroupId = group_id
	u.PaidByUserId = user_id
	_, err := u.SaveExpense()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Expenses Success"})
}
