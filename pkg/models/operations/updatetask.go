package operations

import (
	"github.com/speakeasy-sdks/honeyhive-go-sdk/pkg/models/shared"
)

type UpdateTaskRequest struct {
	Request shared.Task `request:"mediaType=application/json"`
}

type UpdateTaskResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	StatusCode            int
	Task                  *shared.Task
}
