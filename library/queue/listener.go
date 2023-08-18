package queue

import (
	"gin-pro/library/queue/ijobs"
	"time"
)

func Listen() {
	go listenQueue()
}

func listenQueue() {
	for {
		select {
		case q := <-JobsChan:
			go func(j ijobs.IJobs, t time.Duration) {
				beforeHook()
				time.Sleep(time.Second * t)
				j.Handle()
				afterHook()
			}(q.Que, q.SleepTime)
		}
	}
}

func beforeHook() {}

func afterHook() {}
