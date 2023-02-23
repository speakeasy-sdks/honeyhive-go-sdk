package shared

type HyperParameters struct {
	FrequencyPenalty *int64   `json:"frequency_penalty,omitempty"`
	MaxTokens        *int64   `json:"max_tokens,omitempty"`
	PrescencePenalty *int64   `json:"prescence_penalty,omitempty"`
	Temperature      *float64 `json:"temperature,omitempty"`
	TopP             *int64   `json:"top_p,omitempty"`
}
