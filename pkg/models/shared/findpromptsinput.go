package shared

type FindPromptsInput struct {
	Name *string `json:"name,omitempty"`
	Task *string `json:"task,omitempty"`
}
