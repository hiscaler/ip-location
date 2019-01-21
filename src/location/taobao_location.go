package location

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TaoBaoLocation struct {
	Location
}

func (t *TaoBaoLocation) Find() (Location, error) {
	fmt.Println("http://ip.taobao.com/service/getIpInfo.php?ip=" + t.Ip)
	resp, err := http.Get("http://ip.taobao.com/service/getIpInfo.php?ip=" + t.Ip)
	if err == nil {
		defer resp.Body.Close()
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
				if respJson.Code == 0 {
					t.CountryId = respJson.Data.CountryId
					t.CountryName = respJson.Data.Country
					t.AreaId = respJson.Data.AreaId
					t.AreaName = respJson.Data.Area
					t.ProvinceId = respJson.Data.RegionId
					t.ProvinceName = respJson.Data.Region
					t.CityId = respJson.Data.CityId
					t.CityName = respJson.Data.City
					t.RegionId = respJson.Data.CountyId
					t.RegionName = respJson.Data.County
					t.IspId = respJson.Data.IspId
					t.IspName = respJson.Data.Isp

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
	} else {
		return t.Location, err
	}
	return t.Location, err
}
