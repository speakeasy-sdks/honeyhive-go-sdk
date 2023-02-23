package shared

type Dataset struct {
	Description *string `json:"description,omitempty"`
	File        *string `json:"file,omitempty"`
	Name        *string `json:"name,omitempty"`
	Purpose     *string `json:"purpose,omitempty"`
	Task        *string `json:"task,omitempty"`
}
