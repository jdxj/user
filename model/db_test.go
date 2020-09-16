package model

import "testing"

func TestInitDB(t *testing.T) {
	err := InitDB("root", "123456", "127.0.0.1:3306", "video")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	defer mysql.Close()
}
