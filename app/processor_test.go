package app

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestIsOverlimit(t *testing.T) {
	ttt := []struct {
		requestTimeout time.Duration
		expected       bool
		name           string
	}{
		{requestTimeout: 0 * time.Millisecond, expected: false, name: "0"},
		{requestTimeout: 1 * time.Millisecond, expected: false, name: "1"},
		{requestTimeout: 1000 * time.Millisecond, expected: false, name: "1000"},
		{requestTimeout: 5000 * time.Millisecond, expected: false, name: "5000"},
		{requestTimeout: 5001 * time.Millisecond, expected: true, name: "5001"},
		{requestTimeout: 100000 * time.Millisecond, expected: true, name: "100000"},
	}
	for _, tt := range ttt {
		assert.Equal(t, tt.expected, isOverlimit(tt.requestTimeout), tt.name)
	}
}
