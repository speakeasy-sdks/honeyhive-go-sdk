package operations

import (
	"github.com/speakeasy-sdks/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type FindPromptsRequest struct {
	Request *shared.FindPromptsInput `request:"mediaType=application/json"`
}

type FindPromptsResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	Prompts               *shared.Prompts
	StatusCode            int
	RawResponse           *http.Response
}
