package main

import "fmt"

func encode(str string) string {
	output := make([]rune, 0)
	bytes := []byte(str)

	// TODO: figure out a simpler way to do this?
	padsNeeded := 0
	if len(bytes)%3 != 0 {
		padsNeeded = 3 - (len(bytes) % 3)
	}

	// append 0 until len(bytes) is multiple of 3
	for len(bytes)%3 != 0 {
		bytes = append(bytes, 0)
	}

	// operate in groups of 3 bytes
	for i := 0; i < len(bytes); i += 3 {
		var bits uint32 = 0
		bits |= uint32(bytes[i+2])
		bits |= (uint32(bytes[i+1])) << 8
		bits |= (uint32(bytes[i])) << 16

		var ch1 uint32 = (bits & (0b111111 << 18)) >> 18
		var ch2 uint32 = (bits & (0b111111 << 12)) >> 12
		var ch3 uint32 = (bits & (0b111111 << 6)) >> 6
		var ch4 uint32 = bits & (0b111111)

		output = append(output, getCharacterFromIndex(ch1))
		output = append(output, getCharacterFromIndex(ch2))
		output = append(output, getCharacterFromIndex(ch3))
		output = append(output, getCharacterFromIndex(ch4))
	}

	for i := range padsNeeded {
		output[len(output)-i-1] = Padding
	}

	return string(output)
}

var charMap = map[uint32]rune{
	// Upper-case letters
	0: 'A', 1: 'B', 2: 'C', 3: 'D', 4: 'E', 5: 'F', 6: 'G', 7: 'H',
	8: 'I', 9: 'J', 10: 'K', 11: 'L', 12: 'M', 13: 'N', 14: 'O', 15: 'P',
	16: 'Q', 17: 'R', 18: 'S', 19: 'T', 20: 'U', 21: 'V', 22: 'W', 23: 'X',
	24: 'Y', 25: 'Z',
	// Lower-case letters
	26: 'a', 27: 'b', 28: 'c', 29: 'd', 30: 'e', 31: 'f', 32: 'g', 33: 'h',
	34: 'i', 35: 'j', 36: 'k', 37: 'l', 38: 'm', 39: 'n', 40: 'o', 41: 'p',
	42: 'q', 43: 'r', 44: 's', 45: 't', 46: 'u', 47: 'v', 48: 'w', 49: 'x',
	50: 'y', 51: 'z',
	// Digits
	52: '0', 53: '1', 54: '2', 55: '3', 56: '4', 57: '5', 58: '6', 59: '7',
	60: '8', 61: '9',
	// Special characters
	62: '+',
	63: '/',
}

func getCharacterFromIndex(val uint32) rune {
	char, exists := charMap[val]

	if !exists {
		panic(fmt.Sprintf("invalid value '%d' in getCharacter", val))
	}

	return char
}
