package main

import (
	"fmt"
)

func main() {
	var encryptedMessageDecode = "CSOITEUIWUIZNSROCNKFD"
	var keyWordDecode = "GOLANG"
	var decryptedMessageDecode = decode(encryptedMessageDecode, keyWordDecode)
	fmt.Println(decryptedMessageDecode)

	var decryptedMessageEncode = "WEDIGYOULUVTHEGOPHERS"
	var keyWordEncode = "GOLANG"
	var encryptedMessageEncode = encode(decryptedMessageEncode, keyWordEncode)
	fmt.Println(encryptedMessageEncode)
}

func decode(encryptedMessage string, keyWord string) string {
	var result = ""
	var keyWordLength = len(keyWord)
	for i := 0; i < len(encryptedMessage); i++ {
		var move = byte(keyWord[i%keyWordLength]) - 'A'
		var currentLetter = encryptedMessage[i] - move
		if currentLetter < 'A' {
			currentLetter += 26
		}
		result += string(currentLetter)
	}
	return result
}

func encode(decryptedMessage string, keyWord string) string {
	var result = ""
	var keyWordLength = len(keyWord)
	for i := 0; i < len(decryptedMessage); i++ {
		var move = byte(keyWord[i%keyWordLength]) - 'A'
		var currentLetter = decryptedMessage[i] + move
		if currentLetter > 'Z' {
			currentLetter -= 26
		}
		result += string(currentLetter)
	}
	return result
}
