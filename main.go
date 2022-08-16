package main

import (
	"fmt"
	"screenshot/module/SysType"
	"screenshot/module/Shot"
)

// 使用sysType.go中的GetSysType函数
// 在main.go中调用该函数
// 并输出结果
func main() {
	fmt.Println(systype.GetSysType())
	// 定义变量
	var goos string = systype.GetSysType()


	

	// 当用户输入快捷键时，调用Shot函数
	systype.Shot(goos, "", 0)
}

