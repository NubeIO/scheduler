package scheduler

import (
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"testing"
	"time"
)

func TestNewScheduler(t *testing.T) {
	sched := NewScheduler()

	// Define a task function
	task := func(a, b int) {
		fmt.Println("Executing task", a+b)
	}

	// Add a job to the scheduler
	job, err := sched.Add("test",
		gocron.DurationJob(5*time.Second),
		gocron.NewTask(task, 1, 1),
	)
	if err != nil {
		fmt.Println("Error adding job:", err)
		return
	}

	sched.Start()

	fmt.Println("Job UUID:", job.UUID)

	select {
	case <-time.After(time.Second * 20):
	}

	// When you're done, shut it down
	err = sched.Shutdown()
	if err != nil {
		fmt.Println("Error shutting down scheduler:", err)
	}

}
