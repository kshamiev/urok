package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

func main() {
	// проверка FS
	const p = "data/session/test.txt"
	err := os.MkdirAll(path.Dir(p), 0o700)
	if err != nil {
		panic(err)
	}
	if err = os.WriteFile(p, []byte("test"), 0o600); err != nil {
		panic(err)
	}
	if err = os.Remove(p); err != nil {
		panic(err)
	}

	// открываем файл
	f, err := os.Open("sample/files/sum/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// пример сохранения в другой файл
	fp, err := os.OpenFile("sample/files/sum/input_new.txt", os.O_CREATE|os.O_WRONLY, 0o600)
	if _, err = io.Copy(fp, f); err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	// пример контрольной суммы 1
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x\n", h.Sum(nil))

	// пример контрольной суммы 2
	buf := &bytes.Buffer{}
	if _, err := io.Copy(buf, f); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x\n", md5.Sum(buf.Bytes()))
}
