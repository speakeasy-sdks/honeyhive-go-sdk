package shared

type Metric struct {
	CodeSnippet *string `json:"code_snippet,omitempty"`
	Name        *string `json:"name,omitempty"`
	TaskName    *string `json:"task_name,omitempty"`
}
