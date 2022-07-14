package workerbus

var instance *WorkerBus

func Init(sizeBufferChanel, workerLimit int) {
	instance = NewWorkerBus(sizeBufferChanel, workerLimit)
}

func Gist() *WorkerBus {
	return instance
}
