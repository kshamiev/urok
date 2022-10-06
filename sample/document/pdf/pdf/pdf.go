package pdf

import (
	"io"
	"log"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type PDFGenerator interface {
	GenerateByTemplate(io.Reader, io.Writer) error
}

type pdfGeneratorService struct {
	pdfg *wkhtmltopdf.PDFGenerator
	cfg  Config
}

func NewPDFGenerator(cfg Config) (PDFGenerator, error) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal("pdf.NewPDFGenerator wkhtmltopdf.NewPDFGenerator")
		return nil, err
	}

	return &pdfGeneratorService{
		pdfg: pdfg,
		cfg:  cfg,
	}, nil
}

func (s *pdfGeneratorService) GenerateByTemplate(inPDF io.Reader, outPDF io.Writer) error {
	s.pdfg.Dpi.Set(300)
	s.pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)

	s.pdfg.MarginLeft.Set(s.cfg.PDFPoAMarginLeft)
	s.pdfg.MarginRight.Set(s.cfg.PDFPoAMarginRight)
	s.pdfg.MarginTop.Set(s.cfg.PDFPoAMarginTop)
	s.pdfg.MarginBottom.Set(s.cfg.PDFPoAMarginBottom)

	page := wkhtmltopdf.NewPageReader(inPDF)
	s.pdfg.AddPage(page)

	err := s.pdfg.Create()
	if err != nil {
		log.Fatal("pdf.GenerateByTemplate pdfg.Create()")
		return err
	}

	if _, err = outPDF.Write(s.pdfg.Bytes()); err != nil {
		log.Fatal("pdf.GenerateByTemplate outPDF.Write")
		return err
	}

	// s.pdfg.ResetPages()
	if s.pdfg, err = wkhtmltopdf.NewPDFGenerator(); err != nil {
		log.Fatal("pdf.GenerateByTemplate wkhtmltopdf.NewPDFGenerator")
		return err
	}

	return nil
}
