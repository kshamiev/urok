package nifnif

import (
	"github.com/kshamiev/urok/tutorial/generics/zzz"
)

type NifNif struct {
	ID      zzz.ID
	Amounts []float64
	SumFn   zzz.SumFn[float64]
}

func NewCalc[T ~string](obj NifNif) zzz.Calc[T, float64] {
	return zzz.Calc[T, float64]{
		ID:      T(obj.ID),
		Amounts: obj.Amounts,
		SumFn:   obj.SumFn,
	}
}

func (obj NifNif) GetCalc() zzz.Calc[string, float64] {
	return zzz.Calc[string, float64]{
		ID:      string(obj.ID),
		Amounts: obj.Amounts,
		SumFn:   obj.SumFn,
	}
}
