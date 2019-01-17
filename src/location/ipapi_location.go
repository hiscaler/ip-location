package location

// http://ip-api.com
import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type IpApiLocation struct {
	Location
}

func (this *IpApiLocation) Find() (*IpApiLocation, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://ip-api.com/json/"+this.Ip+"?lang=zh-CN", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	req.Header.Set("Host", "ip-api.com")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Pragma", "no-cache")
	resp, err := client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			this.rawData = string(body)
			this.data = this.rawData
			t := struct {
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
			if err := json.Unmarshal([]byte(this.data), &t); err == nil {
				this.CountryId = t.CountryCode
				this.CountryName = t.Country
				this.ProvinceId = t.Region
				this.ProvinceName = t.RegionName
				this.CityName = t.City
				this.IspName = t.Isp

				return this, nil
			} else {
				return this, err
			}
		} else {
			return this, err
		}
	}

	return this, err
}
