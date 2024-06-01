package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// SegmentWordsInMongolianImage 将蒙古文图像中的单词分割并保存到指定文件夹
func SegmentWordsInMongolianImage(imagePath string, outputDir string) error {
	// 运行 Tesseract OCR 并将结果保存到临时文件
	tempOutput := "temp_output"
	tesseractCmd := exec.Command("tesseract", imagePath, tempOutput, "-l", "mon", "stdout")
	if err := tesseractCmd.Run(); err != nil {
		return fmt.Errorf("failed to run tesseract: %v", err)
	}

	// 读取临时文件中的文本行信息
	textLines, err := readTextLinesFromFile(tempOutput + ".txt")
	if err != nil {
		return fmt.Errorf("failed to read text lines: %v", err)
	}

	// 针对每个文本行，分割单词并保存
	for i, line := range textLines {
		words := strings.Fields(line)

		// 创建输出文件夹
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}

		// 保存单词图像
		if err := saveWordImages(words, outputDir, i); err != nil {
			return fmt.Errorf("failed to save word images: %v", err)
		}

		fmt.Printf("Saved %d words from text line %d\n", len(words), i)
	}

	// 删除临时文件
	if err := os.Remove(tempOutput + ".txt"); err != nil {
		fmt.Printf("Warning: failed to delete temporary file: %v\n", err)
	}

	return nil
}

// readTextLinesFromFile 从文件中读取文本行信息
func readTextLinesFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	return lines, nil
}

// saveWordImages 保存单词图像
func saveWordImages(words []string, outputDir string, lineIndex int) error {
	//for i, word := range words {
	//	// 在这里你可以根据需要进行单词图像的生成和保存
	//	// 比如使用字体库生成单词图像，然后保存到指定文件夹中
	//}
	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <image_path> <output_dir>")
		return
	}

	imagePath := os.Args[1]
	outputDir := os.Args[2]

	err := SegmentWordsInMongolianImage(imagePath, outputDir)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
