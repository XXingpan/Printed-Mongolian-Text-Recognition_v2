package utils

import (
	"os"
	"path/filepath"
)

func ClearCache(folderPath string) error {
	// 打开文件夹
	dir, err := os.Open(folderPath)
	if err != nil {
		return err
	}
	defer dir.Close()

	// 读取文件夹中的文件和子文件夹
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return err
	}

	// 遍历文件夹中的每个文件和子文件夹
	for _, fileInfo := range fileInfos {
		// 构建文件/子文件夹的完整路径
		fullPath := filepath.Join(folderPath, fileInfo.Name())

		// 如果是文件夹，则递归删除
		if fileInfo.IsDir() {
			err = ClearCache(fullPath)
			if err != nil {
				return err
			}
		} else {
			// 如果是文件，则删除
			err = os.Remove(fullPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
