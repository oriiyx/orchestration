package main

import (
	"fmt"
	"time"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/oriiyx/orchestration/manager"
	"github.com/oriiyx/orchestration/node"
	"github.com/oriiyx/orchestration/task"
	"github.com/oriiyx/orchestration/worker"
)

func main() {
	t := task.Task{
		ID:     uuid.New(),
		Name:   "Task",
		State:  task.Pending,
		Image:  "",
		Memory: 1024,
		Disk:   1,
	}

	te := task.TaskEvent{
		ID:        uuid.New(),
		State:     task.Pending,
		Timestamp: time.Now(),
		Task:      t,
	}

	fmt.Printf("task: %v", t)
	fmt.Printf("task event: %v", te)

	w := worker.Worker{Name: "worker-1", Queue: *queue.New(), Db: make(map[uuid.UUID]*task.Task)}
	fmt.Printf("worker: %v", w)
	w.CollectStats()
	w.RunTask()
	w.StartTask()
	w.StopTask()

	m := manager.Manager{Pending: *queue.New(), TaskDb: make(map[string][]*task.Task), EventDb: make(map[string][]*task.TaskEvent), Workers: []string{w.Name}}
	fmt.Printf("manager: %v", m)
	m.SelectWorker()
	m.UpdateTask()
	m.SendWork()

	n := node.Node{
		Name:   "Node-1",
		Ip:     "192.168.1.1",
		Cores:  4,
		Memory: 1024,
		Disk:   25,
		Role:   "worker",
	}
	fmt.Printf("node: %v", n)
}
