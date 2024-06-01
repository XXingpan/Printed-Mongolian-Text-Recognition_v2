package service

import (
	"backend/dao/model"
	"gorm.io/gorm"
)

type RecognitionResultService struct {
	DB *gorm.DB
}

// CreateRecognitionResult 添加新的识别结果记录
func (s *RecognitionResultService) CreateRecognitionResult(result *model.RecognitionResult) error {
	return s.DB.Create(result).Error
}

// GetRecognitionResultByID 根据ID获取识别结果记录
func (s *RecognitionResultService) GetRecognitionResultByID(id int) (*model.RecognitionResult, error) {
	var result model.RecognitionResult
	resultQuery := s.DB.First(&result, id)
	return &result, resultQuery.Error
}

// UpdateRecognitionResult 更新现有的识别结果记录
func (s *RecognitionResultService) UpdateRecognitionResult(result *model.RecognitionResult) error {
	return s.DB.Save(result).Error
}

// DeleteRecognitionResult 根据ID删除识别结果记录
func (s *RecognitionResultService) DeleteRecognitionResult(id int) error {
	return s.DB.Delete(&model.RecognitionResult{}, id).Error
}

// GetRecognitionResultIDsByImageID 根据图像ID获取识别结果ID
func (s *RecognitionResultService) GetRecognitionResultIDsByImageID(imageID int) ([]int, error) {
	var results []model.RecognitionResult
	var resultIDs []int

	result := s.DB.Select("result_id").Where("image_id = ?", imageID).Find(&results)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, res := range results {
		resultIDs = append(resultIDs, res.ResultID)
	}

	return resultIDs, nil
}
