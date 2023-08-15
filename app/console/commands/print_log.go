package commands

import (
	"gin-pro/app/core/system"
)

func NewPrintLog() *PrintLog {
	return &PrintLog{}
}

type PrintLog struct{}

func (t *PrintLog) Handle() func() {
	
	f := func() {
		system.ZapLog.Info("NewPrintLog sleep print")
	}

	return f
}
