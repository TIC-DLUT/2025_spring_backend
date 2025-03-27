package tool

import "os"

// 判断文件或者文件夹是否存在
func FileExist(_path string) bool {
	// 系统函数如果能取到就是存在，否则不存在
	_, e := os.Stat(_path)
	return e == nil
}

// 如果不存在就创建对应文件夹
func CreateDir(_path string) {
	if !FileExist(_path) {
		os.MkdirAll(_path, os.ModePerm)
	}
}
