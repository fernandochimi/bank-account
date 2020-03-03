package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type OperationType struct {
	gorm.Model
	Description string `json:"description"`
}

func GetOperations() ([]*OperationType) {
	operations := make([]*OperationType, 0)
	err := GetDatabase().Table("operation_types").Find(&operations).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return operations
}

func GetOperation(u uint) (*OperationType) {
	operation := &OperationType{}
	err := GetDatabase().Table("operation_types").Where("id = ?", u).First(&operation).Error
	if err != nil {
		return nil
	}
	return operation
}