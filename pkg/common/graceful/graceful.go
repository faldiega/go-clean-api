package graceful

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Operation is a cleanup function on shutting down
type Operation func(ctx context.Context) error

// Shutdown waits for termination syscalls and does cleanup operations after received it.
// Optional sigChan can be passed to override signal receiving channel (useful for testing).
func Shutdown(ctx context.Context, timeout time.Duration, ops map[string]Operation, sigChan ...chan os.Signal) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		var s chan os.Signal
		if len(sigChan) > 0 {
			s = sigChan[0]
		} else {
			s = make(chan os.Signal, 10)
			signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		}

		sig := <-s
		log.Println("Shutting down due to signal:", sig)

		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Printf("Timeout %d ms elapsed, force exit", timeout.Milliseconds())
			os.Exit(0)
		})
		defer timeoutFunc.Stop()

		var wg sync.WaitGroup
		for key, op := range ops {
			wg.Add(1)
			go func(key string, op Operation) {
				defer wg.Done()
				log.Printf("Cleaning up: %s", key)
				if err := op(ctx); err != nil {
					log.Printf("%s: clean up failed: %s", key, err.Error())
					return
				}
				log.Printf("%s was shutdown gracefully", key)
			}(key, op)
		}

		wg.Wait()
		close(wait)
	}()
	return wait
}
