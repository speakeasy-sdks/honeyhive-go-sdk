package shared

type TaskGenerationInputProviderAPIKeys struct {
	APIKeys []string `json:"api_keys,omitempty"`
}

type TaskGenerationInputUserProperties struct {
	UserID *string `json:"user_id,omitempty"`
}

type TaskGenerationInput struct {
	BestOf          *int64                              `json:"best_of,omitempty"`
	Metric          *string                             `json:"metric,omitempty"`
	ModelID         *string                             `json:"model_id,omitempty"`
	Prompts         *Prompts                            `json:"prompts,omitempty"`
	ProviderAPIKeys *TaskGenerationInputProviderAPIKeys `json:"provider_api_keys,omitempty"`
	Sampling        *string                             `json:"sampling,omitempty"`
	Source          *string                             `json:"source,omitempty"`
	Task            *string                             `json:"task,omitempty"`
	UserProperties  *TaskGenerationInputUserProperties  `json:"user_properties,omitempty"`
}
