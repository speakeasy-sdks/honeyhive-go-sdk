package honeyhive

import (
	"github.com/speakeasy-sdks/honeyhive-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/honeyhive-go-sdk/pkg/utils"
	"net/http"
	"time"
)

var ServerList = []string{
	"https://honeyhive.ai:5000/",
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// SDK Documentation: https://docs.honeyhive.ai/api-reference/authentication - HoneyHive Docs
type Honeyhive struct {
	Dataset        *dataset
	Feedback       *feedback
	FineTunedModel *fineTunedModel
	Generation     *generation
	Metric         *metric
	Prompt         *prompt
	Task           *task

	// Non-idiomatic field names below are to namespace fields from the fields names above to avoid name conflicts
	_defaultClient  HTTPClient
	_securityClient HTTPClient
	_security       *shared.Security
	_serverURL      string
	_language       string
	_sdkVersion     string
	_genVersion     string
}

type SDKOption func(*Honeyhive)

func WithServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Honeyhive) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk._serverURL = serverURL
	}
}

func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Honeyhive) {
		sdk._defaultClient = client
	}
}

func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *Honeyhive) {
		sdk._security = &security
	}
}

func New(opts ...SDKOption) *Honeyhive {
	sdk := &Honeyhive{
		_language:   "go",
		_sdkVersion: "0.2.0",
		_genVersion: "1.5.0",
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk._defaultClient == nil {
		sdk._defaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk._securityClient == nil {

		if sdk._security != nil {
			sdk._securityClient = utils.ConfigureSecurityClient(sdk._defaultClient, sdk._security)
		} else {
			sdk._securityClient = sdk._defaultClient
		}

	}

	if sdk._serverURL == "" {
		sdk._serverURL = ServerList[0]
	}

	sdk.Dataset = newDataset(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Feedback = newFeedback(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.FineTunedModel = newFineTunedModel(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Generation = newGeneration(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Metric = newMetric(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Prompt = newPrompt(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Task = newTask(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	return sdk
}
