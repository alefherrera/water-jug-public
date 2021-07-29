package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJug(t *testing.T) {
	jug := NewJug("x", 3)
	assert.True(t, jug.IsEmpty())
	assert.False(t, jug.IsFull())
	jug.Fill()
	assert.False(t, jug.IsEmpty())
	assert.True(t, jug.IsFull())
	jug.Empty()
	assert.True(t, jug.IsEmpty())
	assert.False(t, jug.IsFull())
}

func TestJug_Transfer(t *testing.T) {
	x := NewJug("x", 3)
	y := NewJug("y", 1)
	x.Fill()
	assert.True(t, x.IsFull())
	assert.True(t, y.IsEmpty())
	x.Transfer(y)
	assert.True(t, y.IsFull())
	assert.False(t, x.IsFull())
}
