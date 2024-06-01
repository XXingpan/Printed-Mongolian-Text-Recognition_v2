package tr_run

import (
	"backend/utils"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)

func TrRun(c *gin.Context) {
	// 获取UserID
	userID := c.PostForm("UserID")

	startTime := time.Now()
	// Parse request body
	err := c.Request.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "Bad Request"})
		return
	}

	// Get image from form or base64 string
	var img image.Image
	if uploadFile, _, err := c.Request.FormFile("file"); err == nil {
		defer uploadFile.Close()
		img, _, err = image.Decode(uploadFile)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "无法解析上传的图片"})
			return
		}
	} else if uploadImgbase64 := c.PostForm("img"); uploadImgbase64 != "" {
		imgBytes, err := base64.StdEncoding.DecodeString(uploadImgbase64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "无效的base64图片数据"})
			return
		}
		img, _, err = image.Decode(bytes.NewReader(imgBytes))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "无法解码base64图片数据"})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "未找到上传的图片"})
		return
	}

	compressStr := c.DefaultPostForm("compress", "")

	var compress int
	if compressStr == "" {
		compress = 1600
	} else {
		compress, err = strconv.Atoi(compressStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "无效的压缩参数"})
			return
		}
	}
	// 压缩图片
	if compress > 0 {
		img = imaging.Resize(img, compress, 0, imaging.Lanczos)
	}

	// 保存上传的图像到本地文件
	downloadFile, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		// 处理错误
	}
	defer downloadFile.Close()
	err = utils.UploadImage(fileHeader, userID)
	if err != nil {
		log.Fatalf("Image recognition failure: %v", err)
	} else {
		log.Println("Image Recognition Success")
	}
	resp, err := http.Get("http://localhost:8000/api/run-ocr")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "无法发送图像"})
		return
	}
	defer resp.Body.Close()
	//调用ReadJSONFile函数并获取结果
	ocrResult, err := utils.ReadJSONFile("./internal/crnn/result_text/ocr_result.json")

	//将识别结果保存到数据库，对应的图像ID为最新
	userid, err := strconv.Atoi(userID)
	err = utils.SaveOCRResult(userid)
	if err != nil {
		log.Fatalf("Failed to save OCR result: %v", err)
	} else {
		log.Println("OCR result copied successfully")
	}
	//fmt.Printf("\n(%v)\n", ocrResult)
	if err != nil {
		fmt.Println("读取文件时发生错误:", err)
		return
	}

	// Prepare response data
	responseData := make(map[string]interface{})
	responseData["code"] = http.StatusOK
	responseData["msg"] = "success"

	isDrawStr := c.DefaultPostForm("is_draw", "1")
	isDraw, err := strconv.Atoi(isDrawStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "无效的is_draw参数"})
		return
	}
	if isDraw == 0 {
		responseData["data"] = map[string]interface{}{
			"raw_out": ocrResult,
		}
	} else {
		// Convert image to base64 string
		// 读取 result.jpg 文件
		resultFile, err1 := os.Open("./internal/crnn/preprocess/label_image/label.jpg")
		if err1 != nil {
			fmt.Println("无法打开文件:", err)
			return
		}

		// 解码 JPEG 图像
		resultImg, _, err2 := image.Decode(resultFile)
		if err2 != nil {
			fmt.Println("无法解码 JPEG 图像:", err)
			return
		}
		defer resultFile.Close()
		var imgBuffer bytes.Buffer

		err := jpeg.Encode(&imgBuffer, resultImg, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "无法转换图片格式"})
			return
		}
		imgBase64 := base64.StdEncoding.EncodeToString(imgBuffer.Bytes())
		//fmt.Printf("%s", imgBase64)
		responseData["data"] = map[string]interface{}{
			"img_detected": "data:image/jpeg;base64," + imgBase64,
			"raw_out":      ocrResult,
			"speed_time":   time.Since(startTime).Seconds(),
		}
	}

	// Convert response data to JSON
	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "无法生成JSON响应"})
		return
	}

	// Write response
	c.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
}
