package shared

type FineTunedModelsInput struct {
	ModelID *string `json:"model_id,omitempty"`
	Task    *string `json:"task,omitempty"`
}
