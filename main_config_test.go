package main_test

import (
	"fmt"
	"os"
	"testing"
	"tiktok-lite/bootstrap"
	"tiktok-lite/util"
)

// 多版本测试
// 如果成功会在控制台输出 config 配置对象信息
func TestMultiVersionViper(t *testing.T) {
	bootstrap.InitConfig()
}

// 测试环境变量 `GO_ENV` 的获取
// 自测配置是否正确
func TestGetEnv(t *testing.T) {
	env := os.Getenv("GO_ENV")
	fmt.Printf("GO_ENV >> %#q\n", env)
	if env == "" {
		fmt.Println("def")
	}
}

func TestPath(t *testing.T) {
	fmt.Println(util.ParentPath2Suffix("/usr/home"))
}
