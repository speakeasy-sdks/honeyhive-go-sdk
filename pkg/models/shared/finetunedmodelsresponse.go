package shared

type FineTunedModelsResponse struct {
	CreatedAt       *string  `json:"created_at,omitempty"`
	Hyperparams     *string  `json:"hyperparams,omitempty"`
	Model           *string  `json:"model,omitempty"`
	ModelID         *string  `json:"model_id,omitempty"`
	Object          *string  `json:"object,omitempty"`
	OrgID           *string  `json:"org_id,omitempty"`
	ResultFiles     []string `json:"result_files,omitempty"`
	Status          *string  `json:"status,omitempty"`
	TrainingFiles   []string `json:"training_files,omitempty"`
	UpdatedAt       *string  `json:"updated_at,omitempty"`
	ValidationFiles []string `json:"validation_files,omitempty"`
}
