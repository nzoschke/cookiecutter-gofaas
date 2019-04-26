package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	t.Run("handler", func(t *testing.T) {
		out, err := handler(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, "Hello world!", out)
	})
}
