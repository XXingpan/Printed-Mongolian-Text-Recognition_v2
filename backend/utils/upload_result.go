package utils

import (
	"backend/config"
	"backend/dao/model"
	"backend/dao/service"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

// copyFile copies a file from src to dest
func copyFile(src, dest string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

// SaveOCRResult saves the OCR result to a directory and database
func SaveOCRResult(userID int) error {
	uploadTime := time.Now()
	db, err := config.SetupDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	imageService := service.ImageService{DB: db}
	resultService := service.RecognitionResultService{DB: db}

	// Source path for OCR result file
	srcPath := "./internal/crnn/result_text/ocr_result.json"

	// Get the latest image ID
	imageID, err := imageService.GetLatestImageIDByUserID(userID)
	if err != nil {
		return err
	}

	// Destination directory
	destDir := filepath.Join("./internal/images/result", fmt.Sprint(userID))
	err = os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		return err
	}

	// Destination file path with the image ID as the filename
	destPath := filepath.Join(destDir, fmt.Sprintf("%d.json", imageID))

	// Copy the file
	err = copyFile(srcPath, destPath)
	if err != nil {
		return err
	}

	// Create recognition result entry
	recResult := model.RecognitionResult{
		ImageID:               imageID,
		RecognitionContentURL: destPath,
		RecognitionTime:       uploadTime,
	}

	// Save recognition result to the database
	return resultService.CreateRecognitionResult(&recResult)
}

//
//func main() {
//	// Example usage
//	userID := 1
//	err := SaveOCRResult(userID)
//	if err != nil {
//		log.Fatalf("Failed to copy OCR result: %v", err)
//	} else {
//		log.Println("OCR result copied successfully")
//	}
//}
