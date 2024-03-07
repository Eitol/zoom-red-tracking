package zoom

import (
	"errors"
	"net/http"
)

const DefaultBaseURL = "http://webservices.zoom.red"
const DefaultEndpoint = "/baaszoom/public/canguroazul/getZoomTrackWs?tipo_busqueda=1&web=1&codigo="

var errResponse = errors.New("error response")
var errNotFound = errors.New("not found")

type HttpClient struct {
	baseUrl  string
	endpoint string
	client   *http.Client
}

func NewClientWithURL(baseUrl, endpoint string) HttpClient {
	return HttpClient{
		baseUrl:  baseUrl,
		endpoint: endpoint,
		client: &http.Client{
			Transport: &http.Transport{
				DisableKeepAlives: false,
			},
		},
	}
}

func NewDefaultClient() HttpClient {
	return NewClientWithURL(DefaultBaseURL, DefaultEndpoint)
}
