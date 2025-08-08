package strings

import (
	"fmt"
	"strings"
	"unicode"
)

func ReverseString() {

	str := "hello"
	s := ""

	for i := len(str) - 1; i >= 0; i-- {
		s += string(str[i])
	}

	fmt.Println(s)
}

func Palindrome() {
	str := "madamss"
	s := ""

	for i := len(str) - 1; i >= 0; i-- {
		s += string(str[i])
	}

	if str == s {
		fmt.Println("palindrome")
	} else {
		fmt.Println("not palindrome")
	}
}

func Vowels() {
	str := "helloossssssss"
	str = strings.ToLower(str)
	count := 0
	constant := 0
	for _, val := range str {
		if unicode.IsLetter(val) {
			switch val {
			case 'a', 'e', 'i', 'o', 'u':
				count++
			default:
				constant++
			}
		}
	}

	fmt.Println(count, constant)
}
