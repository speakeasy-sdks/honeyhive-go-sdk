package operations

import (
	"github.com/speakeasy-sdks/honeyhive-go-sdk/pkg/models/shared"
)

type ChangePromptPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type ChangePromptRequest struct {
	PathParams ChangePromptPathParams
	Request    shared.Prompt `request:"mediaType=application/json"`
}

type ChangePromptResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	Prompt                *shared.Prompt
	StatusCode            int
}
