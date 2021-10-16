package assembly

import (
	"fmt"
	"strconv"
	"unicode/utf8"

	"github.com/xuri/excelize/v2"

	"github.com/kshamiev/urok/sample/excel/excelb"
	"github.com/kshamiev/urok/sample/excel/typs"
)

type InvoiceTC struct {
	Number      string             `json:"number"`
	Initiator   typs.Contact       `json:"initiator"`
	From        []typs.Participant `json:"from"`
	To          []typs.Participant `json:"to"`
	Cargos      []typs.Cargo       `json:"cargos"`
	PackageList []typs.PackageYarg `json:"package_list"`
	Comment     string             `json:"comment"`
	Extra       typs.Extra         `json:"extra"`
	Type        string             `json:"type"`
}

func InvoiceTCTrucking(tplFilePath string, data []InvoiceTC) (*excelize.File, error) {
	fp, err := excelize.OpenFile(tplFilePath)
	if err != nil {
		return nil, err
	}
	sheetList := fp.GetSheetList()
	defer func() {
		for i := range sheetList {
			fp.DeleteSheet(sheetList[i])
		}
	}()

	build, err := excelb.NewBuilder(fp)
	if err != nil {
		return nil, err
	}

	build.NewSheet("main")

	for _, inv := range data {
		build.NewSheet(inv.Number)
		build.Row = 1
		if err := inv.HeaderMain(build); err != nil {
			return nil, err
		}
		if err := inv.InitiatorA(build); err != nil {
			return nil, err
		}
		if err := inv.FromA(build); err != nil {
			return nil, err
		}
		if err := inv.ToA(build); err != nil {
			return nil, err
		}
		if err := inv.CargosA(build); err != nil {
			return nil, err
		}
		if err := inv.HeaderAdditional(build); err != nil {
			return nil, err
		}
		if err := inv.Additional(build); err != nil {
			return nil, err
		}
		if err := inv.CommentA(build); err != nil {
			return nil, err
		}
	}
	return fp, nil
}

func (inv InvoiceTC) HeaderMain(b *excelb.Builder) (err error) {
	if err := b.Header1("B", "S", b.Row, b.Row).Height(21).
		Value("Заявка на организацию транспортно-экспедиционного обслуживания № " + inv.Number); err != nil {
		return err
	}
	b.Row++
	return nil
}

func (inv InvoiceTC) InitiatorA(b *excelb.Builder) error {
	if err := b.Header3("B", "S", b.Row, b.Row).Height(18).Value("Инициатор"); err != nil {
		return err
	}
	b.Row++
	if err := b.Header4("B", "B", b.Row, b.Row).Height(16).Value("ФИО"); err != nil {
		return err
	}
	if err := b.Content1("C", "H", b.Row, b.Row).Value("${fullname}"); err != nil {
		return err
	}
	if err := b.Header4("I", "I", b.Row, b.Row).Value("Тел"); err != nil {
		return err
	}
	if err := b.Content1("J", "M", b.Row, b.Row).Value("${phone}"); err != nil {
		return err
	}
	if err := b.Header4("N", "N", b.Row, b.Row).Value("E-mail"); err != nil {
		return err
	}
	if err := b.Content1("O", "S", b.Row, b.Row).Value("${email}"); err != nil {
		return err
	}
	b.Row += 2
	return nil
}

func (inv InvoiceTC) FromA(b *excelb.Builder) error {
	return inv.fromToA(b, false)
}

func (inv InvoiceTC) ToA(b *excelb.Builder) error {
	return inv.fromToA(b, true)
}

