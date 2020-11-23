package app

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
)

func TestSlowProcessor_Process(t *testing.T) {
	ttt := []struct {
		requestTimeout time.Duration
		expected       error
		name           string
	}{
		{requestTimeout: 1 * time.Millisecond, expected: nil, name: "1"},
		{requestTimeout: -1000 * time.Millisecond, expected: OnlyPositiveTimeoutErr, name: "-1000"},
	}
	for _, tt := range ttt {
		t.Run(tt.name, func(t *testing.T) {
			p := NewSlowProcessor()
			assert.Equal(t, tt.expected, p.Process(tt.requestTimeout), tt.name)
		})
	}
}
