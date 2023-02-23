package shared

type GenerationInput struct {
	Limit   *int64  `json:"limit,omitempty"`
	Model   *string `json:"model,omitempty"`
	Task    *string `json:"task,omitempty"`
	Version *string `json:"version,omitempty"`
}
