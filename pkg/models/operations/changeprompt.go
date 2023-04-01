// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type ChangePromptRequest struct {
	// Prompt to add to task
	Prompt shared.Prompt `request:"mediaType=application/json"`
	// The id of the prompt
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type ChangePromptResponse struct {
	// Bad Request error
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	// Successful response
	Prompt      *shared.Prompt
	StatusCode  int
	RawResponse *http.Response
}
