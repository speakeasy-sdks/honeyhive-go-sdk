package operations

import (
	"github.com/speakeasy-sdks/honeyhive-go-sdk/pkg/models/shared"
)

type FindAllTasksRequest struct {
	Request *string `request:"mediaType=application/json"`
}

type FindAllTasksResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	StatusCode            int
	Tasks                 *shared.Tasks
}
