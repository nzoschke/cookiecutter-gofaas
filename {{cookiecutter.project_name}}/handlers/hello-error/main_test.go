package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	t.Run("handler", func(t *testing.T) {
		_, err := handler(context.Background())
		assert.Error(t, err)
	})
}
