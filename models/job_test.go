package models

import (
	"fmt"
	"testing"
	"time"
)

func TestJobExecutionChecker(t *testing.T) {

	fmt.Println("Running test to check if its time to run the job")

	ct := CronTime{}
	ct.Second = 0
	ct.Minute = 15
	ct.Hour = 22
	ct.DayOfWeek = 6
	ct.DayOfMonth = 19
	ct.Month = 1

	jb := Job{ct, "somecallbackurl", "SomeTask"}

	someTime, err := time.Parse("2006-01-02 15:04", "2013-01-19 22:15")

	if err != nil {
		t.Errorf("Time cannot be parsed", err)
		return
	}

	if jb.JobExecutionChecker(someTime) {

		fmt.Println("Test passed")
		return

	}

	t.Fail()

}
