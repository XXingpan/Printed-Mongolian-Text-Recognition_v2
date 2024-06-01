package utils

import (
	"net"
)

func HostIP() string {
	ip := "0.0.0.0"
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		hostname, _ := net.LookupHost("")
		ip = hostname[0]
	} else {
		defer conn.Close()
		localAddr := conn.LocalAddr().(*net.UDPAddr)
		ip = localAddr.IP.String()
	}

	return ip
}

//func main() {
//	fmt.Println(hostIP())
//}
