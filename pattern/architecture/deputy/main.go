// Заместитель
// Ни чего нового
package main

import "fmt"

func main() {
	proxy := ImageProxy{FileName: "1.png"}
	// operation without creating a RealSubject
	fileName := proxy.FileName
	// forwarded to the RealSubject
	proxy.Draw()

	fmt.Println(fileName)
}

// Graphic is Subject
type Graphic interface {
	Draw()
}

// Image is RealSubjec
type Image struct {
	FileName string
}

// Draw is Request()
func (im Image) Draw() {
	fmt.Println("draw " + im.FileName)
}

// ImageProxy is Proxy
type ImageProxy struct {
	FileName string
	_image   *Image
}

// GetImage creates an Subject
func (ip ImageProxy) GetImage() *Image {
	if ip._image == nil {
		ip._image = &Image{ip.FileName}
	}
	return ip._image
}

// Draw is Request()
func (ip ImageProxy) Draw() {
	ip.GetImage().Draw()
}
