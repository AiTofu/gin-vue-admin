package main

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGinModes(t *testing.T) {
	// 1. 默认模式（debug）
	t.Log("Default Mode:", gin.Mode())

	// 2. 设置为发布模式
	gin.SetMode(gin.ReleaseMode)
	t.Log("After SetMode to Release:", gin.Mode())

	// 3. 通过环境变量设置为测试模式
	t.Setenv("GIN_MODE", "test")
	t.Log("After setting env to test:", gin.Mode())

	// 4. 恢复为调试模式
	gin.SetMode(gin.DebugMode)
	t.Log("After SetMode back to Debug:", gin.Mode())
}

func ExampleGinServer() {
	// 设置为发布模式
	gin.SetMode(gin.ReleaseMode)

	r := gin.New() // 使用 New() 而不是 Default() 来避免默认中间件

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"mode":    gin.Mode(),
		})
	})

	// 注意：这只是示例，实际运行服务器需要在 main() 函数中
	// r.Run(":8080")
}

// 各模式的特点：
/*
1. DebugMode (默认):
   - 打印详细的日志
   - 包含完整的错误堆栈
   - 适合开发环境

2. ReleaseMode:
   - 最小化日志输出
   - 隐藏详细错误信息
   - 优化性能
   - 适合生产环境

3. TestMode:
   - 用于测试环境
   - 禁用日志颜色
   - 适合自动化测试
*/
