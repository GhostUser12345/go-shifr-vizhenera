package main

import (
	"fmt"
)

func main() {
	//var message = "Hello, Vizhener!"
	var encryptedMessage = "CSOITEUIWUIZNSROCNKFD"
	var keyWord = "GOLANG"
	var decryptedMessage = decode(encryptedMessage, keyWord)
	fmt.Println(decryptedMessage)
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
