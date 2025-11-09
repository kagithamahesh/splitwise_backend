package models

type Expenses struct {
	ExpensesId   uint   `gorm:"primarykey;column:expense_id" json:"expense_id"`
	GroupId      string `gorm:"size:100;not null" json:"group_id"`
	Description  string `gorm:"size:100;not null" json:"description"`
	Amount       string `gorm:"size:100;not null" json:"amount"`
	PaidByUserId string `gorm:"size:100;not null" json:"paid_by_user_id"`
	Date         string `gorm:"size:100;not null" json:"date"`
	SplitMethod  string `gorm:"primarykey;column:split_method" json:"split_method"`

	// `gorm:"size:100;not null" json:"group_id"`

}

func (e *Expenses) SaveExpense() (*Expenses, error) {
	var err error
	err = DB.Create(&e).Error
	if err != nil {
		return &Expenses{}, err
	}
	return e, nil
}
