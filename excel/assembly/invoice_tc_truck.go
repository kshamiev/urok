package assembly

import (
	"github.com/xuri/excelize/v2"

	"github.com/kshamiev/urok/excel/core"
)

func InvoiceTCTrucking(tplFilePath string, data []InvoiceTC) (*excelize.File, error) {
	fp, err := excelize.OpenFile(tplFilePath)
	if err != nil {
		return nil, err
	}
	// defer fp.DeleteSheet(core.TplList)

	comp, err := core.NewBuilder(fp)
	if err != nil {
		return nil, err
	}

	comp.NewSheet("main")

	for _, inv := range data {
		comp.NewSheet(inv.Number)
		comp.Row = 1
		if err := inv.HeaderMain(comp); err != nil {
			return nil, err
		}
		if err := inv.InitiatorA(comp); err != nil {
			return nil, err
		}
		if err := inv.FromA(comp); err != nil {
			return nil, err
		}
		if err := inv.ToA(comp); err != nil {
			return nil, err
		}
		if err := inv.CargosA(comp); err != nil {
			return nil, err
		}
		if err := inv.HeaderAdvanced(comp); err != nil {
			return nil, err
		}
		if err := inv.CommentA(comp); err != nil {
			return nil, err
		}
	}
	return fp, nil
}
