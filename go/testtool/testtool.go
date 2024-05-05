package testtool

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

func CaptureOutput(function func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)

	output := make(chan string)
	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(1)
	go func() {
		var buffer bytes.Buffer
		waitGroup.Done()
		io.Copy(&buffer, reader)
		output <- buffer.String()
	}()
	waitGroup.Wait()
	function()
	writer.Close()
	return <-output
}

func IsExecutionTimeInRange(function func(), min time.Duration, max time.Duration) bool {
	if min > max {
		panic("min within the time range cannot be greater than max")
	}

	startTime := time.Now()
	function()
	executionTime := time.Since(startTime)
	return min <= executionTime && executionTime <= max
}
