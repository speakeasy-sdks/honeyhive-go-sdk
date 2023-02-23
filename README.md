# github.com/speakeasy-sdks/honeyhive-go-sdk

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/honeyhive-go-sdk
```
<!-- End SDK Installation -->

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
