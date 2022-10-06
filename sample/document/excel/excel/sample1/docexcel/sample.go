package docexcel

import (
	"fmt"
	"strconv"
	"unicode/utf8"

	"github.com/xuri/excelize/v2"

	excel2 "github.com/kshamiev/urok/sample/document/excel/excel"
	"github.com/kshamiev/urok/sample/document/excel/excel/sample1/typs"
)

type Sample struct {
	tplPath string
	s       styleSheet
}

func NewSample(tplPath string) Sample {
	return Sample{
		tplPath: tplPath,
	}
}

func (doc Sample) Compile(data []typs.InvoiceTC) (*excelize.File, error) {
	bu, err := excel2.NewBuilderFile(doc.tplPath)
	if err != nil {
		return nil, err
	}
	defer func() {
		bu.DeleteStartSheet()
	}()

	for _, inv := range data {
		b := bu.NewSheet(inv.Number)
		if err := doc.headerMain(b, inv); err != nil {
			return nil, err
		}
		if err := doc.initiatorA(b, inv); err != nil {
			return nil, err
		}
		if err := doc.fromA(b, inv); err != nil {
			return nil, err
		}
		if err := doc.toA(b, inv); err != nil {
			return nil, err
		}
		if err := doc.cargosA(b, inv); err != nil {
			return nil, err
		}
		if err := doc.headerAdditional(b, inv); err != nil {
			return nil, err
		}
		if err := doc.additional(b, inv); err != nil {
			return nil, err
		}
		if err := doc.commentA(b, inv); err != nil {
			return nil, err
		}
	}
	return bu.GetFp(), nil
}

func (doc Sample) headerMain(b *excel2.Build, inv typs.InvoiceTC) (err error) {
	doc.s.Head1(b).Cell("B", "S").Height(21).
		Value("Заявка на организацию транспортно-экспедиционного обслуживания № " + inv.Number)
	b.Row++
	return b.Err
}

func (doc Sample) initiatorA(b *excel2.Build, inv typs.InvoiceTC) (err error) {
	doc.s.Head3(b).CellRow("B", b.Row, "S", b.Row).Height(18).Value("Инициатор")
	b.Row++
	doc.s.Head4(b).Cell("B", "B").Height(16).Value("ФИО")
	doc.s.Body1(b).Cell("C", "H").Value(inv.Initiator.FullName)
	doc.s.Head4(b).Cell("I", "I").Value("Тел")
	doc.s.Body1(b).Cell("J", "M").Value(inv.Initiator.Phone)
	doc.s.Head4(b).Cell("N", "N").Value("E-mail")
	doc.s.Body1(b).Cell("O", "S").Value(inv.Initiator.Email)
	b.Row += 2
	return b.Err
}

func (doc Sample) fromA(b *excel2.Build, inv typs.InvoiceTC) (err error) {
	return doc.fromToA(b, inv.From, false)
}

func (doc Sample) toA(b *excel2.Build, inv typs.InvoiceTC) (err error) {
	return doc.fromToA(b, inv.To, true)
}

func (doc Sample) fromToA(b *excel2.Build, inv typs.Participant, flag bool) (err error) {
	actor := "Грузоотправитель"
	date := "Дата отправления"
	point := "Пункт отправления"
	otvet := "Ответственное лицо грузоотправителя:"
	if flag {
		actor = "Грузополучатель"
		date = "Дата получения"
		point = "Пункт получения"
		otvet = "Ответственное лицо грузополучателя:"
	}
	doc.s.Head3(b).Cell("B", "D").Height(18).Value(actor)
	doc.s.Body1(b).Cell("E", "O").Value(inv.Company)
	doc.s.Head3(b).Cell("P", "Q").Value(date)
	doc.s.Body1(b).Cell("R", "S").Value(inv.Date)
	b.Row++
	doc.s.Head3(b).CellRow("B", b.Row, "D", b.Row+1).Value(point)
	doc.s.Head4(b).Cell("E", "J").Height(16).Value("Область/Республика/Край")
	doc.s.Head4(b).Cell("K", "M").Value("Населенный пункт")
	doc.s.Head4(b).Cell("N", "S").Value("Улица, дом, корп.")
	b.Row++
	doc.s.Body1(b).Cell("E", "J").Height(16).Value(inv.Destination.Province)
	doc.s.Body1(b).Cell("K", "M").Value(inv.Destination.City)
	doc.s.Body1(b).Cell("N", "S").Value(inv.Destination.House)
	b.Row++
	doc.s.Head3(b).CellRow("B", b.Row, "D", b.Row+1).Value(otvet)
	doc.s.Head4(b).Cell("E", "J").Height(16).Value("ФИО")
	doc.s.Head4(b).Cell("K", "M").Value("Телефон")
	doc.s.Head4(b).Cell("N", "S").Value("e-mail")
	b.Row++
	doc.s.Body1(b).Cell("E", "J").Height(16).Value(inv.Contact.FullName)
	doc.s.Body1(b).Cell("K", "M").Value(inv.Contact.Phone)
	doc.s.Body1(b).Cell("N", "S").Value(inv.Contact.Email)
	b.Row += 3
	return b.Err
}

