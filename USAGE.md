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