package config

import (
	"backend/dao/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB // 假设使用 gorm 作为 ORM
var PythonScript string
var PythonInterpreter string

func init() {
	PythonScript = "./internal/crnn/ocr.py"
	PythonInterpreter = "D:\\Anaconda\\anaconda3\\envs\\crnn2\\python.exe"
	var err error
	DB, err = SetupDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrate the schema
	if err := model.Migrate(DB); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}

func SetupDatabase() (*gorm.DB, error) {
	dsn := "root:3344255@tcp(localhost:3306)/pmtr?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
