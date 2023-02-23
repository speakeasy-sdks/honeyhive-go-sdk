package shared

type CreateDatasetResponse struct {
	CreatedAt   *string `json:"created_at,omitempty"`
	Description *string `json:"description,omitempty"`
	FileName    *string `json:"file_name,omitempty"`
	ID          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Purpose     *string `json:"purpose,omitempty"`
	Task        *string `json:"task,omitempty"`
	UpdatedAt   *string `json:"updated_at,omitempty"`
}
