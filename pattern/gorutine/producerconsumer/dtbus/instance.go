package dtbus

var instance *DtBus

func Init(sizeBufferChanel int) {
	instance = NewDtBus(sizeBufferChanel)
}

func Gist() *DtBus {
	return instance
}
