package service

import (
	"backend/dao/model"
	"gorm.io/gorm"
)

type ImageService struct {
	DB *gorm.DB
}

// CreateImage 添加新的图像记录
func (s *ImageService) CreateImage(image *model.Image) error {
	return s.DB.Create(image).Error
}

// GetImageByID 根据ID获取图像记录
func (s *ImageService) GetImageByID(id int) (*model.Image, error) {
	var image model.Image
	result := s.DB.First(&image, id)
	return &image, result.Error
}

// UpdateImage 更新现有的图像记录
func (s *ImageService) UpdateImage(image *model.Image) error {
	return s.DB.Save(image).Error
}

// DeleteImage 根据ID删除图像记录
func (s *ImageService) DeleteImage(id int) error {
	return s.DB.Delete(&model.Image{}, id).Error
}

// GetLatestImageIDByUserID 根据用户ID获取最新的图像ID
func (s *ImageService) GetLatestImageIDByUserID(userID int) (int, error) {
	var image model.Image
	result := s.DB.Where("user_id = ?", userID).Order("upload_time DESC").First(&image)
	if result.Error != nil {
		return 0, result.Error
	}
	return image.ImageID, nil
}

// GetAllImageIDsByUserID 根据用户ID获取所有图像ID
func (s *ImageService) GetAllImageIDsByUserID(userID int) ([]int, error) {
	var images []model.Image
	var imageIDs []int

	result := s.DB.Select("image_id").Where("user_id = ?", userID).Order("upload_time DESC").Find(&images)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, img := range images {
		imageIDs = append(imageIDs, img.ImageID)
	}

	return imageIDs, nil
}
