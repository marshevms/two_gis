package closer

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	mx      sync.Mutex
	closers []func()

	chSignal chan os.Signal
	wg       sync.WaitGroup
)

func init() {

	chSignal = make(chan os.Signal, 1)
	mx = sync.Mutex{}

	signal.Notify(chSignal, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-chSignal
		signal.Stop(chSignal)
		close(chSignal)

		mx.Lock()

		clsrs := closers
		wg.Add(len(clsrs))
		closers = nil

		mx.Unlock()

		for _, foo := range clsrs {
			foo()
			wg.Done()
		}

		wg.Wait()
	}()
}

func Add(foo func()) {
	mx.Lock()
	defer mx.Unlock()

	closers = append(closers, foo)
}

func Wait() {
	wg.Wait()
}
