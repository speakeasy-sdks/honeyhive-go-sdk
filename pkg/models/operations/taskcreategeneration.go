package operations

import (
	"github.com/speakeasy-sdks/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type TaskCreateGenerationRequest struct {
	Request shared.TaskGenerationInput `request:"mediaType=application/json"`
}

type TaskCreateGenerationResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	StatusCode            int
	RawResponse           *http.Response
	TaskGenerationOutput  *shared.TaskGenerationOutput
}
