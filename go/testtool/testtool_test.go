package testtool

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/undefinedlabs/go-mpatch"
)

func TestCaptureOutput(t *testing.T) {
	patch, _ := mpatch.PatchMethod(os.Pipe, func() (r *os.File, w *os.File, err error) {
		return nil, nil, errors.New("mock")
	})
	assert.Panics(t, func() { CaptureOutput(func() {}) })
	patch.Unpatch()

	assert.Equal(t, CaptureOutput(func() {}), "")
	assert.Equal(t, CaptureOutput(func() { fmt.Println("abc") }), "abc\n")
}

func TestIsExecutionTimeInRange(t *testing.T) {
	assert.Panics(t, func() { IsExecutionTimeInRange(func() {}, 0, -1) })

	assert.True(t, IsExecutionTimeInRange(func() {}, 0, 100*time.Millisecond))
	assert.False(t, IsExecutionTimeInRange(func() {}, -1, -1))
	assert.False(t, IsExecutionTimeInRange(func() {}, 100*time.Millisecond, 100*time.Millisecond))

	assert.True(t, IsExecutionTimeInRange(func() { time.Sleep(time.Second) }, 1, 1100*time.Millisecond))
	assert.False(t, IsExecutionTimeInRange(func() { time.Sleep(time.Second) }, 0, 0))
	assert.False(t, IsExecutionTimeInRange(func() { time.Sleep(time.Second) }, 1100*time.Millisecond, 1100*time.Millisecond))
}
