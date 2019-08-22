package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg := &sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		go start(ctx, wg, 5*time.Second, i)
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)
	fmt.Println("waiting for signal")
	<-sig
	cancel()
	fmt.Println("got quit signal")
	wg.Wait()
}

func start(ctx context.Context, wg *sync.WaitGroup, intervalDuration time.Duration, id int) {
	defer wg.Done()
	wg.Add(1)
	ticker := time.NewTicker(intervalDuration)
	defer ticker.Stop()
	for {
		select {
		case x := <-ticker.C:
			time.Sleep(time.Second)
			fmt.Printf("id: %v, x: %v\n", id, x)
		case <-ctx.Done():
			fmt.Printf("done for id: %v \n", id)
			return
		}
	}
}
