package graceful

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/teng231/executor"
)

func TestGrace(t *testing.T) {
	exec := executor.RunSafeQueue(&executor.SafeQueueConfig{
		NumberWorkers: 4,
		Capacity:      100,
	})
	for i := 0; i < 100; i++ {
		exec.Send(&executor.Job{
			Params: []any{i},
			Exectutor: func(i ...interface{}) (interface{}, error) {
				log.Print("done ", i)
				time.Sleep(100 * time.Millisecond)
				return nil, nil
			},
		})
	}

	log.Print("preshutdown")
	PreShutdown(2*time.Second, map[string]func(context.Context) error{
		"exec": func(c context.Context) error {
			log.Print("shutdown wait")
			jobs := exec.TerminatingHandler()
			log.Print("xxxx: ", len(jobs))

			if len(jobs) > 0 {
				for _, job := range jobs {
					job.Exectutor(job.Params...)
				}
			}

			return nil
		},
		"http server": func(ctx context.Context) error {
			time.Sleep(500 * time.Millisecond)
			log.Print("shutdown http")
			return nil
		},
	})
}
