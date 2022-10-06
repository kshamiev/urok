package pdf

type Config struct {
	PDFPoAMarginLeft   uint `json:"pdf_po_a_margin_left"`   // Отступ слева в PDF (в mm) для доверенности и соглашения об ЭДО
	PDFPoAMarginRight  uint `json:"pdf_po_a_margin_right"`  // Отступ справа в PDF (в mm) для доверенности и соглашения об ЭДО
	PDFPoAMarginTop    uint `json:"pdf_po_a_margin_top"`    // Отступ сверху в PDF (в mm) для доверенности и соглашения об ЭДО
	PDFPoAMarginBottom uint `json:"pdf_po_a_margin_bottom"` // Отступ снизу в PDF (в mm) для доверенности и соглашения об ЭДО
}
