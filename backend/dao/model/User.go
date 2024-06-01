package model

// User represents the user table in the database.
type User struct {
	UserID   int    `gorm:"primaryKey;autoIncrement" json:"userID"` // Unique identifier for the user
	Username string `json:"username"`                               // User's username
	Password string `json:"password"`                               // User's password
	Email    string `json:"email"`                                  // User's email address
}
