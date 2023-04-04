package ascii_conversion

import (
	"fmt"
	"strings"
)

func AsciiToBinary(ascii string) string {
	var binaryStrings []string
	for _, char := range ascii {
		binaryStrings = append(binaryStrings, fmt.Sprintf("%08b", char))
	}
	return strings.Join(binaryStrings, " ")
}

func AsciiToHex(input string) string {
	var hexStrings []string
	for _, char := range input {
		hexStrings = append(hexStrings, fmt.Sprintf("%02x", char))
	}
	return strings.Join(hexStrings, " ")
}

func BinaryToAscii(binary string) string {
	var asciiStrings []string
	for _, char := range strings.Split(binary, " ") {
		asciiStrings = append(asciiStrings, string(rune(binaryToDecimal(char))))
	}
	return strings.Join(asciiStrings, "")
}

func HexToAscii(hex string) string {
	var asciiStrings []string
	for _, char := range strings.Split(hex, " ") {
		asciiStrings = append(asciiStrings, string(rune(hexToDecimal(char))))
	}
	return strings.Join(asciiStrings, "")
}

func binaryToDecimal(binary string) int {
	var decimal int
	for i, char := range binary {
		if char == '1' {
			decimal += 1 << uint(len(binary)-i-1)
		}
	}
	return decimal
}

func hexToDecimal(hex string) int {
	var decimal int
	for i, char := range hex {
		if char >= '0' && char <= '9' {
			decimal += int(char-'0') << uint(len(hex)-i-1) * 4
		} else if char >= 'a' && char <= 'f' {
			decimal += int(char-'a'+10) << uint(len(hex)-i-1) * 4
		} else if char >= 'A' && char <= 'F' {
			decimal += int(char-'A'+10) << uint(len(hex)-i-1) * 4
		}
	}
	return decimal
}

// Path: main.go
