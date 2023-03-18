<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/speakeasy-sdks/honeyhive-go-sdk"
    "github.com/speakeasy-sdks/honeyhive-go-sdk/pkg/models/shared"
    "github.com/speakeasy-sdks/honeyhive-go-sdk/pkg/models/operations"
)

func main() {
    s := honeyhive.New(
        honeyhive.WithSecurity(shared.Security{
            BearerAuth: shared.SchemeBearerAuth{
                Authorization: "Bearer YOUR_BEARER_TOKEN_HERE",
            },
        }),
    )

    req := operations.CreateDatasetRequest{
        Request: shared.Dataset{
            Description: "code snippet",
            File: "????",
            Name: "my task",
            Purpose: "new metric",
            Task: "code snippet",
        },
    }

    ctx := context.Background()
    res, err := s.Dataset.CreateDataset(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateDatasetResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->