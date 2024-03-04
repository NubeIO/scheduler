package scheduler

import (
	"fmt"
	"github.com/go-co-op/gocron/v2"
)

// Scheduler interface defines the methods for scheduling tasks.
type Scheduler interface {
	GetScheduler() gocron.Scheduler
	All() []*Job
	Add(name string, job gocron.JobDefinition, task gocron.Task) (*Job, error)
	Delete(uuid string) error
	Start()
	Shutdown() error
}

// scheduler is an implementation of the Scheduler interface.
type scheduler struct {
	all           []*Job
	cronScheduler gocron.Scheduler
}

func (sch *scheduler) GetScheduler() gocron.Scheduler {
	return sch.cronScheduler
}

// Job represents a scheduled job.
type Job struct {
	Name    string
	UUID    string
	CronJob gocron.Job
}

// NewScheduler creates a new scheduler instance.
func NewScheduler() Scheduler {
	s, _ := gocron.NewScheduler()
	return &scheduler{
		all:           []*Job{},
		cronScheduler: s,
	}
}

// All returns all scheduled jobs.
func (sch *scheduler) All() []*Job {
	return sch.all
}

// Add adds a new job to the scheduler.
func (sch *scheduler) Add(name string, job gocron.JobDefinition, task gocron.Task) (*Job, error) {
	j, err := sch.cronScheduler.NewJob(job, task)
	if err != nil {
		return nil, err
	}
	newJob := &Job{
		Name:    name,
		UUID:    j.ID().String(),
		CronJob: j,
	}
	sch.all = append(sch.all, newJob)
	return newJob, nil
}

// Delete removes a job from the scheduler by its UUID.
func (sch *scheduler) Delete(uuid string) error {
	for i, job := range sch.all {
		if job.UUID == uuid {
			err := sch.cronScheduler.RemoveJob(job.CronJob.ID())
			if err != nil {
				return err
			}
			sch.all = append(sch.all[:i], sch.all[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("job with UUID %s cronScheduler not found", uuid)
}

func (sch *scheduler) Start() {
	sch.cronScheduler.Start()
}

func (sch *scheduler) Shutdown() error {
	return sch.cronScheduler.Shutdown()
}
