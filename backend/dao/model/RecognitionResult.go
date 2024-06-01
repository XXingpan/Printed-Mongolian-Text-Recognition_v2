package model

import "time"

// RecognitionResult represents the recognition result table in the database.
type RecognitionResult struct {
	ResultID              int       `gorm:"primaryKey;autoIncrement" json:"resultID"` // Unique identifier for the recognition result
	ImageID               int       `json:"imageID"`                                  // Identifier for the image associated with this result
	RecognitionContentURL string    `json:"recognitionContentURL"`                    // URL or path where the recognition content is stored
	RecognitionTime       time.Time `json:"recognitionTime"`                          // The time when the recognition result was recorded
}
