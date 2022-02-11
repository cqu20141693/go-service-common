package utils_test

import (
	"fmt"
	"github.com/cqu20141693/go-service-common/v2/utils"
	"testing"
)

func TestGetIp(t *testing.T) {
	ip, err := utils.GetOutBoundIP()
	if err != nil {
		return
	}
	// 192.168.0.123
	fmt.Println("ip:", ip)
}
