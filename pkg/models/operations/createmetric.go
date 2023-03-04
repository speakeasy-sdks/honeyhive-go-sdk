package operations

import (
	"github.com/speakeasy-sdks/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateMetricRequest struct {
	Request shared.Metric `request:"mediaType=application/json"`
}

type CreateMetricResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	Metric                *shared.Metric
	StatusCode            int
	RawResponse           *http.Response
}
