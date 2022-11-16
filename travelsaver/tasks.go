package travelsaver

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/procyon-projects/chrono"
)

type scheduler struct {
	taskScheduler chrono.TaskScheduler
}

func (s *scheduler) new() {
	s.taskScheduler = chrono.NewDefaultTaskScheduler()
}

func (s *scheduler) createTask(msg vitaliksMsg) (chrono.ScheduledTask, error) {

	now := time.Now()

	seconds := msg.TS - int(now.Unix())

	fmt.Printf("firing in: %v secs\n", seconds)

	startTime := now.Add(time.Second * time.Duration(seconds))

	task, err := s.taskScheduler.Schedule(func(ctx context.Context) {
		log.Print("One-Shot Task Fired:")
		prettyPrint(msg)

		err := writeCreateTravelPaymentPlan(msg)
		if err != nil {
			log.Printf("writeCreateTravelPaymentPlan err: %v\n", err)
		}

	}, chrono.WithTime(startTime))

	if err != nil {
		return task, fmt.Errorf("taskScheduler err: %v", err)
	}

	return task, nil

}
