package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fp, err := os.OpenFile("input.txt", os.O_RDONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	buf := make([]byte, 32)
	var offset int64 = 0
	for {
		block, err := fp.ReadAt(buf, offset)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if block == 0 {
			break
		}
		fmt.Print(string(buf))
		offset += int64(block)
	}

}
