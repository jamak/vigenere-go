package main

import (
	"fmt"
	"strings"
)

var letters = map[string]int{
	"a": 0, "b": 1, "c": 2, "d": 3, "e": 4, "f": 5, "g": 6, "h": 7, "i": 8,
	"j": 9, "k": 10, "l": 11, "m": 12, "n": 13, "o": 14, "p": 15, "q": 16,
	"r": 17, "s": 18, "t": 19, "u": 20, "v": 21, "w": 22, "x": 23, "y": 24,
	"z": 25, "A": 0, "B": 1, "C": 2, "D": 3, "E": 4, "F": 5, "G": 6, "H": 7,
	"I": 8, "J": 9, "K": 10, "L": 11, "M": 12, "N": 13, "O": 14, "P": 15,
	"Q": 16, "R": 17, "S": 18, "T": 19, "U": 20, "V": 21, "W": 22, "X": 23,
	"Y": 24, "Z": 25,
}

var numbers = map[int]string{
	0: "A", 1: "B", 2: "C", 3: "D", 4: "E", 5: "F", 6: "G", 7: "H", 8: "I", 9: "J", 10: "K",
	11: "L", 12: "M", 13: "N", 14: "O", 15: "P", 16: "Q", 17: "R", 18: "S", 19: "T", 20: "U",
	21: "V", 22: "W", 23: "X", 24: "Y", 25: "Z",
}

//lowercase letters are between 97 - 122
//uppercase letters are between 65 - 90

func isLower(s byte) bool {
	return ((s >= 97) && (s <= 122))
}

func isUpper(s byte) bool {
	return ((s >= 65) && (s <= 90))
}

func encrypt(ptext, key string) string {
	var ctext = ""
	no_space := strings.Join(strings.Fields(ptext), "")
	for i, letter := range no_space {
		ctext += numbers[(letters[string(letter)]+letters[string(key[i%len(key)])]+26)%26]
	}
	return string(ctext)
}

func decrypt(ctext, key string) string {
	var ptext = ""
	for i, letter := range ctext {
		ptext += numbers[(letters[string(letter)]-letters[string(key[i%len(key)])]+26)%26]
	}
	return ptext
}

func main() {
	cipher := encrypt("THIS is a test message", "rbdjg")
	fmt.Println(cipher)
	plaintext := decrypt(cipher, "rbdjg")
	fmt.Println(plaintext)
}
