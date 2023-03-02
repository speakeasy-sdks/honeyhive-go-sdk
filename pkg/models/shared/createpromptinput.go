package shared

type CreatePromptInput struct {
	FewShotExamples []string         `json:"few_shot_examples,omitempty"`
	Hyperparameters *HyperParameters `json:"hyperparameters,omitempty"`
	InputVariables  []string         `json:"input_variables,omitempty"`
	Model           *int64           `json:"model,omitempty"`
	Provider        *string          `json:"provider,omitempty"`
	Required        interface{}      `json:"required,omitempty"`
	Task            *string          `json:"task,omitempty"`
	Text            *string          `json:"text,omitempty"`
	Version         *string          `json:"version,omitempty"`
}
