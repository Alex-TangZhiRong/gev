// +build linux

package eventloop

import (
	"log"
	"testing"
	"time"
)

func TestEventLoop_RunLoop(t *testing.T) {
	el, err := New()
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		go func() {
			el.QueueInLoop(func() {
				log.Println("runinloop")
			})
		}()
	}

	time.Sleep(time.Second)

	go func() {
		if err := el.Stop(); err != nil {
			panic(err)
		}
	}()

	el.RunLoop()
}
