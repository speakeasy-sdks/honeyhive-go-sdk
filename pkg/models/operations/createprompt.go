package operations

import (
	"github.com/speakeasy-sdks/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type CreatePromptRequest struct {
	Request shared.CreatePromptInput `request:"mediaType=application/json"`
}

type CreatePromptResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	CreatePromptResponse  *shared.CreatePromptResponse
	StatusCode            int
	RawResponse           *http.Response
}
