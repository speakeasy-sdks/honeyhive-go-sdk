package operations

import (
	"github.com/speakeasy-sdks/honeyhive-go-sdk/pkg/models/shared"
)

type GetGenerationRequest struct {
	Request shared.GenerationInput `request:"mediaType=application/json"`
}

type GetGenerationResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	GenerationOutputs     *shared.GenerationOutputs
	StatusCode            int
}
