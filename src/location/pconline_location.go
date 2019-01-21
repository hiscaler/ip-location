package location

// whois.pconline.com.cn
import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type PcOnlineLocation struct {
	Location
}

func (t *PcOnlineLocation) Find() (Location, error) {
	t.Name = "PC Online"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://whois.pconline.com.cn/ipJson.jsp?ip="+t.Ip, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	req.Header.Set("Host", "whois.pconline.com.cn")
	req.Header.Set("Pragma", "no-cache")
	resp, err := client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			body, _ := ToUTF8("gb18030", body)
			t.rawData = string(body)
			s := strings.TrimSpace(t.rawData)
			s = strings.Replace(s, "if(window.IPCallBack) {IPCallBack(", "", 1)
			s = strings.Replace(s, ");}", "", 1)
			t.data = s
			respJson := struct {
				Ip          string
				ProCode     string
				Pro         string
				CityCode    string
				City        string
				RegionCode  string
				Region      string
				Addr        string
				RegionNames string
				Err         string
			}{}
			if err := json.Unmarshal([]byte(t.data), &respJson); err == nil {
				if respJson.Err == "" {
					t.Success = true
					t.Ip = respJson.Ip
					t.ProvinceId = respJson.ProCode
					t.ProvinceName = respJson.Pro
					t.CityId = respJson.CityCode
					t.CityName = respJson.City
					t.RegionId = respJson.RegionCode
					t.RegionName = respJson.Region
					if respJson.Addr != "" && strings.Contains(respJson.Addr, " ") {
						addr := strings.Fields(respJson.Addr)
						if len(addr) >= 2 {
							t.Address = addr[0]
							t.IspName = addr[1]
						} else {
							t.Address = respJson.Addr
						}
					} else {
						t.Address = respJson.Addr
					}

					return t.Location, nil
				} else {
					return t.Location, errors.New(respJson.Err)
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
