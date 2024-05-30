package tools

import (
	"math/rand"
	"strings"
)

// project gotools
// file tools/password
// Create by RainHan on 2024/05/30 15:53:33

const (
	LowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	UppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numbers          = "0123456789"
	Symbols          = "!@#$%^&*()_+{}:<>?|"
	AllCharacters    = LowercaseLetters + UppercaseLetters + Numbers + Symbols
)

func GeneratePassword(length int) string {
	var password strings.Builder
	if length < 8 {
		return "Password length must be at least 8 characters"
	}

	for i := 0; i < 1; i++ {
		random := rand.Intn(len(LowercaseLetters))
		password.WriteString(string(LowercaseLetters[random]))
	}
	for i := 0; i < 1; i++ {
		random := rand.Intn(len(UppercaseLetters))
		password.WriteString(string(UppercaseLetters[random]))
	}
	for i := 0; i < 1; i++ {
		random := rand.Intn(len(Numbers))
		password.WriteString(string(Numbers[random]))
	}
	for i := 0; i < 1; i++ {
		random := rand.Intn(len(Symbols))
		password.WriteString(string(Symbols[random]))
	}

	for i := 0; i < length-4; i++ {
		random := rand.Intn(len(AllCharacters))
		password.WriteString(string(AllCharacters[random]))
	}

	return password.String()
}
