package zoom

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
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

func (h *HttpClient) GetTrackingInfo(tracking int) (*Shipment, error) {
	url := h.baseUrl + h.endpoint + strconv.Itoa(tracking)
	method := "GET"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Connection", "keep-alive")
	res, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	respObj := GetTrackingResponse{}
	err = json.Unmarshal(body, &respObj)
	if err != nil {
		return nil, err
	}
	if respObj.Mensaje == "INFORMACION NO EXISTE EN BASE DE DATOS" {
		return nil, errNotFound
	}
	if respObj.Mensaje != "CONSULTA REALIZADA EXITOSAMENTE" {
		return nil, errors.Join(errResponse, errors.New(respObj.Mensaje))
	}
	return &respObj.Shipment, nil
}
