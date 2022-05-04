package address

import (
	"github.com/pkg/errors"
	"log"
	"net"
)

//find the address for the host
func FindAddress() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	// handle err
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Println(err)
			continue
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
			// process IP address
			if ip == nil {
				continue
			}
			//ipString = ip.String()
			if ip.IsGlobalUnicast() {
				return ip.String(), nil
			}

		}
	}

	return "", errors.New("could not find appropriate ip")

}
