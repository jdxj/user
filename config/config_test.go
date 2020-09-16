package config

import (
	"fmt"
	"testing"
)

func TestInitConfig(t *testing.T) {
	if err := Init("../config.yaml"); err != nil {
		t.Fatalf("%s\n", err)
	}
	dbCfg := DB()
	fmt.Printf("%#v\n", dbCfg)

	dbCfg.Password = "abc"
	dbCfg = DB()
	fmt.Printf("%#v\n", dbCfg)
}
