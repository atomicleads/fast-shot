package fastshot

import (
	"errors"
	"github.com/atomicleads/fast-shot/constant"
	"net/http"
	"net/url"
	"time"
)

// ClientConfig is the interface that wraps the basic methods for setting HTTP Client configuration.
var _ ClientConfig[ClientBuilder] = (*ClientConfigBuilder)(nil)

// ClientConfigBuilder allows for setting other client configurations.
type ClientConfigBuilder struct {
	parentBuilder *ClientBuilder
}

// Config returns a new ClientConfigBuilder for setting custom client configurations.
func (b *ClientBuilder) Config() *ClientConfigBuilder {
	return &ClientConfigBuilder{parentBuilder: b}
}

// SetCustomTransport sets custom transport for the HTTP client.
func (b *ClientConfigBuilder) SetCustomTransport(transport http.RoundTripper) *ClientBuilder {
	b.parentBuilder.client.httpClient.Transport = transport
	return b.parentBuilder
}

// SetTimeout sets the timeout for the HTTP client.
func (b *ClientConfigBuilder) SetTimeout(duration time.Duration) *ClientBuilder {
	b.parentBuilder.client.httpClient.Timeout = duration
	return b.parentBuilder
}

// SetFollowRedirects controls whether the HTTP client should follow redirects.
func (b *ClientConfigBuilder) SetFollowRedirects(follow bool) *ClientBuilder {
	if !follow {
		b.parentBuilder.client.httpClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}
	return b.parentBuilder
}

// SetProxy sets the proxy URL for the HTTP client.
func (b *ClientConfigBuilder) SetProxy(proxyURL string) *ClientBuilder {
	parsedURL, err := url.Parse(proxyURL)
	if err != nil {
		b.parentBuilder.client.validations = append(b.parentBuilder.client.validations, errors.Join(errors.New(constant.ErrMsgParseProxyURL), err))
		return b.parentBuilder
	}

	if transport, ok := b.parentBuilder.client.httpClient.Transport.(*http.Transport); ok {
		transport.Proxy = http.ProxyURL(parsedURL)
	} else {
		b.parentBuilder.client.httpClient.Transport = &http.Transport{
			Proxy: http.ProxyURL(parsedURL),
		}
	}

	return b.parentBuilder
}
