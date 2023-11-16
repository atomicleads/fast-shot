package fastshot

import (
	"encoding/base64"
	"github.com/atomicleads/fast-shot/constant/header"
)

// Auth is the interface that wraps the basic methods for setting authentication configurations.
var _ Auth[RequestBuilder] = (*RequestAuthBuilder)(nil)

// RequestAuthBuilder allows for setting authentication configurations.
type RequestAuthBuilder struct {
	parentBuilder *RequestBuilder
}

// Auth returns a new ClientAuthBuilder for setting authentication options.
func (b *RequestBuilder) Auth() *RequestAuthBuilder {
	return &RequestAuthBuilder{parentBuilder: b}
}

// Set sets the Authorization header for custom authentication.
func (b *RequestAuthBuilder) Set(value string) *RequestBuilder {
	b.parentBuilder.request.httpHeader.Set(header.Authorization, value)
	return b.parentBuilder
}

// BearerToken sets the Authorization header for Bearer token authentication.
func (b *RequestAuthBuilder) BearerToken(token string) *RequestBuilder {
	b.Set("Bearer " + token)
	return b.parentBuilder
}

// BasicAuth sets the Authorization header for Basic authentication.
func (b *RequestAuthBuilder) BasicAuth(username, password string) *RequestBuilder {
	encoded := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	b.Set("Basic " + encoded)
	return b.parentBuilder
}
