package assembly

import (
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

func (inv InvoiceTC) HeaderMain(comp *core.Builder, value interface{}) (err error) {
	if err := comp.SetCellStHeaderMain("B", "S", value); err != nil {
		return err
	}
	comp.SetRowMove(1)
	return nil
}

func (inv InvoiceTC) CargosA(comp *core.Builder) (err error) {
	// header
	if err := comp.SetCellStHeader("B", "S", "Информация о грузе"); err != nil {
		return err
	}
	comp.SetRowMove(1)
	// header sub
	if err := comp.SetCellStHeaderSub("B", "C", "Код"); err != nil {
		return err
	}
	if err := comp.SetCellStHeaderSub("D", "I", "Наименование"); err != nil {
		return err
	}
	_ = comp.SetColWidth("J", "J", 13)
	if err := comp.SetCellStHeaderSub("J", "J", "Вес за ед."); err != nil {
		return err
	}
	_ = comp.SetColWidth("K", "M", 16)
	if err := comp.SetCellStHeaderSub("K", "M", "Габариты (длина/ширина/высота) за ед, м"); err != nil {
		return err
	}
	_ = comp.SetColWidth("N", "N", 12)
	if err := comp.SetCellStHeaderSub("N", "N", "Кол-во"); err != nil {
		return err
	}
	_ = comp.SetColWidth("O", "O", 14)
	if err := comp.SetCellStHeaderSub("O", "O", "Объем, м3"); err != nil {
		return err
	}
	_ = comp.SetColWidth("P", "Q", 14)
	if err := comp.SetCellStHeaderSub("P", "Q", "Объемный вес"); err != nil {
		return err
	}
	_ = comp.SetColWidth("R", "R", 14)
	if err := comp.SetCellStHeaderSub("R", "R", "Стоимость"); err != nil {
		return err
	}
	_ = comp.SetColWidth("S", "S", 14)
	if err := comp.SetCellStHeaderSub("S", "S", "Итого вес"); err != nil {
		return err
	}
	comp.SetRowMove(1)
	// data
	for i := range inv.Cargos {
		if err := comp.SetCellStData("B", "C", inv.Cargos[i].ID); err != nil {
			return err
		}
		if err := comp.SetCellStData("D", "I", inv.Cargos[i].Name); err != nil {
			return err
		}
		if err := comp.SetCellStData("J", "J", inv.Cargos[i].Weight); err != nil {
			return err
		}
		if err := comp.SetCellStData("K", "K", inv.Cargos[i].Length); err != nil {
			return err
		}
		if err := comp.SetCellStData("L", "L", inv.Cargos[i].Width); err != nil {
			return err
		}
		if err := comp.SetCellStData("M", "M", inv.Cargos[i].Height); err != nil {
			return err
		}
		if err := comp.SetCellStData("N", "N", inv.Cargos[i].Amount); err != nil {
			return err
		}
		if err := comp.SetCellStData("O", "O", "F 1"); err != nil {
			return err
		}
		if err := comp.SetCellStData("P", "P", "F 2"); err != nil {
			return err
		}
		if err := comp.SetCellStData("Q", "Q", "F 3"); err != nil {
			return err
		}
		if err := comp.SetCellStData("R", "R", inv.Cargos[i].Summ); err != nil {
			return err
		}
		if err := comp.SetCellStData("S", "S", "F 4"); err != nil {
			return err
		}
		comp.SetRowMove(1)
	}
	// footer
	if err := comp.SetCellStFooter("B", "M", "итого"); err != nil {
		return err
	}
	if err := comp.SetCellStFooter("N", "N", "FF 0"); err != nil {
		return err
	}
	if err := comp.SetCellStFooter("O", "O", "FF 1"); err != nil {
		return err
	}
	if err := comp.SetCellStFooter("P", "P", "FF 2"); err != nil {
		return err
	}
	if err := comp.SetCellStFooter("Q", "Q", "FF 3"); err != nil {
		return err
	}
	if err := comp.SetCellStFooter("R", "R", "FF 4"); err != nil {
		return err
	}
	if err := comp.SetCellStFooter("S", "S", "FF 5"); err != nil {
		return err
	}
	comp.SetRowMove(1)
	return nil
}

func (inv InvoiceTC) CommentA(comp *core.Builder) (err error) {
	// header
	if err := comp.SetCellStHeader("B", "S", "Комментарий"); err != nil {
		return err
	}
	comp.SetRowMove(1)
	// data
	h := float64((utf8.RuneCountInString(inv.Comment)/200 + 1) * 14)
	if err := comp.SetRowHeight(h); err != nil {
		return err
	}
	if err := comp.SetCellSt("B", "S", inv.Comment, comp.StData); err != nil {
		return err
	}
	comp.SetRowMove(1)
	return nil
}
