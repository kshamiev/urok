package nufnuf

import (
	"github.com/kshamiev/urok/tutorial/generics/zzz"
)

type NufNuf struct {
	ID      zzz.ID
	Amounts []uint64
	SumFn   zzz.SumFn[uint64]
}

func NewCalc[T ~string](obj NufNuf) zzz.Calc[T, uint64] {
	return zzz.Calc[T, uint64]{
		ID:      T(obj.ID),
		Amounts: obj.Amounts,
		SumFn:   obj.SumFn,
	}
}

func (obj NufNuf) GetCalc() zzz.Calc[string, uint64] {
	return zzz.Calc[string, uint64]{
		ID:      string(obj.ID),
		Amounts: obj.Amounts,
		SumFn:   obj.SumFn,
	}
}
