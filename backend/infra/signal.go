package infra

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	log "github.com/sirupsen/logrus"
)

var (
	ctx         context.Context
	wg          sync.WaitGroup
	cancel      context.CancelFunc
	mux         sync.Mutex
	initialised = false
)

func mustInitialise() {
	mux.Lock()
	defer mux.Unlock()
	if !initialised {
		initialiseContext()
		initialised = true
	}
}

// GetGracefulShutdownCtx returns a context that propagate signals which lead to graceful shutdown.
func GetGracefulShutdownCtx() context.Context {
	mustInitialise()
	return ctx
}

// GetGracefulShutdownWaitGroup returns a context that is waited on so the main() function knows when to return
func GetGracefulShutdownWaitGroup() *sync.WaitGroup {
	mustInitialise()
	mux.Lock()
	wg.Add(1)
	mux.Unlock()
	return &wg
}

// ManualShutdown godoc
func ManualShutdown() {
	cancel()
	log.Info("waiting for goroutines to finish")
	wg.Wait()
}

// WaitForShutdown godoc
func WaitForShutdown() {
	wg.Wait()
}

func initialiseContext() {
	// Create cancel context for shutting down server
	ctx, cancel = context.WithCancel(context.Background())

	// Create signal channel
	sigChan := make(chan os.Signal)
	// Fileter terminate signals
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start goroutine to catch shutdown signals
	go func() {
		defer cancel()

		sig := <-sigChan
		// Use default log to reduce package dependency
		log.WithFields(log.Fields{
			"signal": sig.String(),
		}).Info("initiating graceful shutdown")

		signal.Reset()
	}()
}
