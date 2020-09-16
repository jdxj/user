package model

import (
	"fmt"
	"testing"
)

func initdb() {
	err := InitDB("root", "123456", "127.0.0.1:3306", "video")
	if err != nil {
		panic(err)
	}
}

func TestLoginCheck(t *testing.T) {
	initdb()
	defer mysql.Close()

	u, err := LoginCheck("jdxj", "jdxj")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%#v\n", u)
}
