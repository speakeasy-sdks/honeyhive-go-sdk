package honeyhive

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/honeyhive-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/honeyhive-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/honeyhive-go-sdk/pkg/utils"
	"net/http"
	"strings"
)

type fineTunedModel struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newFineTunedModel(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *fineTunedModel {
	return &fineTunedModel{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// FindFineTunedModels - This endpoint gets the fine-tuned models being managed by the user.
// This endpoint gets the fine-tuned models associated with a particular model or task.
func (s *fineTunedModel) FindFineTunedModels(ctx context.Context, request operations.FindFineTunedModelsRequest) (*operations.FindFineTunedModelsResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/fine_tuned_models"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.FindFineTunedModelsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.FineTunedModelsResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.FineTunedModelsResponse = out
		}
	case httpRes.StatusCode == 400:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.APIResponseBadRequest
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.APIResponseBadRequest = out
		}
	}

	return res, nil
}
