package assembly

import (
	"fmt"
	"strconv"
	"unicode/utf8"

	"github.com/kshamiev/urok/excel/core"
	"github.com/kshamiev/urok/excel/typs"
)

type InvoiceTC struct {
	Number  string       `json:"number"`
	Cargos  []typs.Cargo `json:"cargos"`
	Comment string       `json:"comment"`

	Initiator   typs.Contact       `json:"initiator"`
	From        []typs.Participant `json:"from"`
	To          []typs.Participant `json:"to"`
	PackageList []typs.PackageYarg `json:"package_list"`
	Extra       typs.Extra         `json:"extra"`
	Type        string             `json:"type"`
}

func (inv InvoiceTC) HeaderMain(b *core.Builder) (err error) {
	if err := b.HeaderMain("B", "S", b.Row, b.Row).Height(21).
		Value("Заявка на организацию транспортно-экспедиционного обслуживания № " + inv.Number); err != nil {
		return err
	}
	b.Row++
	return nil
}

func (inv InvoiceTC) InitiatorA(b *core.Builder) error {
	if err := b.Header("B", "S", b.Row, b.Row).Height(18).Value("Инициатор"); err != nil {
		return err
	}
	b.Row++
	if err := b.HeaderSub("B", "B", b.Row, b.Row).Height(16).Value("ФИО"); err != nil {
		return err
	}
	if err := b.Data("C", "H", b.Row, b.Row).Value("${fullname}"); err != nil {
		return err
	}
	if err := b.HeaderSub("I", "I", b.Row, b.Row).Value("Тел"); err != nil {
		return err
	}
	if err := b.Data("J", "M", b.Row, b.Row).Value("${phone}"); err != nil {
		return err
	}
	if err := b.HeaderSub("N", "N", b.Row, b.Row).Value("E-mail"); err != nil {
		return err
	}
	if err := b.Data("O", "S", b.Row, b.Row).Value("${email}"); err != nil {
		return err
	}
	b.Row += 2
	return nil
}

func (inv InvoiceTC) FromA(b *core.Builder) error {
	return inv.fromToA(b, false)
}

func (inv InvoiceTC) ToA(b *core.Builder) error {
	return inv.fromToA(b, true)
}

func (inv InvoiceTC) fromToA(b *core.Builder, flag bool) error {
	actor := "Грузоотправитель"
	date := "Дата отправления"
	if flag {
		actor = "Грузополучатель"
		date = "Дата получения"
	}
	if err := b.Header("B", "D", b.Row, b.Row).Height(18).Value(actor); err != nil {
		return err
	}
	if err := b.Data("E", "O", b.Row, b.Row).Value("${company}"); err != nil {
		return err
	}
	if err := b.Header("P", "Q", b.Row, b.Row).Value(date); err != nil {
		return err
	}
	if err := b.Data("R", "S", b.Row, b.Row).Value("${date}"); err != nil {
		return err
	}
	b.Row++
	if err := b.Header("B", "D", b.Row, b.Row+1).Value("Пункт отправления"); err != nil {
		return err
	}
	if err := b.HeaderSub("E", "J", b.Row, b.Row).Height(16).Value("Область/Республика/Край"); err != nil {
		return err
	}
	if err := b.HeaderSub("K", "M", b.Row, b.Row).Value("Населенный пункт"); err != nil {
		return err
	}
	if err := b.HeaderSub("N", "S", b.Row, b.Row).Value("Улица, дом, корп."); err != nil {
		return err
	}
	b.Row++
	if err := b.Data("E", "J", b.Row, b.Row).Height(16).Value("${province}"); err != nil {
		return err
	}
	if err := b.Data("K", "M", b.Row, b.Row).Value("${city}"); err != nil {
		return err
	}
	if err := b.Data("N", "S", b.Row, b.Row).Value("${house}"); err != nil {
		return err
	}
	b.Row++
	if err := b.Header("B", "D", b.Row, b.Row+1).Value("Ответственное лицо грузополучателя:"); err != nil {
		return err
	}
	if err := b.HeaderSub("E", "J", b.Row, b.Row).Height(16).Value("ФИО"); err != nil {
		return err
	}
	if err := b.HeaderSub("K", "M", b.Row, b.Row).Value("Телефон"); err != nil {
		return err
	}
	if err := b.HeaderSub("N", "S", b.Row, b.Row).Value("e-mail"); err != nil {
		return err
	}
	b.Row++
	if err := b.Data("E", "J", b.Row, b.Row).Height(16).Value("${fullname}"); err != nil {
		return err
	}
	if err := b.Data("K", "M", b.Row, b.Row).Value("${phone}"); err != nil {
		return err
	}
	if err := b.Data("N", "S", b.Row, b.Row).Value("${email}"); err != nil {
		return err
	}
	b.Row += 3
	return nil
}

