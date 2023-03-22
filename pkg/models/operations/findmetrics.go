// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type FindMetricsRequest struct {
	// Find metrics by task name.
	Request *shared.FindMetricsInput `request:"mediaType=application/json"`
}

type FindMetricsResponse struct {
	// Bad Request error
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	// Successful response
	FindMetricsResponse []shared.MetricsResponse
	StatusCode          int
	RawResponse         *http.Response
}
