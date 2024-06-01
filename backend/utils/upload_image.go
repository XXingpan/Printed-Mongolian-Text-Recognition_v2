package utils

import (
	"backend/config"
	"backend/dao/model"
	"backend/dao/service"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func UploadImage(fileHeader *multipart.FileHeader, userID string) error {
	uploadTime := time.Now()
	db, err0 := config.SetupDatabase()
	if err0 != nil {
		log.Fatal("Failed to connect to database:", err0)
	}

	timestamp := strconv.FormatInt(uploadTime.Unix(), 10)
	fileExt := filepath.Ext(fileHeader.Filename)
	fileName := "image_" + timestamp + fileExt
	fileName1 := "image_upload" + fileExt

	// 构建相对路径
	foldPath := filepath.Join(".", "internal", "images", "upload", userID)

	foldPath1 := filepath.Join(".", "internal", "crnn", "upload")
	err := ClearCache(foldPath1)
	if err != nil {
		return err
	}
	// 确保目录存在
	if err := os.MkdirAll(foldPath, 0755); err != nil {
		fmt.Printf("创建目录出错: %v\n", err)
		return err
	}

	if err := os.MkdirAll(foldPath1, 0755); err != nil {
		fmt.Printf("创建目录出错: %v\n", err)
		return err
	}

	filePath := filepath.Join(foldPath, fileName)
	filePath1 := filepath.Join(foldPath1, fileName1)

	imageService := service.ImageService{DB: db}
	userid, err := strconv.Atoi(userID)
	newImage := model.Image{
		UserID:     userid,
		ImageURL:   filePath,
		UploadTime: uploadTime,
	}
	err = imageService.CreateImage(&newImage)
	if err != nil {
		return err
	}

	err = SaveImageToFile(fileHeader, filePath)
	if err != nil {
		// 处理错误
	}

	err = SaveImageToFile(fileHeader, filePath1)
	if err != nil {
		// 处理错误
	}

	return err
}

func SaveImageToFile(fileHeader *multipart.FileHeader, filePath string) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	return err
}
