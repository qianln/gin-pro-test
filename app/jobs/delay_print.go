package jobs

import "fmt"

func NewDelayPrint() *DelayPrint {
	return &DelayPrint{}
}

type DelayPrint struct{}

func (d DelayPrint) Handle() {
	fmt.Println("DelayPrint")
}
