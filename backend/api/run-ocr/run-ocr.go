package run_ocr

import (
	"backend/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"os/exec"
)

func RunPythonScript(c *gin.Context) {

	// 要执行的Python脚本
	pythonScript := config.PythonScript

	// 指定Python解释器的路径
	pythonInterpreter := config.PythonInterpreter
	cmd := exec.Command(pythonInterpreter, pythonScript)

	// 捕获标准错误输出
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Printf("Failed to get stderr pipe: %v\n", err)
		return
	}

	// 启动命令
	if err := cmd.Start(); err != nil {
		fmt.Printf("Failed to start command: %v\n", err)
		return
	}

	// 读取标准错误输出
	errOutput := make([]byte, 0)
	buf := make([]byte, 1024)
	for {
		n, err := stderr.Read(buf)
		if n > 0 {
			errOutput = append(errOutput, buf[:n]...)
		}
		if err != nil {
			break
		}
	}

	// 等待命令执行完毕
	if err := cmd.Wait(); err != nil {
		fmt.Printf("Command execution failed: %v\n", err)
		fmt.Printf("Error output: %s\n", string(errOutput))
		return
	}

	// 如果没有错误输出
	if len(errOutput) > 0 {
		fmt.Printf("Error output: %s\n", string(errOutput))
	}
}
