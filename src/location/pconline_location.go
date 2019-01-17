package location

// whois.pconline.com.cn
import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type PcOnlineLocation struct {
	Location
}

func (this *PcOnlineLocation) Find() (*PcOnlineLocation, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://whois.pconline.com.cn/ipJson.jsp", strings.NewReader("ip="+this.Ip))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	req.Header.Set("Host", "whois.pconline.com.cn")
	req.Header.Set("Pragma", "no-cache")
	resp, err := client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			body, _ := ToUTF8("gb18030", body)
			this.rawData = string(body)
			s := strings.TrimSpace(this.rawData)
			s = strings.Replace(s, "if(window.IPCallBack) {IPCallBack(", "", 1)
			s = strings.Replace(s, ");}", "", 1)
			this.data = s
			t := struct {
				Ip          string
				ProCode     string
				Pro         string
				CityCode    string
				City        string
				RegionCode  string
				Region      string
				Addr        string
				RegionNames string
			}{}
			if err := json.Unmarshal([]byte(this.data), &t); err == nil {
				this.ProvinceId = t.ProCode
				this.ProvinceName = t.Pro
				this.CityId = t.CityCode
				this.CityName = t.City
				this.RegionId = t.RegionCode
				this.RegionName = t.Region

				if t.Addr != "" && strings.Contains(t.Addr, " ") {
					addr := strings.Fields(t.Addr)
					if len(addr) >= 2 {
						this.Address = addr[0]
						this.IspName = addr[1]
					} else {
						this.Address = t.Addr
					}
				} else {
					this.Address = t.Addr
				}
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
