package utils

import (
	"io/ioutil"
)

// ReadJSONFile 从给定路径读取JSON文件并返回其内容字符串
func ReadJSONFile(filePath string) (string, error) {
	// 读取JSON文件
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// 将字节切片转换为字符串
	jsonString := string(data)

	return jsonString, nil
}

//func main() {
//	// 调用ReadJSONFile函数并获取结果
//	jsonString, err := ReadJSONFile("D:\\Projects\\Python\\DataSet\\result\\ocr_result.json")
//	if err != nil {
//		fmt.Println("读取文件时发生错误:", err)
//		return
//	}
//
//	// 打印JSON字符串
//	fmt.Println(jsonString)
//}
