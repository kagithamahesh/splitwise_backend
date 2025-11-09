package models

import (
	"errors"
	"fmt"
	"html"
	"strings"

	"example.com/splitwise_backend/utils/token"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Userlist struct {
	UserID   uint   `gorm:"primarykey;column:user_id" json:"user_id"`
	UserName string `gorm:"size:100;not null" json:"user_name"`
	Password string `gorm:"size:100;not null" json:"password"`
	Email    string `gorm:"size:100;not null" json:"email"`
	// CreatedAt time.Time `json:"created_at"`
}

func GetUserByID(uid uint) (Userlist, error) {
	var u Userlist
	if err := DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}
	// Set Password to empty string to prevent it from being sent in the API response
	u.Password = ""
	return u, nil
}

// PrepareGive is only needed if you want to use it within other model methods.
// For API responses, setting u.Password = "" directly in the controller or GetUserByID is clearer.
func (u *Userlist) PrepareGive() {
	u.Password = ""
}

func VerifyPassword(password, hashedPassword string) error {

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string, error) {
	u := Userlist{}

	// FIX: Ensure we specifically check if the record was not found (gorm.ErrRecordNotFound)
	err := DB.Model(Userlist{}).Where("user_name = ?", username).Take(&u).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", errors.New("User not found")
	}
	if err != nil {
		return "", err // Handle other DB errors
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return "", errors.New("Incorrect password") // Return a generic error for security
	}

	token, err := token.GenerateToken(u.UserID)
	fmt.Println(token)
	if err != nil {
		return "", err
	}

	return token, nil
}
func (u *Userlist) SaveUser() (*Userlist, error) {
	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &Userlist{}, err
	}
	return u, nil
}
func (u *Userlist) BeforeSave(tx *gorm.DB) error {
	// Turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	// Remove spaces in username
	u.UserName = html.EscapeString(strings.TrimSpace(u.UserName))

	return nil
}
