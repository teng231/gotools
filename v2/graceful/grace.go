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

// PreShutdownSync sync to graceful shutdown
// task sync to stop task by task, priority is necessary
func PreShutdownSync(timeout time.Duration, syncHandler func(context.Context) error) {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, os.Interrupt)

	<-quit

	log.Println("shutting down")

	to := time.AfterFunc(timeout, func() {
		log.Print("time up")
		os.Exit(1)
	})
	defer to.Stop()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	syncHandler(ctx)

	log.Println("Shutdown Server")
}

// PreShutdownSync async to graceful shutdown
// task concurrent shutdown.
func PreShutdown(timeout time.Duration, syncHandlers map[string]func(context.Context) error) {
	if len(syncHandlers) == 0 {
		return
	}
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, os.Interrupt)

	<-quit

	log.Println("shutting down")

	to := time.AfterFunc(timeout, func() {
		log.Print("time up")
		os.Exit(1)
	})
	defer to.Stop()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	wg := &sync.WaitGroup{}
	wg.Add(len(syncHandlers))
	for _, handler := range syncHandlers {
		go func(handler func(context.Context) error) {
			defer wg.Done()
			handler(ctx)
			log.Print("done ")
		}(handler)
	}
	wg.Wait()

	log.Println("Shutdown Server")
}
