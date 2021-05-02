package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type UserData struct {
	db *gorm.DB
}

func NewUserData(db *gorm.DB) *UserData {
	return &UserData{db: db}
}

type User struct {
	Id          int    `gorm:"primaryKey"`
	AccountId   int    `gorm:"account_id"`
	FirstName   string `gorm:"first_name"`
	LastName    string `gorm:"last_name"`
	RoleId      int    `gorm:"role_id"`
	LocationId  int    `gorm:"location_id"`
	PhoneNumber string `gorm:"phone_number"`
}

type SelectUser struct {
	Id                  int
	FirstName           string
	LastName            string
	Role                string
	LocationAddress     string
	LocationPhoneNumber string
	PhoneNumber         string
}

func (u UserData) AddUser(user User) (int, error) {
	result := u.db.Create(&user)
	if result.Error != nil {
		return -1, fmt.Errorf("can't insert User to database, error: %w", result.Error)
	}
	return user.Id, nil
}

func (u UserData) DeleteUser(entry User) error {
	result := u.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete User to database, error: %w", result.Error)
	}
	return nil
}

func (u UserData) UpdateUser(user User) error {
	result := u.db.Model(&user).Updates(user)
	if result.Error != nil {
		return fmt.Errorf("can't update User to database, error: %w", result.Error)
	}
	return nil
}

func (u UserData) FindByIdUser(entry User) (User, error) {
	result := u.db.First(&entry)
	if result.Error != nil {
		return User{}, fmt.Errorf("can't find User to database, error: %w", result.Error)
	}
	return entry, nil
}

func (u UserData) ReadAllUsers(account Account) ([]SelectUser, error) {
	var users []SelectUser
	rows, err := u.db.Table("users").Select("users.id, users.first_name, "+
		"users.last_name, roles.name, locations.address, locations.phone_number, users.phone_number").
		Joins("JOIN roles on users.role_id = roles.id").
		Joins("JOIN locations on locations.id = users.account_id").
		Where("users.account_id = ?", account.Id).
		Rows()
	if err != nil {
		return nil, fmt.Errorf("can't get users from database, error:%w", err)
	}
	for rows.Next() {
		var temp SelectUser
		err = rows.Scan(&temp.Id, &temp.FirstName, &temp.LastName, &temp.Role,
			&temp.LocationAddress, &temp.LocationPhoneNumber, &temp.PhoneNumber)
		if err != nil {
			return nil, fmt.Errorf("can't scan users from database, error:%w", err)
		}
		users = append(users, temp)
	}
	return users, nil
}
