package employee

import (
	"multithreading/meat"
	"multithreading/testtool"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/undefinedlabs/go-mpatch"
)

type meat1NoSeconds struct{}

func (meat1NoSeconds) GetName() string {
	return "肉1"
}
func (meat1NoSeconds) GetProcessingSeconds() int {
	return 0
}

type meat2NoSeconds struct{}

func (meat2NoSeconds) GetName() string {
	return "肉2"
}
func (meat2NoSeconds) GetProcessingSeconds() int {
	return 0
}

type meat1With1Seconds struct{}

func (meat1With1Seconds) GetName() string {
	return "肉1"
}
func (meat1With1Seconds) GetProcessingSeconds() int {
	return 1
}

type meat2With2Seconds struct{}

func (meat2With2Seconds) GetName() string {
	return "肉2"
}
func (meat2With2Seconds) GetProcessingSeconds() int {
	return 2
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

func TestEmployeeProcess_output(t *testing.T) {
	patch, _ := mpatch.PatchMethod(time.Now, func() time.Time {
		return time.Date(2001, 01, 01, 00, 00, 00, 0, time.UTC)
	})
	defer patch.Unpatch()

	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(2)
	channelMeat := make(chan meat.Meat, 2)
	channelMeat <- meat1NoSeconds{}
	channelMeat <- meat2NoSeconds{}
	output := testtool.CaptureOutput(func() {
		go Employee{Id: "ID"}.process(channelMeat, waitGroup)
		waitGroup.Wait()
	})
	assert.Equal(t, output, ("ID 在 2001-01-01 00:00:00 取得肉1\n" +
		"ID 在 2001-01-01 00:00:00 處理完肉1\n" +
		"ID 在 2001-01-01 00:00:00 取得肉2\n" +
		"ID 在 2001-01-01 00:00:00 處理完肉2\n"))
}

func TestEmployeeProcess_processingTime(t *testing.T) {
	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(2)
	channelMeat := make(chan meat.Meat, 2)
	channelMeat <- meat1With1Seconds{}
	channelMeat <- meat2With2Seconds{}
	assert.True(t, testtool.IsExecutionTimeInRange(func() {
		go Employee{Id: "ID"}.process(channelMeat, waitGroup)
		waitGroup.Wait()
	}, 3*time.Second, 3500*time.Millisecond))
}

func TestEmployeesProcess_output(t *testing.T) {
	meatList := []meat.Meat{meat1NoSeconds{}, meat2NoSeconds{}}
	employees := Employees{
		All: []Employee{{Id: "A"}, {Id: "B"}},
	}
	output := testtool.CaptureOutput(func() {
		employees.Process(meatList)
	})
	assert.Equal(t, strings.Count(output, "\n"), 2*2)
}

func TestEmployeesProcess_processingTime(t *testing.T) {
	meatList := []meat.Meat{meat1With1Seconds{}, meat2With2Seconds{}}
	employees := Employees{
		All: []Employee{{Id: "A"}, {Id: "B"}},
	}
	totalProcessingTime := 1*time.Second + 2*time.Second
	assert.True(t, testtool.IsExecutionTimeInRange(func() {
		employees.Process(meatList)
	}, totalProcessingTime/2, totalProcessingTime))
}
