package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("sample/files/sum/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x\n", h.Sum(nil))

	buf := &bytes.Buffer{}
	if _, err := io.Copy(buf, f); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x\n", md5.Sum(buf.Bytes()))
}
