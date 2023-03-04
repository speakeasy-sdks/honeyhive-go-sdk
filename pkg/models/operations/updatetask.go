package operations

import (
	"github.com/speakeasy-sdks/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateTaskRequest struct {
	Request shared.Task `request:"mediaType=application/json"`
}

type UpdateTaskResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	StatusCode            int
	RawResponse           *http.Response
	Task                  *shared.Task
}
