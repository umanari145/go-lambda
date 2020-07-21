package main

import (
	"fmt"
	"testing"
)

func TestExample(t *testing.T) {
	req := loadSource()
	fmt.Println(req)
}
