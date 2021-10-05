package assembly

import "github.com/kshamiev/urok/excel/typs"

func NewInvoiceTCSample(number string) InvoiceTC {
	return InvoiceTC{
		Number: number,
		Cargos: []typs.Cargo{
			{
				ID:     "200",
				Name:   "Ловушка для мух",
				Weight: 20,
				Length: 2,
				Width:  2,
				Height: 2,
				Amount: 10,
				Summ:   1000,
			},
			{
				ID:     "200",
				Name:   "Ловушка для мух",
				Weight: 20,
				Length: 2,
				Width:  2,
				Height: 2,
				Amount: 10,
				Summ:   1000,
			},
			{
				ID:     "200",
				Name:   "Ловушка для мух",
				Weight: 20,
				Length: 2,
				Width:  2,
				Height: 2,
				Amount: 10,
				Summ:   1000,
			},
			{
				ID:     "200",
				Name:   "Ловушка для мух",
				Weight: 20,
				Length: 2,
				Width:  2,
				Height: 2,
				Amount: 10,
				Summ:   1000,
			},
			{
				ID:     "200",
				Name:   "Ловушка для мух",
				Weight: 20,
				Length: 2,
				Width:  2,
				Height: 2,
				Amount: 10,
				Summ:   1000,
			},
			{
				ID:     "200",
				Name:   "Ловушка для мух",
				Weight: 20,
				Length: 2,
				Width:  2,
				Height: 2,
				Amount: 10,
				Summ:   1000,
			},
		},
		PackageList: []typs.PackageYarg{
			{Name: "Коробка", Count: "34"},
			{Name: "Упаковка", Count: "567"},
			{Name: "Прыщи пузырчатые", Count: "1289"},
		},
		Comment: "SetSheetName предоставляет функцию для установки имени листа по заданным именам старого и нового листа. В заголовке листа допускается не более 31 символа, и эта функция изменяет только имя листа и не обновляет имя листа в формуле или ссылке, связанной с ячейкой. Таким образом, может быть ошибка формулы проблемы или отсутствует ссылка. SetSheetName предоставляет функцию для установки имени листа по заданным именам старого и нового листа. В заголовке листа допускается не более 31 символа, и эта функция изменяет только имя листа и не обновляет имя листа в формуле или ссылке, связанной с ячейкой. Таким образом, может быть ошибка формулы проблемы или отсутствует ссылка. ",
	}
}
