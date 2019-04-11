package ip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

type dnsResponse struct {
	Answer []dnsAnswer
}

type dnsAnswer struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

var dnsCache = map[string]net.IP{}

// ResolveHost resolves a hostname using an external DNS service
func ResolveHost(hostname string) (net.IP, error) {
	// check cache
	if ip, exists := dnsCache[hostname]; exists {
		return ip, nil
	}

	// query googles dns service
	url := fmt.Sprintf("https://dns.google.com/resolve?name=%s&type=A", hostname)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)

	// unmarshal json response
	response := dnsResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	// ensure we've got something
	if len(response.Answer) == 0 {
		return nil, fmt.Errorf("No results")
	}

	// convert to net.IP
	ipstr := response.Answer[0].Data
	ip := net.ParseIP(string(ipstr))

	// cache result
	dnsCache[hostname] = ip

	return ip, nil
}
