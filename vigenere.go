package main

import (
    /* "strings" */
    "fmt"
)

/* letters := map[string]int{ */
/* 'a': 0 , 'b': 1 , 'c': 2 , 'd': 3 , 'e': 4 , 'f': 5 , 'g': 6 , 'h': 7 , 'i': 8 , */
/* 'j': 9 , 'k': 10 , 'l': 11 , 'm': 12 , 'n': 13 , 'o': 14 , 'p': 15 , 'q': 16 , */
/* 'r': 17 , 's': 18 , 't': 19 , 'u': 20 , 'v': 21 , 'w': 22 ,'x': 23 , 'y': 24 , */
/* 'z': 25 ,'A': 0 , 'B': 1 , 'C': 2 , 'D': 3 , 'E': 4 , 'F': 5 , 'G': 6 , 'H': 7 , */
/* 'I': 8 , 'J': 9 , 'K': 10 , 'L': 11 , 'M': 12 , 'N': 13 , 'O': 14 , 'P': 15 , */
/* 'Q': 16 , 'R': 17 , 'S': 18 , 'T': 19 , 'U': 20 , 'V': 21 , 'W': 22 ,'X': 23 , */
/* 'Y': 24 , 'Z': 25 */
/*     } */

//lowercase letters are between 97 - 122
//uppercase letters are between 65 - 90

func isLower(s byte) bool {
        return ((s >= 97) && (s <= 122))
}

func isUpper(s byte) bool {
        return ((s >= 65) && (s <= 90))
}

func encrypt(ptext, key string) string {
    ctext := make([]byte, len(ptext))
    for i, letter := range ptext {
	    p := byte(letter)
        k := byte(key[i%len(key)])
	    if isLower(p) {
                ctext[i] = p - 97
                if isLower(k) {
                    ctext[i] = ((p + key[i%len(key)]) % 97 ) % 26
                } else {
                    ctext[i] = ((p + key[i%len(key)]%65) % 97 ) % 26
                }
                ctext[i] += 97
	    }else if isUpper(p) {
                ctext[i] = p - 65
                if isUpper(k) {
                    ctext[i] = ((p + key[i%len(key)]) % 65 ) % 26
                } else {
                    ctext[i] = ((p + key[i%len(key)]%97) % 65 ) % 26
                }
                ctext[i] += 65
	    }else {
                ctext[i] = p
	    }
    }
    return string(ctext)
}

func decrypt(ctext, key string) string {
    ptext := make([]byte, len(ctext))
    for i, letter := range ctext {
	    c := byte(letter)
        k := byte(key[i%len(key)])
	    if isLower(c) {
                ptext[i] = c - 97
                if isLower(k) {
                    ptext[i] = ((c - key[i%len(key)]) % 97 ) % 26
                } else {
                    ptext[i] = ((c - key[i%len(key)]%65) % 97 ) % 26
                }
                ptext[i] += 97
	    }else if isUpper(c) {
                ptext[i] = c - 65
                if isUpper(k) {
                    ptext[i] = ((c - key[i%len(key)]) % 65 ) % 26
                } else {
                    ptext[i] = ((c - key[i%len(key)]%97) % 65 ) % 26
                }
                ptext[i] += 65
	    }else {
                ptext[i] = c
	    }
    }
    return string(ptext)
}

func main() {
    cipher := encrypt("THIS is a test message", "aaaa")
    fmt.Println(cipher)
    fmt.Println(decrypt(cipher,"aaaa"))
}
