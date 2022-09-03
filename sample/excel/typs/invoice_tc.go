package typs

func NewInvoiceTC(number string) InvoiceTC {
	return InvoiceTC{
		Number: number,
		Cargos: []Cargo{
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
		PackageList: []PackageYarg{
			{Name: "Коробка", Count: "34"},
			{Name: "Упаковка", Count: "567"},
			{Name: "Прыщи пузырчатые", Count: "1289"},
		},
		Comment: "SetSheetName предоставляет функцию для установки имени листа по заданным именам старого и нового листа. В заголовке листа допускается не более 31 символа, и эта функция изменяет только имя листа и не обновляет имя листа в формуле или ссылке, связанной с ячейкой. Таким образом, может быть ошибка формулы проблемы или отсутствует ссылка. SetSheetName предоставляет функцию для установки имени листа по заданным именам старого и нового листа. В заголовке листа допускается не более 31 символа, и эта функция изменяет только имя листа и не обновляет имя листа в формуле или ссылке, связанной с ячейкой. Таким образом, может быть ошибка формулы проблемы или отсутствует ссылка. ",
	}
}

type InvoiceTC struct {
	Number      string        `json:"number"`
	Initiator   Contact       `json:"initiator"`
	From        []Participant `json:"from"`
	To          []Participant `json:"to"`
	Cargos      []Cargo       `json:"cargos"`
	PackageList []PackageYarg `json:"package_list"`
	Comment     string        `json:"comment"`
	Extra       Extra         `json:"extra"`
	Type        string        `json:"type"`
}

type Cargo struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Weight float64 `json:"weight"`
	Length float64 `json:"length"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Amount int     `json:"amount"`
	Summ   float64 `json:"summ"`
}

type PackageYarg struct {
	Name  string `json:"name"`
	Count string `json:"count"`
}

type Extra struct {
	Loading        string `json:"loading"`
	Loadingdop     string `json:"loadingdop"`
	Crane          string `json:"crane"`
	AerialPlatform string `json:"aerial_platform"`
	Special        string `json:"special"`
	Package        string `json:"package"`
	Packagedop     string `json:"packagedop"`
	Insurance      string `json:"insurance"`
	Fragile        string `json:"fragile"`
	Dispatcher     string `json:"dispatcher"`
	Comment        string `json:"comment"`
}

type Destination struct {
	Province string `json:"province"`
	City     string `json:"city"`
	House    string `json:"house"`
}

type Participant struct {
	Contact     Contact     `json:"contact"`
	Company     string      `json:"company"`
	Date        string      `json:"date"`
	Destination Destination `json:"destination"`
}

type Contact struct {
	FullName string `json:"fullname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}
