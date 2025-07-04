package scrapper

import (
	"github.com/shivamhw/content-pirate/commons"
	"github.com/shivamhw/content-pirate/store"
)

type TaskStatusEnum string

const (
	TaskCreated TaskStatusEnum = "TASK_CREATED"
	TaskStarted TaskStatusEnum = "TASK_STARTED"
	TaskDone    TaskStatusEnum = "TASK_DONE"
)

type TaskStatus struct {
	ItemDone  int64
	TotalItem int64
	Status    TaskStatusEnum
}

type Task struct {
	Id     string
	J      Job
	I      []commons.Item
	Status TaskStatus
	S      []store.Store `json:"-"`
}

type TaskUpdateOpts struct {
	*TaskStatus
	Items []commons.Item
}
