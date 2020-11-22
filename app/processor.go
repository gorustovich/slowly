package app

import (
	"time"
	"errors"
)

const limitTimeout = time.Second * 5

type SlowProcessor interface {
	Process(timeout time.Duration) error
}

func NewSlowProcessor() *slowProcessor {
	return &slowProcessor{}
}

type slowProcessor struct{}

func (sp *slowProcessor) Process(timeout time.Duration) error {
	if isOverlimit(timeout) {
		return errors.New("timeout too long")
	}
	if timeout < 0 {
		return errors.New("only positive timeout is allowed")
	}
	time.Sleep(timeout)
	return nil
}

func isOverlimit(timeout time.Duration) bool {
	return timeout > limitTimeout
}
