# HoneyHive Go SDK

<div align="center">
   <img src="https://user-images.githubusercontent.com/6267663/220803812-cd7e27bd-06cb-49b0-87c1-d85fe21a3557.png" />
   <p>HoneyHive is a model observability and evaluation platform that helps you continuously improve your models in production. We help you evaluate, deploy, monitor and fine-tune both closed and open-source large language models for production use-cases, allowing you to optimize model performance & align your models with your users’ preferences.</p>
   <a href="https://github.com/speakeasy-sdks/honeyhive-go-sdk/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/honeyhive-go-sdk/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
   <a href="https://docs.honeyhive.ai/introduction"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=2ca47c&style=for-the-badge" /></a>
</div> 

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/honeyhive-go-sdk
```
<!-- End SDK Installation -->

## Authentication

After signing up on the app, you can find your API key in the [Settings page](https://app.honeyhive.ai/settings/account).

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
    "log"
    "github.com/speakeasy-sdks/honeyhive-go-sdk"
    "github.com/speakeasy-sdks/honeyhive-go-sdk/pkg/models/shared"
    "github.com/speakeasy-sdks/honeyhive-go-sdk/pkg/models/operations"
)

func main() {
    opts := []honeyhive.SDKOption{
        honeyhive.WithSecurity(
            shared.Security{
                BearerAuth: shared.SchemeBearerAuth{
                    Authorization: "Bearer YOUR_BEARER_TOKEN_HERE",
                },
            }
        ),
    }

    s := honeyhive.New(opts...)
    
    req := operations.CreateDatasetRequest{
        Request: shared.Dataset{
            Description: "unde",
            File: "deserunt",
            Name: "porro",
            Purpose: "nulla",
            Task: "id",
        },
    }
    
    res, err := s.Dataset.CreateDataset(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateDatasetResponse != nil {
        // handle response
    }
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## SDK Available Operations


### Dataset

* `CreateDataset` - Create Dataset
* `DeleteDataset` - Delete Datasets
* `FindDatasets` - Get Datasets

### Feedback

* `CreateFeedback` - Create Feedback

### FineTunedModel

* `FindFineTunedModels` - This endpoint gets the fine-tuned models being managed by the user.

### Generation

* `GetGeneration` - Get Generations
* `IngestGenerations` - Create Generation for Task
* `ModelPromptCreateGeneration` - Create Generation for Model and Prompt
* `TaskCreateGeneration` - Create Generation for Task

### Metric

* `CreateMetric` - Create Metric
* `FindMetrics` - Get Metrics

### Prompt

* `ChangePrompt` - Update Prompts
* `CreatePrompt` - Create Prompt
* `DeletePrompt` - Delete Prompts
* `FindPrompts` - Get Prompts

### Task

* `CreateTask` - Create a new task
* `FindAllTasks` - Find all Tasks
* `UpdateTask` - Update an existing task
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)