package test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
	"net/http"
)

func TestSlowApi(t *testing.T) {
	cl := NewClient()

	t.Run("timeout zero __ ok", func(t *testing.T) {
		t.Parallel()
		r, err := cl.SlowApiPost(0)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, r.Code)
		assert.Equal(t, "ok", r.Status)
	})

	t.Run("timeout 3000 __ ok", func(t *testing.T) {
		t.Parallel()
		r, err := cl.SlowApiPost(3000)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, r.Code)
		assert.Equal(t, "ok", r.Status)
	})

	t.Run("timeout 5000 __ ok", func(t *testing.T) {
		t.Parallel()
		r, err := cl.SlowApiPost(5000)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, r.Code)
		assert.Equal(t, "ok", r.Status)
	})

	t.Run("timeout 5001 __ too long", func(t *testing.T) {
		t.Parallel()
		r, err := cl.SlowApiPost(5001)
		require.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, r.Code)
		assert.Equal(t, "timeout too long", r.Error)
	})

	t.Run("timeout 700000 __ too long", func(t *testing.T) {
		t.Parallel()
		r, err := cl.SlowApiPost(700000)
		require.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, r.Code)
		assert.Equal(t, "timeout too long", r.Error)
	})

	t.Run("timeout -1000 __ negative not allowed", func(t *testing.T) {
		t.Parallel()
		r, err := cl.SlowApiPost(-1000)
		require.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, r.Code)
		assert.Equal(t, "only positive timeout is allowed", r.Error)
	})

}
