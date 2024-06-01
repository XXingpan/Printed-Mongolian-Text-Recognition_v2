package utils

import (
	"fmt"
	"os/exec"
)

func RunPythonScript(scriptPath string, interpreterPath string) (string, error) {
	// 执行Python脚本的命令
	cmd := exec.Command(interpreterPath, scriptPath)

	// 执行命令并获取输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

// RunPythonScriptAndKeepServerRunning 函数用于执行Python脚本并保持服务器一直运行
func RunOCR(scriptPath string, interpreterPath string) error {
	// 异步执行Python脚本
	go func() {
		// 执行Python脚本并获取输出
		output, err := RunPythonScript(scriptPath, interpreterPath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// 输出Python脚本的结果
		fmt.Println("Python script output:", output)
	}()

	// 主程序保持运行
	fmt.Println("Server is running...")
	// 这里可以添加一些其他操作或者直接返回，保持服务器一直运行
	select {}
}

//func main() {
//	// 要执行的Python脚本
//	pythonScript := "D:\\Projects\\Python\\DataSet\\ocr_model\\run_ocr.py"
//
//	// 指定Python解释器的路径
//	pythonInterpreter := "D:\\Anaconda\\anaconda3\\envs\\DataSet\\python.exe"
//
//	// 执行Python脚本并获取输出
//	output, err := RunPythonScript(pythonScript, pythonInterpreter)
//	if err != nil {
//		fmt.Println("Error:", err)
//		return
//	}
//
//	// 输出Python脚本的结果
//	fmt.Println("Python script output:", output)
//}
