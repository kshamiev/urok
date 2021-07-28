package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"os"
	"path/filepath"
	"runtime"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/ean"
	"github.com/disintegration/imaging"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
)

// apt-get install ttf-xfree86-nonfree
func main() {

	// 13 digits
	code := "5901234123457"

	fmt.Println("Generating Datamatrix barcode for : ", code)

	// see https://godoc.org/github.com/boombuler/barcode/ean
	bcode, err := ean.Encode(code)

	// uncomment if checksum missmatch
	// fmt.Println(err)

	if err != nil {
		fmt.Printf("String %s cannot be encoded\n", code)
		os.Exit(1)
	}

	// scale to 100x100
	bcodeNew, err := barcode.Scale(bcode, 100, 100)

	if err != nil {
		fmt.Println("EAN scaling error : ", err)
		os.Exit(1)
	}

	// now we want to append the code at the bottom
	// of the EAN

	// Create an new image with text data
	// From https://github.com/llgcode/draw2d.samples/tree/master/helloworld
	// Set the global folder for searching fonts
	_, filePath, _, _ := runtime.Caller(0)
	draw2d.SetFontFolder(filepath.Dir(filePath))

	// Initialize the graphic context on an RGBA image
	img := image.NewRGBA(image.Rect(0, 0, 250, 50))

	// set background to white
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

	gc := draw2dimg.NewGraphicContext(img)

	gc.FillStroke()

	// Set the font Montereymbi.ttf
	gc.SetFontData(draw2d.FontData{"Monterey", draw2d.FontFamilyMono, draw2d.FontStyleBold | draw2d.FontStyleItalic})
	// Set the fill text color to black
	gc.SetFillColor(image.Black)

	gc.SetFontSize(12)

	gc.FillStringAt(code, 50, 20)

	// create a new blank image with white background
	newImg := imaging.New(300, 200, color.NRGBA{255, 255, 255, 255})

	// paste the codabar to new blank image
	newImg = imaging.Paste(newImg, bcodeNew, image.Pt(100, 30))

	// paste the text to the new blank image
	newImg = imaging.Paste(newImg, img, image.Pt(50, 150))

	err = draw2dimg.SaveToPngFile(filepath.Dir(filePath)+"/ean.png", newImg)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// everything ok
	fmt.Println("EAN barcode generated and saved to ean.png")

}
