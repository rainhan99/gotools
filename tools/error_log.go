package tools

import "fmt"

// project gotools
// file tools/error_log.go
// Create by RainHan on 2024/06/12 14:43:50

func UseError(err error) {
	if err != nil {
		fmt.Println(err)
	}
	return
}

func UseErrorf(msg string, err error) {
	if err != nil {
		fmt.Printf("%s=>%v\n", msg, err)
	}
	return
}

func UseNilErrorf(msg string, err error) {
	if err == nil {
		fmt.Printf("%s\n", msg)
	}
}