func (inv InvoiceTC) fromToA(b *excelb.Builder, flag bool) error {
	actor := "Грузоотправитель"
	date := "Дата отправления"
	if flag {
		actor = "Грузополучатель"
		date = "Дата получения"
	}
	if err := b.Header3("B", "D", b.Row, b.Row).Height(18).Value(actor); err != nil {
		return err
	}
	if err := b.Content1("E", "O", b.Row, b.Row).Value("${company}"); err != nil {
		return err
	}
	if err := b.Header3("P", "Q", b.Row, b.Row).Value(date); err != nil {
		return err
	}
	if err := b.Content1("R", "S", b.Row, b.Row).Value("${date}"); err != nil {
		return err
	}
	b.Row++
	if err := b.Header3("B", "D", b.Row, b.Row+1).Value("Пункт отправления"); err != nil {
		return err
	}
	if err := b.Header4("E", "J", b.Row, b.Row).Height(16).Value("Область/Республика/Край"); err != nil {
		return err
	}
	if err := b.Header4("K", "M", b.Row, b.Row).Value("Населенный пункт"); err != nil {
		return err
	}
	if err := b.Header4("N", "S", b.Row, b.Row).Value("Улица, дом, корп."); err != nil {
		return err
	}
	b.Row++
	if err := b.Content1("E", "J", b.Row, b.Row).Height(16).Value("${province}"); err != nil {
		return err
	}
	if err := b.Content1("K", "M", b.Row, b.Row).Value("${city}"); err != nil {
		return err
	}
	if err := b.Content1("N", "S", b.Row, b.Row).Value("${house}"); err != nil {
		return err
	}
	b.Row++
	if err := b.Header3("B", "D", b.Row, b.Row+1).Value("Ответственное лицо грузополучателя:"); err != nil {
		return err
	}
	if err := b.Header4("E", "J", b.Row, b.Row).Height(16).Value("ФИО"); err != nil {
		return err
	}
	if err := b.Header4("K", "M", b.Row, b.Row).Value("Телефон"); err != nil {
		return err
	}
	if err := b.Header4("N", "S", b.Row, b.Row).Value("e-mail"); err != nil {
		return err
	}
	b.Row++
	if err := b.Content1("E", "J", b.Row, b.Row).Height(16).Value("${fullname}"); err != nil {
		return err
	}
	if err := b.Content1("K", "M", b.Row, b.Row).Value("${phone}"); err != nil {
		return err
	}
	if err := b.Content1("N", "S", b.Row, b.Row).Value("${email}"); err != nil {
		return err
	}
	b.Row += 3
	return nil
}

