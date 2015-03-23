/*
  This will be a library that will implement a number of
  string functions without using any of the built in string
  functions in GO.
  The functions we will be implemetning are as follows:
  IsAlpha(str string) bool
  IsDigit(str string) bool
  ToLower(str string) string
  ToUpper(str string) string
  FindChar(str string, char string) int
  FindString(str string, substring string) int
  ReplaceChar(str string, charToReplace string, replaceWith string) string
  ReplaceString(str string, strToReplace string, replaceWith string) string
 */

package stringlib

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const (
	asciiLowecase = "abcdefghijklmnopqrstuvwxyz"
	asciiUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	decimalDigits = "0123456789"
)

func Start() {
	fmt.Println(ToLower("BUTTHASH"))
	fmt.Println(ToUpper("asss Pirate!"))
	fmt.Println(IsDigit("6"))
	fmt.Println(IsDigit("h"))
	fmt.Println(FindChar("ButtHash", "h"))
	fmt.Println(FindString("bababahbah", "baba", 0))
}

// IsAlpaha will return true if the character passed in
// is a member of the alphabet, false otherwise.
func IsAlpha(str string) bool {
	_, inString := stringIndex(asciiLowecase, ToLower(str), 0)
	return inString
}

// IsDigit will return true if the digit passed in is from
// 1 to 9, false otherwise.
func IsDigit(digit string) bool {
	_, inDigitStr := stringIndex(decimalDigits, digit, 0)
	return inDigitStr
}

// StringIndex will return the index of the substring and true
// if the substring can be found in the string, -1 for the index
// and false otherwise.
func stringIndex(str string, substring string, start int) (int, bool) {
	for i := start; i < utf8.RuneCountInString(str); i++ {
		if string(str[i]) == substring {
          return i, true
		}
	}
	return -1, false
}

// FindChar will find the index of the char we pass in
// in the string we pass in, -1 if we can't find it.
func FindChar(str string, char string) int {
	lowerStr := ToLower(str)
	lowerChar := ToLower(char)
	index, _ := stringIndex(lowerStr, lowerChar, 0)
	return index
}

// FindString will find the index of a given string and substring
// -1 if the substring does not exist in the string.
func FindString(str string, substring string, start int) int {
	index, ok := stringIndex(str, string(substring[0]), start)
	if ok && index+utf8.RuneCountInString(substring) <= utf8.RuneCountInString(str) {
		for i := 0; i < utf8.RuneCountInString(substring); i++ {
			if string(str[index+i]) != string(substring[i]) {
				return FindString(str, substring, index+1);
			}
		}
		return index
	} else {
		return -1
	}
}

// ToUpper will take in a string and uppercase all of the letters
func ToUpper(str string) string {
	return SwitchCase("upper", str)
}

// ToLower will take in a string and lowercase all of the letters.
func ToLower(str string) string {
	return SwitchCase("lower", str)
}

// SwitchCase will will turn a string into upper case or lower
// case depending on which flag we give it.
func SwitchCase(upperOrLower string, str string) string {
	var asciiToUse string
	if upperOrLower == "upper" {
		asciiToUse = asciiUppercase
	} else {
		asciiToUse = asciiLowecase
	}
	lowerSlice := make([]string,0)
	for _, val := range str {
		index, ok := stringIndex(asciiUppercase, string(val), 0)
		if ok {
			lowerSlice = append(lowerSlice, string(asciiToUse[index]))
		} else {
			index, ok := stringIndex(asciiLowecase, string(val), 0)
			if ok {
				lowerSlice = append(lowerSlice, string(asciiToUse[index]))
			} else {
				lowerSlice = append(lowerSlice, string(val))
			}
		}
	}
	return strings.Join(lowerSlice, "")
}