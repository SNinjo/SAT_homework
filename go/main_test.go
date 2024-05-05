package main

import (
	"multithreading/testtool"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	output := testtool.CaptureOutput(func() { main() })
	assert.Equal(t, strings.Count(output, "\n"), (10+7+5)*2)

	totalProcessingTime := 10*1*time.Second + 7*2*time.Second + 5*3*time.Second
	assert.True(t, testtool.IsExecutionTimeInRange(func() {
		main()
	}, totalProcessingTime/5, totalProcessingTime))
}
