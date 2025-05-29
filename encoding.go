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

		output = append(output, getCharacter(ch1))
		output = append(output, getCharacter(ch2))
		output = append(output, getCharacter(ch3))
		output = append(output, getCharacter(ch4))

		// Prints for debugging
		// fmt.Printf("octet 1 : %08b\n", bytes[i])
		// fmt.Printf("octet 2 :         %08b\n", bytes[i+1])
		// fmt.Printf("octet 3 :                 %08b\n", bytes[i+2])
		// fmt.Printf("combined: %024b\n", bits)
		// fmt.Printf("ch1     : %06b\n", ch1)
		// fmt.Printf("ch2     :       %06b\n", ch2)
		// fmt.Printf("ch3     :             %06b\n", ch3)
		// fmt.Printf("ch4     :                   %06b\n", ch4)
	}

	for i := range padsNeeded {
		output[len(output)-i-1] = Padding
	}

	return string(output)
}

func getCharacter(val uint32) rune {
	if val >= 64 {
		panic(fmt.Sprintf("invalid value '%d' in getCharacter", val))
	}

	if val < 26 { // Upper-case
		return rune(val + 'A')
	} else if val < 52 { // Lower-case
		return rune((val - 26) + 'a')
	} else if val < 62 { // digits
		return rune(val - 52 + '0')
	} else if val == 62 {
		return '+'
	} else {
		return '/'
	}
}
