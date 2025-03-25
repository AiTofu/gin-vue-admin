package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	// 1. 初始化 Viper
	v := viper.New()
	v.SetConfigName("config") // 配置文件名称（无扩展名）
	v.SetConfigType("yaml")   // 配置文件类型
	v.AddConfigPath(".")      // 查找配置文件所在的路径

	// 2. 首次读取配置文件
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	// 3. 注册配置文件变化的回调函数
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s\n", e.Name)
		fmt.Printf("New port value: %d\n", v.GetInt("server.port"))
		fmt.Printf("New mode value: %s\n", v.GetString("server.mode"))
	})

	// 4. 开始监听配置文件变化
	v.WatchConfig()
	fmt.Println("Starting to watch config file...")

	// 5. 保持程序运行
	for {
		time.Sleep(time.Second)
		// 你可以修改 config.yaml 文件来测试配置变化
		// 比如修改 server.port 或 server.mode 的值
	}
}

/*
示例配置文件 (config.yaml):
server:
  port: 8080
  mode: "debug"
*/

/*
Viper 监听配置文件的工作原理：

1. 初始化阶段：
   - 创建 fsnotify.Watcher
   - 获取配置文件的完整路径
   - 开始监听配置文件所在目录

2. 事件处理：
   - Write：配置文件被修改
   - Create：配置文件被创建
   - Remove：配置文件被删除
   - Rename：配置文件被重命名

3. 同步机制：
   - initWG：确保监听器初始化完成
   - eventsWG：处理文件事件

4. 配置更新：
   - 检测到变化后重新读取配置
   - 触发 OnConfigChange 回调
   - 使用新的配置值

5. 错误处理：
   - 监听器创建失败
   - 配置文件读取错误
   - 文件系统事件错误
*/
