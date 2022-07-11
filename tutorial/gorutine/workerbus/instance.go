package workerbus

var instance *WorkerBus

func Init(sizeBufferChanel, workerLimit int, debug bool) {
	instance = NewWorkerBus(sizeBufferChanel, workerLimit, debug)
}

func Gist() *WorkerBus {
	return instance
}
