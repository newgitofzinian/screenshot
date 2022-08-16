package mode1

// 主要实现的是保存在指定文件夹

// 根据当前时间生成文件名
var dir string
dir = "/Users/zinian/Desktop/scrshot"
// 存放图片的路径
var path string = dir + "/" + systype.GetTime() + ".png" 
fmt.Println(path)