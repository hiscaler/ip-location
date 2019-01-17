package main

import (
	"fmt"
	"location"
)

// IP 地址信息查询
func main() {
	ip := location.PcOnlineLocation{}
	ip.SetIp("47.75.37.108")
	if v, err := ip.Find(); err == nil {
		fmt.Println(fmt.Sprintf("%#v", v))
	} else {
		fmt.Println("Error", err)
	}
}
