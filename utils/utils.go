package utils

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

func ToJSONString(o interface{}) string {

	marshal, err := json.Marshal(o)
	if err != nil {
		return ""
	}
	return string(marshal)
}