func (doc Sample) cargosA(b *excel2.Build, inv typs.InvoiceTC) (err error) {

	var formula string
	// header
	doc.s.Head3(b).Cell("B", "S").Height(18).Value("Информация о грузе")
	b.Row++
	// header sub
	doc.s.Head4(b).Cell("B", "C").Height(16).Value("Код")
	doc.s.Head4(b).Cell("D", "I").Value("Наименование")
	doc.s.Head4(b).Cell("J", "J").Width(13).Value("Вес за ед.")
	doc.s.Head4(b).Cell("K", "M").Width(16).Value("Габариты (длина/ширина/высота) за ед, м")
	doc.s.Head4(b).Cell("N", "N").Width(12).Value("Кол-во")
	doc.s.Head4(b).Cell("O", "O").Width(14).Value("Объем, м3")
	doc.s.Head4(b).Cell("P", "Q").Width(14).Value("Объемный вес")
	doc.s.Head4(b).Cell("R", "R").Width(14).Value("Стоимость")
	doc.s.Head4(b).Cell("S", "S").Width(14).Value("Итого вес")
	b.Row++
	// data
	rStart := strconv.Itoa(b.Row)
	for i := range inv.Cargos {
		r := strconv.Itoa(b.Row)
		doc.s.Body1(b).Cell("B", "C").Height(16).Value(inv.Cargos[i].ID)
		doc.s.Body1(b).Cell("D", "I").Value(inv.Cargos[i].Name)
		doc.s.Body1(b).Cell("J", "J").Value(inv.Cargos[i].Weight)
		doc.s.Body1(b).Cell("K", "K").Value(inv.Cargos[i].Length)
		doc.s.Body1(b).Cell("L", "L").Value(inv.Cargos[i].Width)
		doc.s.Body1(b).Cell("M", "M").Value(inv.Cargos[i].Height)
		doc.s.Body1(b).Cell("N", "N").Value(inv.Cargos[i].Amount)
		formula = fmt.Sprintf("=K%s*L%s*M%s*N%s", r, r, r, r)
		doc.s.Formula1(b).Cell("O", "O").ValueFormula(formula)
		formula = fmt.Sprintf("=O%s*1000/6*N%s", r, r)
		doc.s.Formula1(b).Cell("P", "P").ValueFormula(formula)
		formula = fmt.Sprintf("=O%s*250", r)
		doc.s.Formula1(b).Cell("Q", "Q").ValueFormula(formula)
		doc.s.Body1(b).Cell("R", "R").Value(inv.Cargos[i].Summ)
		formula = fmt.Sprintf("=N%s*J%s", r, r)
		doc.s.Formula1(b).Cell("S", "S").ValueFormula(formula)
		b.Row++
	}
	rStop := strconv.Itoa(b.Row - 1)
	// footer
	doc.s.Formula1(b).Cell("B", "M").Height(16).Value("итого")
	formula = fmt.Sprintf("=SUM(N%s:N%s)", rStart, rStop)
	doc.s.Formula1(b).Cell("N", "N").ValueFormula(formula)
	formula = fmt.Sprintf("=SUM(O%s:O%s)", rStart, rStop)
	doc.s.Formula1(b).Cell("O", "O").ValueFormula(formula)
	formula = fmt.Sprintf("=IF(AND(SUM(Q%s:Q%s)>=0,SUM(Q%s:Q%s)<=0.5),0.5,ROUND(SUM(P%s:P%s),0))",
		rStart, rStop, rStart, rStop, rStart, rStop)
	doc.s.Formula1(b).Cell("P", "P").ValueFormula(formula)
	formula = fmt.Sprintf("=ROUND(SUM(Q%s:Q%s),0)", rStart, rStop)
	doc.s.Formula1(b).Cell("Q", "Q").ValueFormula(formula)
	formula = fmt.Sprintf("=SUM(R%s:R%s)", rStart, rStop)
	doc.s.Formula1(b).Cell("R", "R").ValueFormula(formula)
	formula = fmt.Sprintf("=SUM(S%s:S%s)", rStart, rStop)
	doc.s.Formula1(b).Cell("S", "S").ValueFormula(formula)
	b.Row += 2

	return b.Err
}
func (doc Sample) headerAdditional(b *excel2.Build, inv typs.InvoiceTC) (err error) {
	doc.s.Head2(b).Cell("B", "S").Height(21).Value("Дополнительные условия перевозки")
	b.Row += 2
	return b.Err
}
func (doc Sample) additional(b *excel2.Build, inv typs.InvoiceTC) (err error) {
	if len(inv.PackageList) > 0 {
		doc.s.Head3(b).Cell("B", "I").Height(18).Value("Упаковка")
		doc.s.Head3(b).Cell("J", "K").Height(18).Value("количество")
		b.Row++
		for i := range inv.PackageList {
			doc.s.Body1(b).Cell("B", "I").Height(16).Value(inv.PackageList[i].Name)
			doc.s.Body1(b).Cell("J", "K").Height(16).Value(inv.PackageList[i].Count)
			b.Row++
		}
	}
	b.Row += 2
	return b.Err
}
func (doc Sample) commentA(b *excel2.Build, inv typs.InvoiceTC) (err error) {
	// header
	doc.s.Head3(b).Cell("B", "S").Height(16).Value("Комментарий")
	b.Row++
	// data
	h := float64((utf8.RuneCountInString(inv.Comment)/200 + 1) * 14)
	doc.s.Body1(b).Cell("B", "S").Height(h).Value(inv.Comment)
	b.Row += 2
	return b.Err
}

// ////

type styleSheet struct {
}

func (doc styleSheet) Head1(b *excel2.Build) *excel2.Build {
	return b.Style("C2")
}
func (doc styleSheet) Head2(b *excel2.Build) *excel2.Build {
	return b.Style("C4")
}
func (doc styleSheet) Head3(b *excel2.Build) *excel2.Build {
	return b.Style("C6")
}
func (doc styleSheet) Head4(b *excel2.Build) *excel2.Build {
	return b.Style("C8")
}
func (doc styleSheet) Body1(b *excel2.Build) *excel2.Build {
	return b.Style("C15")
}
func (doc styleSheet) Footer1(b *excel2.Build) *excel2.Build {
	return b.Style("C29")
}
func (doc styleSheet) Formula1(b *excel2.Build) *excel2.Build {
	return b.Style("C22")
}
