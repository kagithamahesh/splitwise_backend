package models

type ExpensesList struct {
	ExpenseId   uint   `gorm:"size:100;not null" json:"expense_id"`
	UserId      string `gorm:"size:100;not null" json:"user_id"`
	ShareAmount string `gorm:"size:100;not null" json:"share_amount"`
}

func (e *ExpensesList) SaveExpenseList() (*ExpensesList, error) {
	var err error
	err = DB.Create(&e).Error
	if err != nil {
		return &ExpensesList{}, err
	}
	return e, nil
}
