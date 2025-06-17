package store

import (
	"sync"
	"time"
)

type Task struct {
	ID                string        `json:"id"`
	Status            string        `json:"status"`
	CreatedAt         time.Time     `json:"createdAt"`
	EstimatedDuration time.Duration `json:"estimatedtime"`
}
type StoreTask struct {
	tasks map[string]*Task
	mu    sync.RWMutex
}

func NewTaskStore() *StoreTask {
	return &StoreTask{
		tasks: make(map[string]*Task),
	}
}

func (s *StoreTask) AddTask(task *Task) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tasks[task.ID] = task
}

func (s *StoreTask) GetTask(id string) (*Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	task, ok := s.tasks[id]
	return task, ok
}

func (s *StoreTask) UpdateTaskStatus(id string, status string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if task, exists := s.tasks[id]; exists {
		task.Status = status
	}
}
