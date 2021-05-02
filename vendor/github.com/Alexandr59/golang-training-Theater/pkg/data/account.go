package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type AccountData struct {
	db *gorm.DB
}

func NewAccountData(db *gorm.DB) *AccountData {
	return &AccountData{db: db}
}

type Account struct {
	Id          int    `gorm:"primaryKey"`
	FirstName   string `gorm:"first_name"`
	LastName    string `gorm:"last_name"`
	PhoneNumber string `gorm:"phone_number"`
	Email       string `gorm:"email"`
}

func (a AccountData) AddAccount(account Account) (int, error) {
	result := a.db.Create(&account)
	if result.Error != nil {
		return -1, fmt.Errorf("can't insert account to database, error: %w", result.Error)
	}
	return account.Id, nil
}

func (a AccountData) DeleteAccount(entry Account) error {
	result := a.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Account to database, error: %w", result.Error)
	}
	return nil
}

func (a AccountData) UpdateAccount(account Account) error {
	result := a.db.Model(&account).Updates(account)
	if result.Error != nil {
		return fmt.Errorf("can't update account to database, error: %w", result.Error)
	}
	return nil
}

func (a AccountData) FindByIdAccount(entry Account) (Account, error) {
	result := a.db.First(&entry)
	if result.Error != nil {
		return Account{}, fmt.Errorf("can't find Account to database, error: %w", result.Error)
	}
	return entry, nil
}
