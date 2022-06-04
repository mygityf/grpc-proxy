package stop

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 停止信号处理
func SignalHandler() {
	shutdownHook := make(chan os.Signal, 1)
	signal.Notify(shutdownHook,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		os.Interrupt)
	sig := <-shutdownHook

	log.Printf("caught sig exit sig:%v", sig)
	time.Sleep(3 * time.Second)
}