func (inv InvoiceTC) CargosA(b *core.Builder) error {
	var formula string
	// header
	if err := b.Header("B", "S", b.Row, b.Row).Height(18).Value("Информация о грузе"); err != nil {
		return err
	}
	b.Row++
	// header sub
	if err := b.HeaderSub("B", "C", b.Row, b.Row).Height(16).Value("Код"); err != nil {
		return err
	}
	if err := b.HeaderSub("D", "I", b.Row, b.Row).Value("Наименование"); err != nil {
		return err
	}
	if err := b.HeaderSub("J", "J", b.Row, b.Row).Width(13).Value("Вес за ед."); err != nil {
		return err
	}
	if err := b.HeaderSub("K", "M", b.Row, b.Row).Width(16).Value("Габариты (длина/ширина/высота) за ед, м"); err != nil {
		return err
	}
	if err := b.HeaderSub("N", "N", b.Row, b.Row).Width(12).Value("Кол-во"); err != nil {
		return err
	}
	if err := b.HeaderSub("O", "O", b.Row, b.Row).Width(14).Value("Объем, м3"); err != nil {
		return err
	}
	if err := b.HeaderSub("P", "Q", b.Row, b.Row).Width(14).Value("Объемный вес"); err != nil {
		return err
	}
	if err := b.HeaderSub("R", "R", b.Row, b.Row).Width(14).Value("Стоимость"); err != nil {
		return err
	}
	if err := b.HeaderSub("S", "S", b.Row, b.Row).Width(14).Value("Итого вес"); err != nil {
		return err
	}
	b.Row++
	// data
	rStart := strconv.Itoa(b.Row)
	for i := range inv.Cargos {
		r := strconv.Itoa(b.Row)
		if err := b.Data("B", "C", b.Row, b.Row).Height(16).Value(inv.Cargos[i].ID); err != nil {
			return err
		}
		if err := b.Data("D", "I", b.Row, b.Row).Value(inv.Cargos[i].Name); err != nil {
			return err
		}
		if err := b.Data("J", "J", b.Row, b.Row).Value(inv.Cargos[i].Weight); err != nil {
			return err
		}
		if err := b.Data("K", "K", b.Row, b.Row).Value(inv.Cargos[i].Length); err != nil {
			return err
		}
		if err := b.Data("L", "L", b.Row, b.Row).Value(inv.Cargos[i].Width); err != nil {
			return err
		}
		if err := b.Data("M", "M", b.Row, b.Row).Value(inv.Cargos[i].Height); err != nil {
			return err
		}
		if err := b.Data("N", "N", b.Row, b.Row).Value(inv.Cargos[i].Amount); err != nil {
			return err
		}
		formula = fmt.Sprintf("=K%s*L%s*M%s*N%s", r, r, r, r)
		if err := b.Formula("O", "O", b.Row, b.Row).Formula(formula); err != nil {
			return err

		}
		formula = fmt.Sprintf("=O%s*1000/6*N%s", r, r)
		if err := b.Formula("P", "P", b.Row, b.Row).Formula(formula); err != nil {
			return err
		}
		formula = fmt.Sprintf("=O%s*250", r)
		if err := b.Formula("Q", "Q", b.Row, b.Row).Formula(formula); err != nil {
			return err
		}
		if err := b.Data("R", "R", b.Row, b.Row).Value(inv.Cargos[i].Summ); err != nil {
			return err
		}
		formula = fmt.Sprintf("=N%s*J%s", r, r)
		if err := b.Formula("S", "S", b.Row, b.Row).Formula(formula); err != nil {
			return err
		}
		b.Row++
	}
	rStop := strconv.Itoa(b.Row - 1)
	// footer
	if err := b.HeaderSub("B", "M", b.Row, b.Row).Height(16).Value("итого"); err != nil {
		return err
	}
	formula = fmt.Sprintf("=SUM(N%s:N%s)", rStart, rStop)
	if err := b.Formula("N", "N", b.Row, b.Row).Formula(formula); err != nil {
		return err
	}
	formula = fmt.Sprintf("=SUM(O%s:O%s)", rStart, rStop)
	if err := b.Formula("O", "O", b.Row, b.Row).Formula(formula); err != nil {
		return err
	}
	formula = fmt.Sprintf("=IF(AND(SUM(Q%s:Q%s)>=0,SUM(Q%s:Q%s)<=0.5),0.5,ROUND(SUM(P%s:P%s),0))",
		rStart, rStop, rStart, rStop, rStart, rStop)
	if err := b.Formula("P", "P", b.Row, b.Row).Formula(formula); err != nil {
		return err
	}
	formula = fmt.Sprintf("=ROUND(SUM(Q%s:Q%s),0)", rStart, rStop)
	if err := b.Formula("Q", "Q", b.Row, b.Row).Formula(formula); err != nil {
		return err
	}
	formula = fmt.Sprintf("=SUM(R%s:R%s)", rStart, rStop)
	if err := b.Formula("R", "R", b.Row, b.Row).Formula(formula); err != nil {
		return err
	}
	formula = fmt.Sprintf("=SUM(S%s:S%s)", rStart, rStop)
	if err := b.Formula("S", "S", b.Row, b.Row).Formula(formula); err != nil {
		return err
	}
	b.Row += 2
	return nil
}

func (inv InvoiceTC) HeaderAdvanced(b *core.Builder) (err error) {
	if err := b.Cell("B", "S", b.Row, b.Row).Height(21).
		Value("Дополнительные условия перевозки"); err != nil {
		return err
	}
	b.Row += 2
	return nil
}

func (inv InvoiceTC) CommentA(b *core.Builder) (err error) {
	// header
	if err := b.Header("B", "S", b.Row, b.Row).Height(16).Value("Комментарий"); err != nil {
		return err
	}
	b.Row++
	// data
	h := float64((utf8.RuneCountInString(inv.Comment)/200 + 1) * 14)
	if err := b.Data("B", "S", b.Row, b.Row).Height(h).Value(inv.Comment); err != nil {
		return err
	}
	b.Row += 2
	return nil
}
