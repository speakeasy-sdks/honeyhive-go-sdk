package shared

type TaskGenerationOutput struct {
	Cost         *float64 `json:"cost,omitempty"`
	Generation   *string  `json:"generation,omitempty"`
	GenerationID *string  `json:"generation_id,omitempty"`
	Latency      *float64 `json:"latency,omitempty"`
	Prompt       *string  `json:"prompt,omitempty"`
	TotalTokens  *int64   `json:"total_tokens,omitempty"`
}