func (inv InvoiceTC) CargosA(b *excelb.Builder) error {
	var formula string
	// header
	if err := b.Header3("B", "S", b.Row, b.Row).Height(18).Value("Информация о грузе"); err != nil {
		return err
	}
	b.Row++
	// header sub
	if err := b.Header4("B", "C", b.Row, b.Row).Height(16).Value("Код"); err != nil {
		return err
	}
	if err := b.Header4("D", "I", b.Row, b.Row).Value("Наименование"); err != nil {
		return err
	}
	if err := b.Header4("J", "J", b.Row, b.Row).Width(13).Value("Вес за ед."); err != nil {
		return err
	}
	if err := b.Header4("K", "M", b.Row, b.Row).Width(16).Value("Габариты (длина/ширина/высота) за ед, м"); err != nil {
		return err
	}
	if err := b.Header4("N", "N", b.Row, b.Row).Width(12).Value("Кол-во"); err != nil {
		return err
	}
	if err := b.Header4("O", "O", b.Row, b.Row).Width(14).Value("Объем, м3"); err != nil {
		return err
	}
	if err := b.Header4("P", "Q", b.Row, b.Row).Width(14).Value("Объемный вес"); err != nil {
		return err
	}
	if err := b.Header4("R", "R", b.Row, b.Row).Width(14).Value("Стоимость"); err != nil {
		return err
	}
	if err := b.Header4("S", "S", b.Row, b.Row).Width(14).Value("Итого вес"); err != nil {
		return err
	}
	b.Row++
	// data
	rStart := strconv.Itoa(b.Row)
	for i := range inv.Cargos {
		r := strconv.Itoa(b.Row)
		if err := b.Content1("B", "C", b.Row, b.Row).Height(16).Value(inv.Cargos[i].ID); err != nil {
			return err
		}
		if err := b.Content1("D", "I", b.Row, b.Row).Value(inv.Cargos[i].Name); err != nil {
			return err
		}
		if err := b.Content1("J", "J", b.Row, b.Row).Value(inv.Cargos[i].Weight); err != nil {
			return err
		}
		if err := b.Content1("K", "K", b.Row, b.Row).Value(inv.Cargos[i].Length); err != nil {
			return err
		}
		if err := b.Content1("L", "L", b.Row, b.Row).Value(inv.Cargos[i].Width); err != nil {
			return err
		}
		if err := b.Content1("M", "M", b.Row, b.Row).Value(inv.Cargos[i].Height); err != nil {
			return err
		}
		if err := b.Content1("N", "N", b.Row, b.Row).Value(inv.Cargos[i].Amount); err != nil {
			return err
		}
		formula = fmt.Sprintf("=K%s*L%s*M%s*N%s", r, r, r, r)
		if err := b.Formula1("O", "O", b.Row, b.Row).Formula(formula); err != nil {
			return err

		}
		formula = fmt.Sprintf("=O%s*1000/6*N%s", r, r)
		if err := b.Formula1("P", "P", b.Row, b.Row).Formula(formula); err != nil {
			return err
		}
		formula = fmt.Sprintf("=O%s*250", r)
		if err := b.Formula1("Q", "Q", b.Row, b.Row).Formula(formula); err != nil {
			return err
		}
		if err := b.Content1("R", "R", b.Row, b.Row).Value(inv.Cargos[i].Summ); err != nil {
			return err
		}
		formula = fmt.Sprintf("=N%s*J%s", r, r)
		if err := b.Formula1("S", "S", b.Row, b.Row).Formula(formula); err != nil {
			return err
		}
		b.Row++
	}
	rStop := strconv.Itoa(b.Row - 1)
	// footer
	if err := b.Footer1("B", "M", b.Row, b.Row).Height(16).Value("итого"); err != nil {
		return err
	}
	formula = fmt.Sprintf("=SUM(N%s:N%s)", rStart, rStop)
	if err := b.Formula1("N", "N", b.Row, b.Row).Formula(formula); err != nil {
		return err
	}
	formula = fmt.Sprintf("=SUM(O%s:O%s)", rStart, rStop)
	if err := b.Formula1("O", "O", b.Row, b.Row).Formula(formula); err != nil {
		return err
	}
	formula = fmt.Sprintf("=IF(AND(SUM(Q%s:Q%s)>=0,SUM(Q%s:Q%s)<=0.5),0.5,ROUND(SUM(P%s:P%s),0))",
		rStart, rStop, rStart, rStop, rStart, rStop)
	if err := b.Formula1("P", "P", b.Row, b.Row).Formula(formula); err != nil {
		return err
	}
	formula = fmt.Sprintf("=ROUND(SUM(Q%s:Q%s),0)", rStart, rStop)
	if err := b.Formula1("Q", "Q", b.Row, b.Row).Formula(formula); err != nil {
		return err
	}
	formula = fmt.Sprintf("=SUM(R%s:R%s)", rStart, rStop)
	if err := b.Formula1("R", "R", b.Row, b.Row).Formula(formula); err != nil {
		return err
	}
	formula = fmt.Sprintf("=SUM(S%s:S%s)", rStart, rStop)
	if err := b.Formula1("S", "S", b.Row, b.Row).Formula(formula); err != nil {
		return err
	}
	b.Row += 2
	return nil
}

func (inv InvoiceTC) HeaderAdditional(b *excelb.Builder) (err error) {
	if err := b.Header2("B", "S", b.Row, b.Row).Height(21).
		Value("Дополнительные условия перевозки"); err != nil {
		return err
	}
	b.Row += 2
	return nil
}

func (inv InvoiceTC) Additional(b *excelb.Builder) (err error) {
	if len(inv.PackageList) > 0 {
		if err := b.Header3("B", "I", b.Row, b.Row).Height(18).Value("Упаковка"); err != nil {
			return err
		}
		if err := b.Header3("J", "K", b.Row, b.Row).Height(18).Value("количество"); err != nil {
			return err
		}
		b.Row++
		for i := range inv.PackageList {
			if err := b.Content1("B", "I", b.Row, b.Row).Height(16).Value(inv.PackageList[i].Name); err != nil {
				return err
			}
			if err := b.Content1("J", "K", b.Row, b.Row).Height(16).Value(inv.PackageList[i].Count); err != nil {
				return err
			}
			b.Row++
		}
	}
	b.Row += 2
	return nil
}

func (inv InvoiceTC) CommentA(b *excelb.Builder) (err error) {
	// header
	if err := b.Header3("B", "S", b.Row, b.Row).Height(16).Value("Комментарий"); err != nil {
		return err
	}
	b.Row++
	// data
	h := float64((utf8.RuneCountInString(inv.Comment)/200 + 1) * 14)
	if err := b.Content1("B", "S", b.Row, b.Row).Height(h).Value(inv.Comment); err != nil {
		return err
	}
	b.Row += 2
	return nil
}
