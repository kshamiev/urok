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
