package workerbus

var instance *WorkerBus

func Init(sizeBufferChanel, workerLimit int) {
	if instance == nil {
		instance = NewWorkerBus(sizeBufferChanel, workerLimit)
	}
}

func Gist() *WorkerBus {
	return instance
}
