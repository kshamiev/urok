package tpl

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"io"
	"log"
	"os"
	"strings"
)

// GetImgBase64File преобразование файла в форматированную строку с картинкой для шаблона html (pdf)
func GetImgBase64File(filename string) (string, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return GetImgBase64Byte(buf), nil
}

// GetImgBase64Byte преобразование байтов в форматированную строку с картинкой для шаблона html (pdf)
func GetImgBase64Byte(buf []byte) string {
	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(buf)
}

func WordReplace(oldFilename, newFilename string, replaceWord map[string]string) error {
	zipReader, err := zip.OpenReader(oldFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer zipReader.Close()

	newFile, err := os.Create(newFilename)
	if err != nil {
		return err
	}
	defer newFile.Close()

	var k, v, newContent string
	var writer io.Writer
	var readCloser io.ReadCloser
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(newFile)
	for _, file := range zipReader.File {
		if writer, err = zipWriter.Create(file.Name); err != nil {
			return err
		}
		if readCloser, err = file.Open(); err != nil {
			return err
		}
		buf.Reset()
		_, _ = buf.ReadFrom(readCloser)

		if file.Name == "word/document.xml" {
			newContent = string(buf.Bytes())
			for k, v = range replaceWord {
				newContent = strings.Replace(newContent, k, v, -1)
			}
			_, _ = writer.Write([]byte(newContent))
		} else {
			_, _ = writer.Write(buf.Bytes())
		}
	}
	return zipWriter.Close()
}
