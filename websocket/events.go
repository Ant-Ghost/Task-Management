package websocket

import "encoding/json"

type EventName int

const (
	TaskCreated EventName = iota
	TaskUpdated
	TaskDeleted
	TaskStatusUpdated
	SubTaskUpdated
)

func (eventName EventName) String() string {
	switch eventName {
	case TaskCreated:
		return "TaskCreated"
	case TaskUpdated:
		return "TaskUpdated"
	case TaskDeleted:
		return "TaskDeleted"
	case TaskStatusUpdated:
		return "TaskStatusUpdated"
	case SubTaskUpdated:
		return "SubTaskAdded"
	default:
		return "UnknownEvent"
	}
}

type Event struct {
	Name     string `json:"name"`
	TaskId   uint   `json:"task_id"`
	ParentId uint   `json:"parent_id"`
}

func EmitEvent(eventName EventName, taskId uint, parentId *uint) {
	if parentId == nil {
		parentId = new(uint)
		*parentId = 0
	}
	event := Event{
		Name:     eventName.String(),
		TaskId:   taskId,
		ParentId: *parentId,
	}

	jsonBytes, err := json.Marshal(event)
	if err != nil {
		return
	}

	GlobalHub.broadcast <- jsonBytes

}
