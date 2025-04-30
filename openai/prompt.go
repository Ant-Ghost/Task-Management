package openai

import (
	"encoding/json"

	"github.com/sashabaranov/go-openai"
)

type Subtask struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type AiResponse struct {
	Subtasks []Subtask `json:"subtasks"`
}

func (r AiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(r)
}

var JsonResponseFomat *openai.ChatCompletionResponseFormat
var SystemMessage openai.ChatCompletionMessage

func generateSystemMessage() {

	requirement := "You are a task manager AI. Please break down the task into smaller subtasks"
	note := "NOTE: Reply in JSON format with the following structure:"

	sampleJson := `{
		"subtasks": [
			{
				"title": "Subtask 1",
				"description": "Description of subtask 1"
			},
			{
				"title": "Subtask 2",
				"description": "Description of subtask 2"
			}
		]
	}`

	prompt := requirement + "\n" + note + "\n" + sampleJson

	SystemMessage = openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: prompt,
	}
}

func generateResponseFormat() {

	JsonResponseFomat = &openai.ChatCompletionResponseFormat{
		Type: openai.ChatCompletionResponseFormatTypeJSONSchema,
		JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
			Name:        "subtask_list",
			Description: "A list of subtasks",
			Schema:      AiResponse{},
			Strict:      true,
		},
	}
}
