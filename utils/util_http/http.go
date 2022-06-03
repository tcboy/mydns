package util_http

import (
	"net/http"
	"regexp"
	"strings"
)

var (
	SimpleIpPattern  = regexp.MustCompile("^[0-9]+.[0-9]+.[0-9]+.[0-9]+$")
	PrivateIpPattern = regexp.MustCompile("^192.168.|^10.|^172.(1[6-9]|2[0-9]|3[01]).")
)

func GetRealIp(req *http.Request) string {
	ipsStr := req.Header.Get("Ali-CDN-Real-IP")
	if ipsStr == "" {
		ipsStr = req.Header.Get("X-Real-IP")
	}

	if ipsStr == "" {
		ipsStr = req.Header.Get("HTTP_X_FORWARDED_FOR")
	}

	if ipsStr != "" {
		ips := strings.Split(ipsStr, ",")
		for _, ip := range ips {
			ip = strings.TrimSpace(ip)
			if SimpleIpPattern.MatchString(ip) && !PrivateIpPattern.MatchString(ip) {
				return ip
			}
		}
		return "0.0.0.0"
	} else if remoteAddr := req.RemoteAddr; remoteAddr != "" {
		ipPort := strings.Split(remoteAddr, ":")
		return ipPort[0]
	}
	return ""
}

func ReplaceUrls(old, new string, urls ...string) {
	for i, rawurl := range urls {
		urls[i] = strings.Replace(rawurl, old, new, -1)
	}
}

func SchemeToHttps(urls ...string) {
	ReplaceUrls("http://", "https://", urls...)
}
