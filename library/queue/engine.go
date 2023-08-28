package queue

import (
	"gin-pro/library/queue/ijobs"
	"time"
)

var JobsChan = make(chan Engine, 100)

func NewEngine() *Engine {
	return &Engine{}
}

type Engine struct {
	SleepTime time.Duration
	Que       ijobs.IJobs
}

func (s *Engine) Dispatch(q ijobs.IJobs) {
	qqs := Engine{
		SleepTime: s.SleepTime,
		Que:       q,
	}
	s.afterHook()
	JobsChan <- qqs
	s.beforeHook()
}

func (s *Engine) Delay(t time.Duration) *Engine {
	return &Engine{
		SleepTime: t,
	}
}

func (s *Engine) beforeHook() {}

func (s *Engine) afterHook() {}
