package ascii_conversion

import (
	"testing"
)

func TestAsciiToBinary(t *testing.T) {
	input := "Hello"
	expected := "01001000 01100101 01101100 01101100 01101111"
	result := AsciiToBinary(input)
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}

func TestAsciiToHex(t *testing.T) {
	input := "Hello"
	expected := "48 65 6c 6c 6f"
	result := AsciiToHex(input)
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}

func TestBinaryToAscii(t *testing.T) {
	input := "01001000 01100101 01101100 01101100 01101111"
	expected := "Hello"
	result := BinaryToAscii(input)
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}

func TestBinaryToAsciiWithSpaces(t *testing.T) {
	input := "01001000 01100101 01101100 01101100 01101111"
	expected := "Hello"
	result := BinaryToAscii(input)
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}
