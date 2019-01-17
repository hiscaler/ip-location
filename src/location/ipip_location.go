package location

// freeapi.ipip.net
import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type IpIpLocation struct {
	Location
}

func (this *IpIpLocation) Find() (*IpIpLocation, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://freeapi.ipip.net/"+this.Ip, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	req.Header.Set("Host", "freeapi.ipip.net")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Pragma", "no-cache")
	resp, err := client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			this.rawData = string(body)
			this.data = this.rawData
			t := [5]string{}
			if err := json.Unmarshal([]byte(this.data), &t); err == nil {
				this.CountryName = t[0]
				this.ProvinceName = t[1]
				this.CityName = t[2]
				this.RegionName = t[3]
				this.IspName = t[4]

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
