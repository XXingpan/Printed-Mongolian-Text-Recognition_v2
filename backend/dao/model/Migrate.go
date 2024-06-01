package model

import "gorm.io/gorm"

// Migration to set up database schema
func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&User{}, &Image{}, &RecognitionResult{})
	if err != nil {
		return err
	}
	return nil
}
