package main

import "fmt"
import "os"
import "encoding/hex"
import "strconv"

func count_characters(byte_arr []byte) [256]byte {
	var result [256]byte
	for i := 0; i < len(byte_arr); i++ {
		result[byte_arr[i]]++
	}

	for i := 0; i < 256; i++ {
		//	fmt.Printf("%d ", result[i])
	}
	//fmt.Printf("\n")
	return result
}

func print_vowels(chars []byte) {
	cnt := count_characters(chars)

	prt := "aeiouAEIOU"

	total := 0
	for i := 0; i < len(prt); i++ {
		fmt.Printf("%c %d, ", prt[i], cnt[prt[i]])
		total += int(cnt[prt[i]])
	}
	fmt.Printf(" TOT = %d", total)
}

func do_decode(arr []byte, num byte) []byte {
	var buff = make([]byte, len(arr))

	for i := 0; i < len(arr); i++ {
		buff[i] = arr[i] ^ num
	}

	return buff
}

func do_break(s1 string) string {
	s1b, _ := hex.DecodeString(s1)

	for i := 0; i < 256; i++ {

		buff := do_decode(s1b, byte(i))

		fmt.Printf("%d --- ", i)
		print_vowels(buff)
		fmt.Printf("\n")
	}

	return hex.EncodeToString(s1b)
}

func main() {
	if len(os.Args) == 2 {
		fmt.Printf("%s\n", do_break(os.Args[1]))
	} else {
		stb, _ := hex.DecodeString(os.Args[1])

		num, _ := strconv.Atoi(os.Args[2])
		buff := do_decode(stb, byte(num))

		fmt.Printf("%s\n", buff)
	}
}
