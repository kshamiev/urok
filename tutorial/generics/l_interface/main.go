package main

import (
	"github.com/kshamiev/urok/tutorial/generics/zzz"
	"github.com/kshamiev/urok/tutorial/generics/zzz/nafnaf"
	"github.com/kshamiev/urok/tutorial/generics/zzz/nifnif"
	"github.com/kshamiev/urok/tutorial/generics/zzz/nufnuf"
)

func main() {
	//
	obj1 := nufnuf.NufNuf{
		ID:      "NUFNUF",
		Amounts: []uint64{1, 3, 5, 7, 9, 20, 60},
		SumFn:   zzz.Sum[uint64],
	}
	calc1 := nufnuf.NewCalc[string](obj1)
	calc1.PrintIDAndSum()

	calc1 = obj1.GetCalc()
	calc1.PrintIDAndSum()

	zzz.PrintLedger(calc1)

	//
	obj2 := nafnaf.NafNaf{
		ID:      "NAFNAF",
		Amounts: []int{2, 4, 6, 8, 10, 30, 70},
		SumFn:   zzz.Sum[int],
	}
	calc2 := nafnaf.NewCalc(obj2)
	calc2.PrintIDAndSum()

	calc2 = obj2.GetCalc()
	calc2.PrintIDAndSum()

	zzz.PrintLedger(calc2)

	//
	obj3 := nifnif.NifNif{
		ID:      "NIFNIF",
		Amounts: []float64{1.34, 3.65, 5.39, 7.95, 9.82, 20.46, 60.84},
		SumFn:   zzz.Sum[float64],
	}
	calc3 := nifnif.NewCalc[string](obj3)
	calc3.PrintIDAndSum()

	calc3 = obj3.GetCalc()
	calc3.PrintIDAndSum()

	zzz.PrintLedger(calc3)
}
