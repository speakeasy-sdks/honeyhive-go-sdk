package shared

type FindDatasetsInput struct {
	DatasetID *string `json:"dataset_id,omitempty"`
	Prompt    *string `json:"prompt,omitempty"`
	Purpose   *string `json:"purpose,omitempty"`
	Task      *string `json:"task,omitempty"`
}
