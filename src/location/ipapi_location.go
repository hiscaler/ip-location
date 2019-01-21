package location

// http://ip-api.com
import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type IpApiLocation struct {
	Location
}

func (t *IpApiLocation) Find() (Location, error) {
	t.Name = "IP-API"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://ip-api.com/json/"+t.Ip+"?lang=zh-CN", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	req.Header.Set("Host", "ip-api.com")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Pragma", "no-cache")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err == nil && resp.StatusCode == 200 {
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			t.rawData = string(body)
			t.data = t.rawData
			respJson := struct {
				As          string
				City        string
				Country     string
				CountryCode string
				Isp         string
				Lat         float64
				Lon         float64
				Query       string
				Region      string
				RegionName  string
				Status      string
				TimeZone    string
				Zip         string
			}{}
			if err := json.Unmarshal([]byte(t.data), &respJson); err == nil {
				if respJson.Status != "fail" {
					t.Success = true
					t.Ip = respJson.Query
					t.CountryId = respJson.CountryCode
					t.CountryName = respJson.Country
					t.ProvinceId = respJson.Region
					t.ProvinceName = respJson.RegionName
					t.CityName = respJson.City
					t.IspName = respJson.Isp

					return t.Location, nil
				} else {
					return t.Location, errors.New("Fail")
				}
			} else {
				return t.Location, err
			}
		} else {
			return t.Location, err
		}
	}

	return t.Location, err
}
