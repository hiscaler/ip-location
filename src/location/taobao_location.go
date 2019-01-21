package location

// ip.taobao.com
import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type TaoBaoLocation struct {
	Location
}

func (t *TaoBaoLocation) Find() (Location, error) {
	t.Name = "TaoBao"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://ip.taobao.com/service/getIpInfo.php?ip="+t.Ip, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	req.Header.Set("Host", "ip.taobao.com")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Pragma", "no-cache")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err == nil && resp.StatusCode == 200 {
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			respJson := struct {
				Code int
				Data struct {
					Ip        string
					IspId     string `json:"isp_id"`
					Isp       string
					CountryId string `json:"country_id"`
					Country   string
					AreaId    string `json:"area_id"`
					Area      string
					RegionId  string `json:"region_id"`
					Region    string
					CityId    string `json:"city_id"`
					City      string
					CountyId  string `json:"county_id"`
					County    string
				}
			}{}
			if err := json.Unmarshal(body, &respJson); err == nil {
				parseValue := func(s string) string {
					s = strings.ToLower(s)
					if s == "xx" {
						return ""
					} else {
						return s
					}
				}
				if respJson.Code == 0 {
					t.CountryId = parseValue(respJson.Data.CountryId)
					t.CountryName = parseValue(respJson.Data.Country)
					t.AreaId = parseValue(respJson.Data.AreaId)
					t.AreaName = parseValue(respJson.Data.Area)
					t.ProvinceId = parseValue(respJson.Data.RegionId)
					t.ProvinceName = parseValue(respJson.Data.Region)
					t.CityId = parseValue(respJson.Data.CityId)
					t.CityName = parseValue(respJson.Data.City)
					t.RegionId = parseValue(respJson.Data.CountyId)
					t.RegionName = parseValue(respJson.Data.County)
					t.IspId = parseValue(respJson.Data.IspId)
					t.IspName = parseValue(respJson.Data.Isp)

					return t.Location, nil
				} else {
					return t.Location, errors.New(string(respJson.Code))
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
