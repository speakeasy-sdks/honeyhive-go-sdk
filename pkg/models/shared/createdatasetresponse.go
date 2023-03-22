// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreateDatasetResponse - Successful response
type CreateDatasetResponse struct {
	CreatedAt *string `json:"created_at,omitempty"`
	// description of the dataset
	Description *string `json:"description,omitempty"`
	// This is the file containing the dataset.
	FileName *string `json:"file_name,omitempty"`
	// id of the dataset
	ID *string `json:"id,omitempty"`
	// This is the name of the dataset.
	Name *string `json:"name,omitempty"`
	// This is the purpose of the dataset.
	Purpose *string `json:"purpose,omitempty"`
	// This is the task associated with the dataset.
	Task      *string `json:"task,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}
