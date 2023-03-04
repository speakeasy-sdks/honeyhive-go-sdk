package operations

import (
	"github.com/speakeasy-sdks/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type FindFineTunedModelsRequest struct {
	Request *shared.FineTunedModelsInput `request:"mediaType=application/json"`
}

type FindFineTunedModelsResponse struct {
	APIResponseBadRequest   *shared.APIResponseBadRequest
	ContentType             string
	FineTunedModelsResponse *shared.FineTunedModelsResponse
	StatusCode              int
	RawResponse             *http.Response
}
