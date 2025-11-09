package controllers

import (
	"net/http"
	"strconv"

	"example.com/splitwise_backend/models"
	"github.com/gin-gonic/gin"
)

type ExpensesListInput struct {
	ExpenseId   uint   `gorm:"size:100;not null" json:"expense_id"`
	UserId      string `gorm:"size:100;not null" json:"user_id"`
	ShareAmount string `gorm:"size:100;not null" json:"share_amount"`
}

func ExpenseListCreate(c *gin.Context) {
	user_id := c.Query("id")
	expense_id := c.Query("expenseId")
	share_amount := c.Query("shareAmount")

	var input ExpensesListInput
	expenseIDUint64, err := strconv.ParseUint(expense_id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expense ID format."})
		return
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := models.ExpensesList{}

	u.ShareAmount = share_amount
	u.UserId = user_id
	u.ExpenseId = uint(expenseIDUint64)
	_, err = u.SaveExpenseList()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Expenses List Success"})
}
