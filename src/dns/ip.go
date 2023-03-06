package dns

import (
	"io/ioutil"
	"net"
	"net/http"
)

func ResolveDomain(domain string) ([]string, error) {
	ips, err := net.LookupIP(domain)

	if err != nil {
		return nil, err
	}

	var ipArr []string

	for _, ip := range ips {
		ipArr = append(ipArr, ip.String())
	}

	return ipArr, nil
}

func getIPFromAPI(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func GetPublicIP() string {
	results := make(chan string, 2)

	go func() {
		ip, err := getIPFromAPI("https://api.ipify.org")

		if err != nil {
			results <- ""
		} else {
			results <- ip
		}
	}()

	go func() {
		ip, err := getIPFromAPI("https://api.ipify.org")

		if err != nil {
			results <- ""
		} else {
			results <- ip
		}
	}()

	var ips []string

	for i := 0; i < 2; i++ {
		ip := <-results

		if ip != "" {
			ips = append(ips, ip)
		}
	}

	if len(ips) == 0 {
		return ""
	}

	for _, ip := range ips {
		if ip != ips[0] {
			return ""
		}
	}

	return ips[0]
}
