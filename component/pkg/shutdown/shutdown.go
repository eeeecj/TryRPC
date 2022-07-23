package shutdown

import "sync"

//ShutdownCallback 是回调必须实现的接口。
//OnShutdown 会在需要的时候调用。参数是ShutdownManager需要的。
type ShutdownCallback interface {
	OnShutdown(string2 string) error
}

// ShutdownFunc is a helper type, so you can easily provide anonymous functions
// as ShutdownCallbacks.
type ShutdownFunc func(string2 string) error

// OnShutdown 定义触发关闭时需要运行的操作。
func (f ShutdownFunc) OnShutdown(shutdownManager string) error {
	return f(shutdownManager)
}

//ShutdownManager 是由ShutdownManagers实现的接口。
//GetName 返回ShutdownManager的名字
//Start ShutdownManager开始监听shutdown请求，当GSInterface调用时。
// 首先调用 ShutdownStart()，然后执行所有 ShutdownCallbacks
// 一旦所有 ShutdownCallbacks 返回，就调用 ShutdownFinish。
type ShutdownManager interface {
	GetName() string
	Start(gs GSInterface) error
	ShutdownStart() error
	ShutdownFinish() error
}

// ErrorHandler is an interface you can pass to SetErrorHandler to
// handle asynchronous errors.
type ErrorHandler interface {
	OnError(err error)
}

// ErrorFunc is a helper type, so you can easily provide anonymous functions
// as ErrorHandlers.
type ErrorFunc func(err error)

// OnError defines the action needed to run when errors occurred.
func (f ErrorFunc) OnError(err error) {
	f(err)
}

// GSInterface is an interface implemented by GracefulShutdown,
// that gets passed to ShutdownManager to call StartShutdown when shutdown
// is requested.
type GSInterface interface {
	StartShutdown(sm ShutdownManager)
	ReportError(err error)
	AddShutdownCallback(callback ShutdownCallback)
}

// GracefulShutdown is main struct that handles ShutdownCallbacks and
// ShutdownManagers. Initialize it with New.
type GracefulShutdown struct {
	callbacks    []ShutdownCallback
	managers     []ShutdownManager
	errorHandler ErrorHandler
}

// New initializes GracefulShutdown.
func New() *GracefulShutdown {
	return &GracefulShutdown{
		callbacks:    make([]ShutdownCallback, 0, 10),
		managers:     make([]ShutdownManager, 0, 3),
		errorHandler: nil,
	}
}

// Start calls Start on all added ShutdownManagers. The ShutdownManagers
// start to listen to shutdown requests. Returns an errors if any ShutdownManagers
// return an errors.
func (gs *GracefulShutdown) Start() error {
	for _, manager := range gs.managers {
		//启动manager.start()实现对信号的检测
		if err := manager.Start(gs); err != nil {
			return err
		}
	}
	return nil
}

// AddShutdownManager adds a ShutdownManager that will listen to shutdown requests.
func (gs *GracefulShutdown) AddShutdownManager(manager ShutdownManager) {
	gs.managers = append(gs.managers, manager)
}

// AddShutdownCallback adds a ShutdownCallback that will be called when
// shutdown is requested.
//
// You can provide anything that implements ShutdownCallback interface,
// or you can supply a function like this:
//	AddShutdownCallback(shutdown.ShutdownFunc(func() errors {
//		// callback code
//		return nil
//	}))
func (gs *GracefulShutdown) AddShutdownCallback(callback ShutdownCallback) {
	gs.callbacks = append(gs.callbacks, callback)
}

// SetErrorHandler sets an ErrorHandler that will be called when an errors
// is encountered in ShutdownCallback or in ShutdownManager.
//
// You can provide anything that implements ErrorHandler interface,
// or you can supply a function like this:
//	SetErrorHandler(shutdown.ErrorFunc(func (err errors) {
//		// handle errors
//	}))
func (gs *GracefulShutdown) SetErrorHandler(handler ErrorHandler) {
	gs.errorHandler = handler
}

// StartShutdown is called from a ShutdownManager and will initiate shutdown.
// first call ShutdownStart on Shutdownmanager,
// call all ShutdownCallbacks, wait for callbacks to finish and
// call ShutdownFinish on ShutdownManager.
func (gs *GracefulShutdown) StartShutdown(sm ShutdownManager) {
	gs.ReportError(sm.ShutdownStart())
	var wg sync.WaitGroup
	for _, shutdownCallback := range gs.callbacks {
		wg.Add(1)
		go func(callback ShutdownCallback) {
			defer wg.Done()
			gs.ReportError(callback.OnShutdown(sm.GetName()))
		}(shutdownCallback)
	}
	wg.Wait()
	gs.ReportError(sm.ShutdownFinish())
}

// ReportError is a function that can be used to report errors to
// ErrorHandler. It is used in ShutdownManagers.
func (gs *GracefulShutdown) ReportError(err error) {
	if err != nil && gs.errorHandler != nil {
		gs.errorHandler.OnError(err)
	}
}
