// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateFeedbackResponse struct {
	// Bad Request error
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	// Successful response
	Feedback    *shared.Feedback
	StatusCode  int
	RawResponse *http.Response
}
