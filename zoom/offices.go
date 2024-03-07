package zoom

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func (h *HttpClient) GetOffices() ([]CourierOffice, error) {
	return getOffices()
}

func getOffices() ([]CourierOffice, error) {
	url := "https://zoom.red/wp-admin/admin-ajax.php?action=asl_load_stores&load_all=1"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return loadOfficesFromJson(b)
}

func loadOfficesFromJson(jsonStr []byte) ([]CourierOffice, error) {
	offices := make([]CourierOffice, 250)
	err := json.Unmarshal(jsonStr, &offices)
	if err != nil {
		return nil, err
	}
	for i, o := range offices {
		oo, err := parseOpenHours(o.OpenHours)
		if err != nil {
			return nil, err
		}
		offices[i].OpenHoursMap = *oo
		lat, err := strconv.ParseFloat(o.LatStr, 64)
		if err != nil {
			return nil, err
		}
		lon, err := strconv.ParseFloat(o.LngStr, 64)
		if err != nil {
			return nil, err
		}
		offices[i].Latitude = lat
		offices[i].Longitude = lon
	}
	return offices, nil
}
