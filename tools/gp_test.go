package tools

import "testing"

// project gotools
// file tools/gp_test.go
// Create by RainHan on 2024/05/30 16:03:02

func Test_Ggg(t *testing.T) {
	for i := 7; i <= 15; i++ {
		password := GeneratePassword(i)
		t.Log(password)
	}
}
