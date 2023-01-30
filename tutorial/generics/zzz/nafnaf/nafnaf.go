package nafnaf

import (
	"github.com/kshamiev/urok/tutorial/generics/zzz"
)

type NafNaf struct {
	ID      string
	Amounts []int
	SumFn   zzz.SumFn[int]
}

func NewCalc(obj NafNaf) zzz.Calc[string, int] {
	return zzz.Calc[string, int]{
		ID:      obj.ID,
		Amounts: obj.Amounts,
		SumFn:   obj.SumFn,
	}
}

func (obj NafNaf) GetCalc() zzz.Calc[string, int] {
	return zzz.Calc[string, int]{
		ID:      obj.ID,
		Amounts: obj.Amounts,
		SumFn:   obj.SumFn,
	}
}
