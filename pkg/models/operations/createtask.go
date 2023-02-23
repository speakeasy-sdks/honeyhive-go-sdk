package operations

import (
	"github.com/speakeasy-sdks/honeyhive-go-sdk/pkg/models/shared"
)

type CreateTaskRequest struct {
	Request shared.Task `request:"mediaType=application/json"`
}

type CreateTaskResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	StatusCode            int
	Task                  *shared.Task
}
