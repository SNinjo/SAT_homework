package employee

import (
	"bytes"
	"io"
	"log"
	"multithreading/meat"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/undefinedlabs/go-mpatch"
)

type meat1NoSeconds struct{}

func (_ meat1NoSeconds) GetName() string {
	return "肉1"
}
func (_ meat1NoSeconds) GetProcessingSeconds() int {
	return 0
}

type meat2NoSeconds struct{}

func (_ meat2NoSeconds) GetName() string {
	return "肉2"
}
func (_ meat2NoSeconds) GetProcessingSeconds() int {
	return 0
}

func captureOutput(function func()) string {
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

	out := make(chan string)
	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(1)
	go func() {
		var buf bytes.Buffer
		waitGroup.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	waitGroup.Wait()
	function()
	writer.Close()
	return <-out
}

func TestGetCurrentTime(t *testing.T) {
	patch, _ := mpatch.PatchMethod(time.Now, func() time.Time {
		return time.Date(1970, 01, 01, 00, 00, 00, 0, time.UTC)
	})
	assert.Equal(t, getCurrentTime(), "1970-01-01 00:00:00")
	patch.Unpatch()

	patch, _ = mpatch.PatchMethod(time.Now, func() time.Time {
		return time.Date(2001, 12, 31, 12, 59, 59, 999, time.UTC)
	})
	assert.Equal(t, getCurrentTime(), "2001-12-31 12:59:59")
	patch.Unpatch()
}

func TestEmployeeProcess(t *testing.T) {
	mpatch.PatchMethod(time.Now, func() time.Time {
		return time.Date(2001, 01, 01, 00, 00, 00, 0, time.UTC)
	})
	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(2)
	channelMeat := make(chan meat.Meat, 2)
	channelMeat <- meat1NoSeconds{}
	channelMeat <- meat2NoSeconds{}
	output := captureOutput(func() {
		go Employee{Id: "ID"}.process(channelMeat, waitGroup)
		waitGroup.Wait()
	})
	assert.Equal(t, output, ("ID 在 2001-01-01 00:00:00 取得肉1\n" +
		"ID 在 2001-01-01 00:00:00 處理完肉1\n" +
		"ID 在 2001-01-01 00:00:00 取得肉2\n" +
		"ID 在 2001-01-01 00:00:00 處理完肉2\n"))
}
