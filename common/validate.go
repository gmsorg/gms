package common

import (
	"errors"
	"net"
)

func ValidateIp(ip string) error {
	parseIP := net.ParseIP(ip)
	if parseIP == nil {
		return errors.New("ip error")
	}
	return nil

}

func ValidatePort(port int) error {
	if port < 1024 {
		return errors.New("port less than 1024")
	}
	if port >= 65535 {
		return errors.New("port big than 65535")
	}
	return nil
}
