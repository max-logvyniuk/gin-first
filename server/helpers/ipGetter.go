package helpers

import (
	"fmt"
	"net"
)

func GetIP() []net.IP {

	ifaces, err := net.Interfaces()

	fmt.Println("ifaces ---", ifaces)

	if err != nil {
		fmt.Println("IP getter error ---", err)
	}

	var IPSlice []net.IP

	// handle err
	for _, i := range ifaces {
		addrs, err := i.Addrs()

		fmt.Println("ifaces addrs ---", addrs)

		if err != nil {
			fmt.Println("IP getter in range error ---", err)
		}

		
		// handle err
		for _, addr := range addrs {
			var ip net.IP

			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			IPSlice = append(IPSlice, ip)
			// process IP address
		}
	}

	return IPSlice

}
