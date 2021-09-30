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

	comp, err := core.NewComponent(fp)
	if err != nil {
		return nil, err
	}

	const hMain1 = "Заявка на организацию транспортно-экспедиционного обслуживания № "

	comp.NewSheet("main")

	for _, inv := range data {
		comp.NewSheet(inv.Number)
		if err := inv.HeaderMain(comp, hMain1+inv.Number); err != nil {
			return nil, err
		}
		if err := inv.CargosA(comp); err != nil {
			return nil, err
		}
		comp.SetRowMove(1)
		if err := inv.CommentA(comp); err != nil {
			return nil, err
		}
		comp.SetRowMove(1)
	}
	return fp, nil
}
