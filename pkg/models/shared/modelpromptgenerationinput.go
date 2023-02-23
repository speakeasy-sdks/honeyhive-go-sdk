package shared

type ModelPromptGenerationInput struct {
	Hyperparameters *HyperParameters `json:"hyperparameters,omitempty"`
	Model           *string          `json:"model,omitempty"`
	Prompt          *string          `json:"prompt,omitempty"`
	Source          *string          `json:"source,omitempty"`
	Task            *string          `json:"task,omitempty"`
	Version         *string          `json:"version,omitempty"`
}
