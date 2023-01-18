package maker

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/procyon-projects/chrono"
)

type vitaliksMsg struct {
	TS    int
	ID    int
	ADDR  string
	HTTPS string
}

type scheduler struct {
	taskScheduler chrono.TaskScheduler
}

func (s *scheduler) new() {
	s.taskScheduler = chrono.NewDefaultTaskScheduler()
}

func (s *scheduler) createTask(msg vitaliksMsg, logger *log.Logger) (chrono.ScheduledTask, error) {

	now := time.Now()

	seconds := msg.TS - int(now.Unix())

	logger.Printf("firing in: %v secs\n", seconds)

	startTime := now.Add(time.Second * time.Duration(seconds))

	task, err := s.taskScheduler.Schedule(func(ctx context.Context) {
		logger.Print("One-Shot Task Fired:")
		logger.Println(msg)

		err := writeCreateTravelPaymentPlan(msg, logger)
		if err != nil {
			logger.Printf("writeCreateTravelPaymentPlan err: %v\n", err)
		}

	}, chrono.WithTime(startTime))

	if err != nil {
		return task, fmt.Errorf("taskScheduler err: %v", err)
	}

	return task, nil

}
