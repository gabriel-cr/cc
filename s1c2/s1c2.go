package main

import "fmt"
import "os"
import "encoding/hex"

func do_xor(s1 string, s2 string) string {
	s1b, _ := hex.DecodeString(s1)
	s2b, _ := hex.DecodeString(s2)

	for i := 0; i < len(s1b); i++ {
		s1b[i] ^= s2b[i]
	}

	return hex.EncodeToString(s1b)
}

func main() {
	fmt.Printf("%s\n", do_xor(os.Args[1], os.Args[2]))
}
