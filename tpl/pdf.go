package tpl

import (
	"bytes"
	"io"
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/pdfcpu/pdfcpu/pkg/cli"
)

type PDFGenerator struct {
	configF func(pdfg *wkhtmltopdf.PDFGenerator)
}

func NewPDFGenerator() *PDFGenerator {
	if os.Getenv("WKHTMLTOPDF_PATH") == "" {
		_ = os.Setenv("WKHTMLTOPDF_PATH", "/bin")
	}
	return &PDFGenerator{}
}

func (s *PDFGenerator) config(pdfg *wkhtmltopdf.PDFGenerator) {
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)

	pdfg.MarginLeft.Set(10)
	pdfg.MarginRight.Set(10)
	pdfg.MarginTop.Set(10)
	pdfg.MarginBottom.Set(10)
}

func (s *PDFGenerator) Generate(inHtml io.Reader, outPDF io.Writer) error {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return err
	}
	s.config(pdfg)

	p := wkhtmltopdf.NewPageReader(inHtml)
	p.EnableLocalFileAccess.Set(true)
	p.EnableForms.Set(true)
	pdfg.AddPage(p)

	if err := pdfg.Create(); err != nil {
		return err
	}
	_, err = outPDF.Write(pdfg.Bytes())
	return err
}

func (s *PDFGenerator) GenerateFile(inHtml io.Reader, pathPDF string) error {
	outPDF := &bytes.Buffer{}
	if err := s.Generate(inHtml, outPDF); err != nil {
		return err
	}
	return os.WriteFile(pathPDF, outPDF.Bytes(), 0o600)
}

func (s *PDFGenerator) GenerateToJson(inHtml io.Reader) ([]byte, error) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	s.config(pdfg)

	pdfg.AddPage(wkhtmltopdf.NewPageReader(inHtml))

	return pdfg.ToJSON()
}

func (s *PDFGenerator) GenerateFromJson(jb []byte, outPDF io.Writer) error {
	pdfg, err := wkhtmltopdf.NewPDFGeneratorFromJSON(bytes.NewReader(jb))
	if err != nil {
		return err
	}

	err = pdfg.Create()
	if err != nil {
		return err
	}

	_, err = outPDF.Write(pdfg.Bytes())
	return err
}

func (s *PDFGenerator) AddAttachment(filename string, files []string) error {
	cmd := cli.AddAttachmentsCommand(filename, "", files, nil)
	if _, err := cli.Process(cmd); err != nil {
		return err
	}
	return nil
}

func (s *PDFGenerator) DelAttachment(filename string, files []string) error {
	cmd := cli.RemoveAttachmentsCommand(filename, "", files, nil)
	if _, err := cli.Process(cmd); err != nil {
		return err
	}
	return nil
}
