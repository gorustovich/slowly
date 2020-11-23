package app

import (
	"time"
	"errors"
)

var OnlyPositiveTimeoutErr = errors.New("only positive timeout is allowed")

type SlowProcessor interface {
	Process(timeout time.Duration) error
}

func NewSlowProcessor() *slowProcessor {
	return &slowProcessor{}
}

type slowProcessor struct{}

func (sp *slowProcessor) Process(timeout time.Duration) error {
	if timeout < 0 {
		return OnlyPositiveTimeoutErr
	}
	time.Sleep(timeout)
	return nil
}
