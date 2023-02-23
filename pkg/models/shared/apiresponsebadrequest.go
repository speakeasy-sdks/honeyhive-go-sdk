package shared

type APIResponseBadRequest struct {
	Error  *string `json:"error,omitempty"`
	Status *int    `json:"status,omitempty"`
}
