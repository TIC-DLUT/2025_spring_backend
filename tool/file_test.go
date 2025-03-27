package tool

import "testing"

func TestFileExist(t *testing.T) {
	if !FileExist("file.go") {
		t.Fatal("判断错误")
		return
	}
	if FileExist("something_unexisted.go") {
		t.Fatal("判断错误")
		return
	}
}
