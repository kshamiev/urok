package grpcclient // import "application/components/grpc-client"

import (
	"application/configuration"
	"application/workers"
	"application/workflow"
)

func init() { workflow.Register(New()) }

// New Create object and return interface
func New() Interface {
	var cpn = new(impl)
	return cpn
}

// After Возвращает массив зависимостей - имена пакетов компонентов, которые должны быть запущены до этого компонента
func (cpn *impl) After() []string {
	return []string{
		"application/components/bootstrap",
	}
}

// Init Функция инициализации компонента
func (cpn *impl) Init(appVersion string, appBuild string) (exitCode uint8, err error) {
	cpn.Cfg = configuration.Get()
	if cpn.Jbo, err = workers.Init(); err != nil {
		exitCode = workflow.ErrCantCreateWorkers
		return
	}

	return
}

// Start Выполнение компонента
func (cpn *impl) Start(cmd string) (done bool, exitCode uint8, err error) {
	const (
		keySingle = `ping simple`
		keyStream = `ping stream`
	)
	//var ()

	switch cmd {
	case keySingle, keyStream:
		done, exitCode = true, workflow.ErrNone
	default:
		return
	}
	// Выбор типа пинга
	switch cmd {
	case keySingle:
		err = cpn.PingSimple(
			cpn.Cfg.Configuration().PingServerAddress,
			cpn.Cfg.Configuration().PingInsecure,
			cpn.Cfg.Configuration().PingCount,
		)
	case keyStream:
		err = cpn.PingStream(
			cpn.Cfg.Configuration().PingServerAddress,
			cpn.Cfg.Configuration().PingInsecure,
			cpn.Cfg.Configuration().PingCount,
		)
	}

	return
}

// Stop Функция завершения работы компонента
func (cpn *impl) Stop() (exitCode uint8, err error) {
	return
}
