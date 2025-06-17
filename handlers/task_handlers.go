package handlers

import (
	"github/santori/tasks/store"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskHandler struct {
	Store *store.StoreTask
}

func NewTaskHandler(s *store.StoreTask) *TaskHandler {
	return &TaskHandler{Store: s}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	id := generateID()
	duration := 3 * time.Minute

	task := &store.Task{
		ID:                id,
		Status:            "In progress",
		CreatedAt:         time.Now(),
		EstimatedDuration: duration,
	}

	h.Store.AddTask(task)

	go func() {
		time.Sleep(duration)
		h.Store.UpdateTaskStatus(id, "done")
	}()

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *TaskHandler) GetTask(c *gin.Context) {
	id := c.Param("id")
	task, found := h.Store.GetTask(id)

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task not found"})
		return
	}
	if task.Status == "in_progress" {
		elapsed := time.Since(task.CreatedAt)
		timeLeft := task.EstimatedDuration - elapsed
		if timeLeft < 0 {
			timeLeft = 0
		}
		c.JSON(http.StatusOK, gin.H{
			"status":    task.Status,
			"time_left": timeLeft.Seconds(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": task.Status})

}
func generateID() string {
	id := uuid.New()
	return id.String()
}
