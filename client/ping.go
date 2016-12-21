package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// AnswerPingPingPath computes a request path to the AnswerPing action of ping.
func AnswerPingPingPath(args string) string {
	return fmt.Sprintf("/wifidog/ping/%v", args)
}

// answer wifidog ping GET-Request
func (c *Client) AnswerPingPing(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewAnswerPingPingRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAnswerPingPingRequest create the request corresponding to the AnswerPing action endpoint of the ping resource.
func (c *Client) NewAnswerPingPingRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
