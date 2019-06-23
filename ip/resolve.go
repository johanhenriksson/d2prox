package ip

import (
	"io/ioutil"
	"net"
	"net/http"
)

var publicIP = net.IP{127, 0, 0, 1}

// ResolvePublicIP resolves & caches the machines public ip address using ipify.org
func ResolvePublicIP() (net.IP, error) {
	res, err := http.Get("https://api.ipify.org")
	if err != nil {
		return nil, err
	}
	ipstr, err := ioutil.ReadAll(res.Body)
	if err == nil {
		publicIP = net.ParseIP(string(ipstr)).To4()
	}
	return publicIP, err
}

// Public returns the latest known public ip. If its not yet resolved, loopback will be returned.
func Public() net.IP {
	return publicIP
}

// Loopback returns the loopback ip address
func Loopback() net.IP {
	return net.IP{127, 0, 0, 1}
}
