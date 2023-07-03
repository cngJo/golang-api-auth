package models

import (
	"github.com/cngJo/golang-api-auth/internal/binary_uuid"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	ID binary_uuid.BinaryUUID `gorm:"primary_key;type:uuid" json:"id"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	uuid, err := uuid.NewRandom()
	base.ID = binary_uuid.BinaryUUID(uuid)

	return err
}

type User struct {
	BaseModel
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

type RefreshToken struct {
	BaseModel
	User        User
	Fingerprint string
	Token       string
}
