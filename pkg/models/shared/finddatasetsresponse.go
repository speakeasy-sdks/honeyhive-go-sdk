package shared

type FindDatasetsResponsePurposeEnum string

const (
	FindDatasetsResponsePurposeEnumTraining   FindDatasetsResponsePurposeEnum = "Training"
	FindDatasetsResponsePurposeEnumValidation FindDatasetsResponsePurposeEnum = "Validation"
	FindDatasetsResponsePurposeEnumTest       FindDatasetsResponsePurposeEnum = "Test"
)

type FindDatasetsResponse struct {
	Bytes       *string                          `json:"bytes,omitempty"`
	CreatedAt   *string                          `json:"created_at,omitempty"`
	Description *string                          `json:"description,omitempty"`
	FileName    *string                          `json:"file_name,omitempty"`
	ID          *string                          `json:"id,omitempty"`
	Object      *string                          `json:"object,omitempty"`
	Purpose     *FindDatasetsResponsePurposeEnum `json:"purpose,omitempty"`
}
