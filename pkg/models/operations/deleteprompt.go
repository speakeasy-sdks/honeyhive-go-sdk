package operations

import (
	"github.com/speakeasy-sdks/honeyhive-go-sdk/pkg/models/shared"
)

type DeletePromptPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeletePromptRequest struct {
	PathParams DeletePromptPathParams
}

type DeletePromptResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	PromptDeleteResponse  *shared.PromptDeleteResponse
	StatusCode            int
}
