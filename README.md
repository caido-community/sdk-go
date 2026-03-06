# sdk-go

Community Go SDK for [Caido](https://caido.io) - the lightweight web security auditing toolkit.

This SDK mirrors the API surface of the official [JavaScript SDK](https://github.com/caido/sdk-js) (`@caido/sdk-client`) and uses [genqlient](https://github.com/Khan/genqlient) for type-safe GraphQL code generation from the official Caido schema.

## Installation

```bash
go get github.com/caido-community/sdk-go
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    caido "github.com/caido-community/sdk-go"
)

func main() {
    ctx := context.Background()

    client, err := caido.NewClient(caido.Options{
        URL:  "http://localhost:8080",
        Auth: caido.PATAuth("your-pat-token"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if err := client.Connect(ctx); err != nil {
        log.Fatal(err)
    }

    // List proxied requests
    first := 10
    resp, err := client.Requests.List(ctx, &caido.ListRequestsOptions{
        First: &first,
    })
    if err != nil {
        log.Fatal(err)
    }

    for _, edge := range resp.Requests.Edges {
        req := edge.Node
        status := 0
        if req.Response != nil {
            status = req.Response.StatusCode
        }
        fmt.Printf("%s %s%s -> %d\n", req.Method, req.Host, req.Path, status)
    }
}
```

## Authentication

The SDK supports Personal Access Tokens (PAT), which is the recommended method:

```go
client, err := caido.NewClient(caido.Options{
    URL:  "http://localhost:8080",
    Auth: caido.PATAuth("caido_xxxxx"),
})
```

You can also use access/refresh token pairs from the OAuth device flow:

```go
client, err := caido.NewClient(caido.Options{
    URL:  "http://localhost:8080",
    Auth: caido.TokenAuth(accessToken, refreshToken),
})
```

## Domain SDKs

The client exposes domain-specific SDKs matching the JS SDK:

| SDK | Description |
|-----|-------------|
| `client.Requests` | Proxied HTTP requests (list, get, metadata) |
| `client.Intercept` | MITM intercept entries and message queue |
| `client.Replay` | Replay sessions, entries, and send requests |
| `client.Findings` | Security findings attached to requests |
| `client.Scopes` | Target scope management |
| `client.Projects` | Project management |
| `client.Environments` | Variable environments |
| `client.HostedFiles` | Files served by Caido |
| `client.Workflows` | Automation workflows |
| `client.Tasks` | Background task management |
| `client.Instance` | Runtime info and settings |
| `client.Filters` | Saved HTTPQL filter presets |
| `client.Users` | Current user info |
| `client.Plugins` | Installed plugin packages |
| `client.Automate` | Fuzzing sessions (Automate) |
| `client.Sitemap` | Site structure tree |

## Readiness Polling

Wait for the Caido instance to be ready before making requests:

```go
err := client.ConnectWithOptions(ctx, caido.ConnectOptions{
    WaitForReady:  true,
    ReadyTimeout:  60 * time.Second,
    ReadyInterval: 2 * time.Second,
})
```

## Low-Level GraphQL Access

For operations not yet covered by domain SDKs, use the GraphQL client directly:

```go
import gen "github.com/caido-community/sdk-go/graphql"

resp, err := gen.ListScopes(ctx, client.GraphQL)
```

## Schema Updates

The GraphQL schema is vendored from [`@caido/schema-proxy`](https://www.npmjs.com/package/@caido/schema-proxy) on npm.

```bash
make schema    # Pull latest schema
make generate  # Regenerate Go code
make build     # Verify compilation
```

## Development

```bash
git clone https://github.com/caido-community/sdk-go.git
cd sdk-go
make check     # generate + build + vet
make test      # run tests
```

## License

MIT - see [LICENSE](LICENSE).

## Acknowledgments

- [Caido](https://caido.io) team for the platform and schema
- Built with [genqlient](https://github.com/Khan/genqlient) for type-safe GraphQL
