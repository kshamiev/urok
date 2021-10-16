package typs

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
