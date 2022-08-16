package shot

import (
	"fmt"
	"os/exec"
)

/** 
*	goos: windows|darwin|linux
*	path: 截图保存路径
*	mode: 0: 截图并保存到粘贴板
*		  1: 截图并保存到指定路径	
*/
// 这里主要实现的是截图
func Shot(goos string , path string , mode int){
	// 判断是否是windows系统
	if goos == "windows" {
		// 如果是windows系统，则调用windows的截图函数
		shotWindows(path, mode)
	} else if goos == "darwin" {
		// 如果是mac系统，则调用mac的截图函数
		shotMac(path, mode)
	} else {
		// 如果是linux系统，则调用linux的截图函数
		shotLinux(path, mode)
	}
}

// windows的截图函数
func shotWindows(path string, mode int) {
	// 根据mode选择截图方式
	if mode == 0 {
		// 如果是保存到粘贴板，则调用windows的截图函数
		shotWindows0(path)
	}
	if mode == 1 {
		// 如果是保存到指定路径，则调用windows的截图函数
		shotWindows1(path)
	}
}	

func shotWindows0(path string) {
	// 定义截图命令
	cmd := "clip"
	// 执行截图命令
	exec.Command("cmd", "/c", cmd).Run()
}

func shotWindows1(path string) {
	// 定义截图命令
	cmd := "screencapture -x " + path
	// 执行截图命令
	exec.Command("cmd", "/c", cmd).Run()
}

// mac的截图函数
func shotMac(path string, mode int) {
	fmt.Println("mac")
	// 定义截图命令
	cmd := "screencapture -x " + path
	// 执行截图命令
	exec.Command("cmd", "/c", cmd).Run()
}

// linux的截图函数
func shotLinux(path string, mode int) {
	// 定义截图命令
	cmd := "screencapture -x " + path
	// 执行截图命令
	exec.Command("cmd", "/c", cmd).Run()
}
