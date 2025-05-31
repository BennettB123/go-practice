package main

import "fmt"

func decode(str string) string {
	if len(str)%4 != 0 {
		panic(fmt.Sprintf("this decoder requires padded (=) inputs"))
	}

    runes := []rune(str)
    output := make([]byte, 0)

	// operate in groups of 4 bytes
    for i := 0; i < len(runes); i += 4 {
        bits1 := getIndexFromCharacter(runes[i])
        bits2 := getIndexFromCharacter(runes[i+1])
        bits3 := getIndexFromCharacter(runes[i+2])
        bits4 := getIndexFromCharacter(runes[i+3])

        combined := (bits1 << 18) | (bits2 << 12) | (bits3 << 6) | bits4

        byte1 := byte((combined >> 16) & 0b11111111)
        byte2 := byte((combined >> 8) & 0b11111111)
        byte3 := byte(combined & 0b11111111)

        // Handle padding
        if runes[i+2] == '=' && runes[i+3] == '=' {
            output = append(output, byte1)
        } else if runes[i+3] == '=' {
            output = append(output, byte1, byte2)
        } else {
            output = append(output, byte1, byte2, byte3)
        }
    }

	return string(output)
}

var reverseCharMap = map[rune]uint32{
	// Upper-case letters
	'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4, 'F': 5, 'G': 6, 'H': 7,
	'I': 8, 'J': 9, 'K': 10, 'L': 11, 'M': 12, 'N': 13, 'O': 14, 'P': 15,
	'Q': 16, 'R': 17, 'S': 18, 'T': 19, 'U': 20, 'V': 21, 'W': 22, 'X': 23,
	'Y': 24, 'Z': 25,
	// Lower-case letters
	'a': 26, 'b': 27, 'c': 28, 'd': 29, 'e': 30, 'f': 31, 'g': 32, 'h': 33,
	'i': 34, 'j': 35, 'k': 36, 'l': 37, 'm': 38, 'n': 39, 'o': 40, 'p': 41,
	'q': 42, 'r': 43, 's': 44, 't': 45, 'u': 46, 'v': 47, 'w': 48, 'x': 49,
	'y': 50, 'z': 51,
	// Digits
	'0': 52, '1': 53, '2': 54, '3': 55, '4': 56, '5': 57, '6': 58, '7': 59,
	'8': 60, '9': 61,
	// Special characters
	'+': 62,
	'/': 63,
}

func getIndexFromCharacter(char rune) uint32 {
	// treat pads specially
	if char == Padding {
		return 0
	}

	val, exists := reverseCharMap[char]

	if !exists {
		panic(fmt.Sprintf("invalid character in text '%d'", char))
	}

	return val
}
