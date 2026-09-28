# TempestVideo Go SDK

A Go client for the public TempestVideo API.

```sh
go get github.com/tempestvideo/go-sdk
```

```go
import (
	"connectrpc.com/connect"

	"github.com/tempestvideo/go-sdk"
	"github.com/tempestvideo/go-sdk/tempestvideov1"
)

client := tempestvideo.New(os.Getenv("TEMPESTVIDEO_API_KEY"))
res, err := client.Installations.ListInstallations(ctx, connect.NewRequest(&tempestvideov1.ListInstallationsRequest{}))
```

Create an API key in the TempestVideo dashboard (account menu → API Keys). A
key acts as you, with your access. `tempestvideo.WithBaseURL` points the client
at another environment.

## Branches

- `main` follows the production API.
- `development` follows the dev API (`https://api.dev.tempestvideo.net`) and
  may change at any time.

`tempestvideov1/` is generated from the public API definition
(`tempestvideo/v1/api.proto` in the protocol repository) and updated
automatically. Don't edit it by hand; `client.go` and this README are
maintained here.
