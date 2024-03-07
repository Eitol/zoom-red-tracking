package zoom

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

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
