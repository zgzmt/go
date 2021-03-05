package tcpserver

import (
	"fmt"
	"golang.org/x/net/context"
	"lib/sync/atomic"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"zgz.redis/interface/tcp"
)

type Config struct {
	Address string
	MaxConnect uint32
	Timeout time.Duration
}
func ListenAndServer(cfg *Config,handle tcp.Handle){
	listener,err := net.Listen("tcp",cfg.Address)
	if err != nil {
		log.Fatal(fmt.Sprintf("listen err :v%",err))
	}
	//监听中断信号
	var closing atomic.Atomicbool
	sigCh := make(chan os.Signal,1)
	signal.Notify(sigCh,syscall.SIGHUP,syscall.SIGQUIT,syscall.SIGTERM, syscall.SIGINT)
	go func() {
		sig := <-sigCh
		switch sig{
		case syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Fatal("shuting dowm...")
			closing.SetAtomicUnit32(true)
			_ = listener.Close()
			_ = handle.Close()
		}
	}()
	log.Fatal(fmt.Sprintf("bin:%s,start listening...",cfg.Address))
	defer func(){
		_= listener.Close()
		_= handle.Close()
	}()
	ctx, _ := context.WithCancel(context.Background())
	var waitDone sync.WaitGroup
	for{
		conn, err := listener.Accept()
	}
}
