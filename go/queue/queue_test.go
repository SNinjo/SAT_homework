package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert.Equal(t, New().slice, []interface{}{})
}

func TestIsEmpty(t *testing.T) {
	assert.True(t, New().IsEmpty())
	assert.False(t, Queue{slice: []interface{}{1}}.IsEmpty())
}

func TestPut(t *testing.T) {
	q := New()
	q.Put(0)
	assert.Equal(t, q.slice, []interface{}{0})
	q.Put(1)
	assert.Equal(t, q.slice, []interface{}{0, 1})
}

func TestGet(t *testing.T) {
	q := Queue{slice: []interface{}{0, 1}}
	assert.Equal(t, q.Get(), 0)
	assert.Equal(t, q.slice, []interface{}{1})
	assert.Equal(t, q.Get(), 1)
	assert.Equal(t, q.slice, []interface{}{})
}
