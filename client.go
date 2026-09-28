// Package tempestvideo is a client for the public TempestVideo API.
//
//	client := tempestvideo.New(os.Getenv("TEMPESTVIDEO_API_KEY"))
//	res, err := client.Installations.ListInstallations(ctx, connect.NewRequest(&tempestvideov1.ListInstallationsRequest{}))
//
// The request and response types are in package tempestvideov1, generated
// from the API definition.
package tempestvideo

import (
	"context"
	"net/http"

	"connectrpc.com/connect"

	"github.com/tempestvideo/go-sdk/tempestvideov1/tempestvideov1connect"
)

// DefaultBaseURL is the production API.
const DefaultBaseURL = "https://api.tempestvideo.net"

// Client has one field per API service.
type Client struct {
	Profile       tempestvideov1connect.ProfileServiceClient
	Tenants       tempestvideov1connect.TenantsServiceClient
	Installations tempestvideov1connect.InstallationsServiceClient
	Members       tempestvideov1connect.MembersServiceClient
	Connectors    tempestvideov1connect.ConnectorsServiceClient
	Channels      tempestvideov1connect.ChannelsServiceClient
	Devices       tempestvideov1connect.DevicesServiceClient
	APIKeys       tempestvideov1connect.ApiKeysServiceClient
}

type options struct {
	baseURL    string
	httpClient connect.HTTPClient
	connect    []connect.ClientOption
}

// Option configures New.
type Option func(*options)

// WithBaseURL points the client at another API, e.g.
// https://api.dev.tempestvideo.net.
func WithBaseURL(url string) Option { return func(o *options) { o.baseURL = url } }

// WithHTTPClient replaces http.DefaultClient.
func WithHTTPClient(c connect.HTTPClient) Option { return func(o *options) { o.httpClient = c } }

// WithClientOptions passes options to every service client, e.g.
// connect.WithInterceptors.
func WithClientOptions(opts ...connect.ClientOption) Option {
	return func(o *options) { o.connect = append(o.connect, opts...) }
}

// New returns a client that authenticates with apiKey. Create API keys in the
// dashboard (account menu, API Keys); a key acts as you, with your access.
func New(apiKey string, opts ...Option) *Client {
	o := options{baseURL: DefaultBaseURL, httpClient: http.DefaultClient}
	for _, opt := range opts {
		opt(&o)
	}
	auth := connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			req.Header().Set("Authorization", "Bearer "+apiKey)
			return next(ctx, req)
		}
	})
	c := append([]connect.ClientOption{connect.WithInterceptors(auth)}, o.connect...)
	return &Client{
		Profile:       tempestvideov1connect.NewProfileServiceClient(o.httpClient, o.baseURL, c...),
		Tenants:       tempestvideov1connect.NewTenantsServiceClient(o.httpClient, o.baseURL, c...),
		Installations: tempestvideov1connect.NewInstallationsServiceClient(o.httpClient, o.baseURL, c...),
		Members:       tempestvideov1connect.NewMembersServiceClient(o.httpClient, o.baseURL, c...),
		Connectors:    tempestvideov1connect.NewConnectorsServiceClient(o.httpClient, o.baseURL, c...),
		Channels:      tempestvideov1connect.NewChannelsServiceClient(o.httpClient, o.baseURL, c...),
		Devices:       tempestvideov1connect.NewDevicesServiceClient(o.httpClient, o.baseURL, c...),
		APIKeys:       tempestvideov1connect.NewApiKeysServiceClient(o.httpClient, o.baseURL, c...),
	}
}
