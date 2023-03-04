package operations

import (
	"github.com/speakeasy-sdks/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type DeleteDatasetPathParams struct {
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

type DeleteDatasetRequest struct {
	PathParams DeleteDatasetPathParams
}

type DeleteDatasetResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	DeleteDatasetResponse *shared.DeleteDatasetResponse
	StatusCode            int
	RawResponse           *http.Response
}
