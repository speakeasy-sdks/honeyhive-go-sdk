package shared

type MetricsResponse struct {
	CreatedAt   *string `json:"created_at,omitempty"`
	ID          *string `json:"id,omitempty"`
	MetricName  *string `json:"metric_name,omitempty"`
	MetricValue *string `json:"metric_value,omitempty"`
	TaskName    *string `json:"task_name,omitempty"`
	UpdatedAt   *string `json:"updated_at,omitempty"`
}
