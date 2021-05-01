package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type RoleData struct {
	db *gorm.DB
}

func NewRoleData(db *gorm.DB) *RoleData {
	return &RoleData{db: db}
}

type Role struct {
	Id   int    `gorm:"primaryKey"`
	Name string `gorm:"name"`
}

func (r RoleData) AddRole(role Role) (int, error) {
	result := r.db.Create(&role)
	if result.Error != nil {
		return -1, fmt.Errorf("can't insert Role to database, error: %w", result.Error)
	}
	return role.Id, nil
}

func (r RoleData) DeleteRole(entry Role) error {
	result := r.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Role to database, error: %w", result.Error)
	}
	return nil
}

func (r RoleData) UpdateRole(role Role) error {
	result := r.db.Model(&role).Updates(role)
	if result.Error != nil {
		return fmt.Errorf("can't update Role to database, error: %w", result.Error)
	}
	return nil
}

func (r RoleData) FindByIdRole(entry Role) (Role, error) {
	result := r.db.First(&entry)
	if result.Error != nil {
		return Role{}, fmt.Errorf("can't find Role to database, error: %w", result.Error)
	}
	return entry, nil
}
