package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
	"github.com/joho/godotenv"
	"fmt"
)

var db *gorm.DB


func init() {
	e := godotenv.Load("github.com/fernandochimi/bank-account/config/api.env")

	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	dbUrl := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		dbHost, dbPort, username, dbName, password)

	fmt.Println(dbUrl)

	connection, err := gorm.Open("postgres", dbUrl)

	if err != nil {
		fmt.Print(e)
	}

	db = connection
	db.Debug().AutoMigrate(&Account{}, &OperationType{}, &Transaction{})

	operation_1 := OperationType{Description: "COMPRA A VISTA"}
	operation_2 := OperationType{Description: "COMPRA PARCELADA"}
	operation_3 := OperationType{Description: "SAQUE"}
	operation_4 := OperationType{Description: "PAGAMENTO"}

	op := db.Where("description IN (?)", []string{"COMPRA A VISTA", "COMPRA PARCELADA", "SAQUE", "PAGAMENTO"}).Find(&OperationType{})

	if err := op.Error; err != nil {
		db.Create(&operation_1)
		db.Create(&operation_2)
		db.Create(&operation_3)
		db.Create(&operation_4)
	} else {
		fmt.Print("Operations Already Exists")
	}
}

func GetDatabase() *gorm.DB {
	return db
}