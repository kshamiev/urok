package typs

func NewInvoiceTC(number string) InvoiceTC {
	return InvoiceTC{
		Number: number,
		Initiator: Contact{
			FullName: "Капитан очевидность",
			Phone:    "-5 (000) 111-11-11",
			Email:    "big@fire.ru",
		},
		From: Participant{
			Contact: Contact{
				FullName: "Шариков Полиграф Полиграфович",
				Phone:    "-1 (000) 999-99-99",
				Email:    "koshmar@boloto.ru",
			},
			Company: "Три пескаря",
			Date:    "2010-10-10",
			Destination: Destination{
				Province: "Болото",
				City:     "Лягушатник",
				House:    "5",
			},
		},
		To: Participant{
			Contact: Contact{
				FullName: "Шаша Бесподобный",
				Phone:    "-2 (000) 888-88-88",
				Email:    "razlojenie@ujas.ru",
			},
			Company: "Колхоз 30 лет без урожая",
			Date:    "2010-10-20",
			Destination: Destination{
				Province: "Топь",
				City:     "Трясина",
				House:    "23",
			},
		},
		Cargos: []Cargo{
			{
				ID:     "201",
				Name:   "Ловушка для мух",
				Weight: 10,
				Length: 1.4,
				Width:  1.5,
				Height: 1.4,
				Amount: 4,
				Summ:   950,
			},
			{
				ID:     "202",
				Name:   "Мышеловка для крыс",
				Weight: 12,
				Length: 1.5,
				Width:  1.3,
				Height: 1.1,
				Amount: 4,
				Summ:   345,
			},
			{
				ID:     "203",
				Name:   "Сачок для бабочек",
				Weight: 14,
				Length: 1.9,
				Width:  1.8,
				Height: 1.7,
				Amount: 8,
				Summ:   1000,
			},
			{
				ID:     "204",
				Name:   "Сморчок на палочке",
				Weight: 16,
				Length: 2,
				Width:  1.5,
				Height: 1.3,
				Amount: 10,
				Summ:   756,
			},
			{
				ID:     "205",
				Name:   "Суслик засушливый",
				Weight: 18,
				Length: 1,
				Width:  1.6,
				Height: 1.4,
				Amount: 12,
				Summ:   802,
			},
			{
				ID:     "206",
				Name:   "Мухомор обыкновенный",
				Weight: 20,
				Length: 1.1,
				Width:  1.2,
				Height: 1.3,
				Amount: 6,
				Summ:   456,
			},
		},
		PackageList: []PackageYarg{
			{Name: "Коробка", Count: "34"},
			{Name: "Упаковка", Count: "567"},
			{Name: "Прыщи пузырчатые", Count: "1289"},
		},
		Comment: "Предоставляет функцию для установки имени листа по заданным именам старого и нового листа. В заголовке листа допускается не более 31 символа, и эта функция изменяет только имя листа и не обновляет имя листа в формуле или ссылке, связанной с ячейкой. Таким образом, может быть ошибка формулы проблемы или отсутствует ссылка. Предоставляет функцию для установки имени листа по заданным именам старого и нового листа. В заголовке листа допускается не более 31 символа, и эта функция изменяет только имя листа и не обновляет имя листа в формуле или ссылке, связанной с ячейкой. Таким образом, может быть ошибка формулы проблемы или отсутствует ссылка.",
	}
}

type InvoiceTC struct {
	Number      string        `json:"number"`
	Initiator   Contact       `json:"initiator"`
	From        Participant   `json:"from"`
	To          Participant   `json:"to"`
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
