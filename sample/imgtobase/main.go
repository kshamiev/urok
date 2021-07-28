package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
)

func Home(w http.ResponseWriter, r *http.Request) {
	imgFile, err := os.Open("product-name.png") // a QR code image

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer imgFile.Close()

	// create a new buffer base on file size
	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	// read file content into buffer
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	// if you create a new image instead of loading from file, encode the image to buffer instead with png.Encode()

	// png.Encode(&buf, image)

	// convert the buffer bytes to base64 string - use buf.Bytes() for new image
	imgBase64Str := base64.StdEncoding.EncodeToString(buf)

	// Embed into an html without PNG file
	img2html := "<html><body><img src=\"data:image/png;base64," + imgBase64Str + "\" /></body></html>"

	w.Write([]byte(fmt.Sprintf(img2html)))

}

func main() {
	// http.Handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)

	http.ListenAndServe(":8080", mux)
}
