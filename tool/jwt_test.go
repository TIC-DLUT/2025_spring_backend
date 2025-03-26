package tool

import (
	"fmt"
	"testing"
)

func TestJWT(t *testing.T) {
	token := GenerateJWToken("password", 1, "111111")
	fmt.Println(token)
	res, e := ParseJWToken("password", token)
	if e != nil {
		t.Fatal(e.Error())
		return
	}
	if res.Telephone != "111111" || res.ID != 1 {
		t.Fatal("数据变化！")
		return
	}
}
