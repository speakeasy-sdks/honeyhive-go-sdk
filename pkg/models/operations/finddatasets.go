package operations

import (
	"github.com/speakeasy-sdks/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type FindDatasetsRequest struct {
	Request *shared.FindDatasetsInput `request:"mediaType=application/json"`
}

type FindDatasetsResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	FindDatasetsResponse  *shared.FindDatasetsResponse
	StatusCode            int
	RawResponse           *http.Response
}
