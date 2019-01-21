package location

// freeapi.ipip.net
import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type IpIpLocation struct {
	Location
}

func (t *IpIpLocation) Find() (Location, error) {
	t.Name = "IPIP"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://freeapi.ipip.net/"+t.Ip, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	req.Header.Set("Host", "freeapi.ipip.net")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Pragma", "no-cache")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err == nil && resp.StatusCode == 200 {
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			t.rawData = string(body)
			t.data = t.rawData
			if !strings.Contains(t.rawData, "not found") {
				respJson := [5]string{}
				if err := json.Unmarshal([]byte(t.data), &respJson); err == nil {
					t.Success = true
					t.CountryName = respJson[0]
					t.ProvinceName = respJson[1]
					t.CityName = respJson[2]
					t.RegionName = respJson[3]
					t.IspName = respJson[4]

					return t.Location, nil
				} else {
					return t.Location, err
				}
			} else {
				return t.Location, errors.New(t.rawData)
			}
		} else {
			return t.Location, err
		}

	}

	return t.Location, err
}
