package reducer

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	DEFAULT_CAPACITY = 500
)

type Task struct {
	Message  interface{}
	Headers  map[string]string
	Topic    string
	Callback func(error)
	Handler  func(data interface{}) error
	Publish  func(string, ...interface{}) error
}

type Reducer struct {
	numberConcurrent int
	topicName        string
	Tasks            chan *Task
	ctx              context.Context
	cancelFunc       context.CancelFunc
}

// Reducer depend on kafka. Slow task for handler
type IReducer interface {
	// Start is fn start reducer flow
	Start()
	// Push is push task to process
	Push()
	// Consume is listen event. It's optional
	Consume()
	// Close reducer instance.
	Close()
}

func (r *Reducer) Start(numberConcurrent, capacity int, topicName string) {
	if capacity == 0 {
		capacity = DEFAULT_CAPACITY
	}
	r.numberConcurrent = numberConcurrent
	r.topicName = topicName
	r.Tasks = make(chan *Task, capacity)
	ctx, cancel := context.WithCancel(context.Background())
	r.ctx = ctx
	r.cancelFunc = cancel
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	// number concurrent  goroutines count
	for i := 0; i < numberConcurrent; i++ {
		go func() {
			for {
				select {
				case task := <-r.Tasks:
					// do somethings
					log.Print(task)
					if task.Publish == nil {
						continue
					}
					err := task.Publish(task.Topic, task.Message)
					if task.Callback != nil && err != nil {
						task.Callback(err)
					}
				case <-sig:
					for task := range r.Tasks {
						if task.Publish == nil {
							continue
						}
						err := task.Publish(task.Topic, task.Message)
						if task.Callback != nil && err != nil {
							task.Callback(err)
						}
					}
				case <-r.ctx.Done():
					close(sig)
					return
				}
			}
		}()
	}
}

func (r *Reducer) Push(task *Task) {
	r.Tasks <- task
}

// Consume is not requred, you can use for pub and sub whe
func (r *Reducer) Consume(topic string) {

}

func (r *Reducer) Close() {
	r.cancelFunc()
	close(r.Tasks)
}
