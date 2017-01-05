package concourse

import (
	"golang.org/x/oauth2"
	"net"
	"net/http"
	"time"
)

// https://github.com/concourse/fly/blob/6fb036ef31f6e6f3e74f0089f2d59d2722f0580c/rc/target.go#L378
type basicAuthTransport struct {
	username string
	password string

	base http.RoundTripper
}

func (t basicAuthTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.SetBasicAuth(t.username, t.password)
	return t.base.RoundTrip(r)
}

func newBasicAuthClient(username, password string) *http.Client {
	httpClient := &http.Client{
		Transport: basicAuthTransport{
			username: username,
			password: password,
			base:     defaultTransport(),
		},
	}
	return httpClient
}

func newOAuthClient(tokenType, tokenValue string) *http.Client {
	var oAuthToken *oauth2.Token
	oAuthToken = &oauth2.Token{
		TokenType:   tokenType,
		AccessToken: tokenValue,
	}

	transport := defaultTransport()

	transport = &oauth2.Transport{
		Source: oauth2.StaticTokenSource(oAuthToken),
		Base:   transport,
	}
	return &http.Client{Transport: transport}
}

func defaultTransport() http.RoundTripper {
	var transport http.RoundTripper
	transport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).Dial,
	}

	return transport
}
