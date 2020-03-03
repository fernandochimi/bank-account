package models

import (
	"fmt"
	"time"
	"github.com/jinzhu/gorm"
	u "github.com/fernandochimi/bank-account/src/utils"
)


type Transaction struct {
	gorm.Model
	AccountId uint `json:accountId,gorm:"foreignkey:ID"`
	OperationId uint `json:operationId,gorm:"foreignkey:ID"`
	Amount string `json:"amount"`
	EventDate time.Time `gorm"default:CURRENT_TIMESTAMP"`
}

func (transaction *Transaction) Create() (map[string] interface{}) {

	// am := transaction.Amount

	opErr := GetDatabase().Table("operation_types").Where("id IN (?)").First(&OperationType{}).Error
	if opErr != nil {
		return u.Message("Operation Type does not exist")
	} else {
		fmt.Println(opErr)
	}
	// if transaction.OperationId in op {
	// 	transaction.Amount = -am 
	// }

	GetDatabase().Create(transaction)

	if transaction.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}

	response := u.Message(true, "Transaction Success")
	response["account_id"] = transaction.AccountId
	response["operation_type_id"] = transaction.OperationId
	response["amount"] = transaction.Amount
	return response
}
