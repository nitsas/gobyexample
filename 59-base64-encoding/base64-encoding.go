package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "abc123!?$*&()'-=@~"
	fmt.Println("s:", s)

	sStdBase64 := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println("sStdBase64 := base64.StdEncoding.EncodeToString([]byte(s))")
	fmt.Println("sStdBase64:", sStdBase64)
	sStdBase64Decoded, err := base64.StdEncoding.DecodeString(sStdBase64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("base64.StdEncoding.DecodeString(sStdBase64): %s, %e\n", sStdBase64Decoded, err)
	fmt.Println()

	sURLBase64 := base64.URLEncoding.EncodeToString([]byte(s))
	fmt.Println("sURLBase64 := base64.URLEncoding.EncodeToString([]byte(s))")
	fmt.Println("sURLBase64:", sURLBase64)
	sURLBase64Decoded, err := base64.URLEncoding.DecodeString(sURLBase64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("base64.URLEncoding.DecodeString(sURLBase64): %s, %e\n", sURLBase64Decoded, err)
	fmt.Println()
}
