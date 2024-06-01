package model

import "time"

// Image represents the image table in the database.
type Image struct {
	ImageID    int       `gorm:"primaryKey;autoIncrement" json:"imageID"` // Unique identifier for the image
	UserID     int       `json:"userID"`                                  // Identifier for the user who owns the image
	ImageURL   string    `json:"imageURL"`                                // URL or path where the image is stored
	UploadTime time.Time `json:"uploadTime"`                              // The time when the image was uploaded
}
