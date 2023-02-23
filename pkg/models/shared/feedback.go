package shared

type FeedbackFeedbackJSON struct {
	Comment *string `json:"comment,omitempty"`
	Score   *int64  `json:"score,omitempty"`
}

type Feedback struct {
	FeedbackJSON *FeedbackFeedbackJSON `json:"feedback_json,omitempty"`
	GenerationID *string               `json:"generation_id,omitempty"`
	Task         *string               `json:"task,omitempty"`
	Timestamp    *string               `json:"timestamp,omitempty"`
}
