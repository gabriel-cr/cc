package main

import "fmt"
import "os"
import "encoding/base64"
import "encoding/hex"

func main() {
	data, _ := hex.DecodeString(os.Args[1])
	res := base64.StdEncoding.EncodeToString(data)
	fmt.Printf(res + "\n")
}
