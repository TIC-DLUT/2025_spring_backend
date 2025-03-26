package tool

import (
	"testing"
)

func TestMD5(t *testing.T) {
	if GenerateMD5("admin") != "21232f297a57a5a743894a0e4a801fc3" {
		t.Fatal("generate MD5 Error")
	}
}
