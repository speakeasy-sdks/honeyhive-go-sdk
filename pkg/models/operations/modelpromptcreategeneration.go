package operations

import (
	"github.com/speakeasy-sdks/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type ModelPromptCreateGenerationRequest struct {
	Request shared.ModelPromptGenerationInput `request:"mediaType=application/json"`
}

type ModelPromptCreateGenerationResponse struct {
	APIResponseBadRequest       *shared.APIResponseBadRequest
	ContentType                 string
	ModelPromptGenerationOutput *shared.ModelPromptGenerationOutput
	StatusCode                  int
	RawResponse                 *http.Response
}
