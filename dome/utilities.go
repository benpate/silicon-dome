package dome

import (
	"net"
	"net/http"
	"strings"

	"github.com/benpate/domain"
)

// Some notes on looking up real IP addresses:
// https://adam-p.ca/blog/2022/03/x-forwarded-for/

// realIPAddress returns the real IP address of the request,
// taking into account X-Real-IP and X-Forwarded-For headers.
func realIPAddress(request *http.Request) string {

	// Get IP address from CloudFlare (this is more trustworthy than other headers)
	if cfConnectingIP := request.Header.Get("CF-Connecting-IP"); cfConnectingIP != "" {
		return cfConnectingIP
	}

	// Get the "true client ip" for Akamai
	if trueClientIP := request.Header.Get("True-Client-IP"); trueClientIP != "" {
		return trueClientIP
	}

	// Get IP address from X-Forwarded-For header
	if forwardedFor := request.Header.Get("X-Forwarded-For"); forwardedFor != "" {

		// Scan for first non-local ip address
		for _, ip := range strings.Split(forwardedFor, ",") {
			ip = strings.TrimSpace(ip)
			if !domain.IsLocalhost(ip) {
				return ip
			}
		}
	}

	// Get IP address from X-Real-IP header
	if realIP := request.Header.Get("X-Real-Ip"); realIP != "" {
		return realIP
	}

	// Get IP address from request.RemoteAddr
	result, _, _ := net.SplitHostPort(request.RemoteAddr)
	return result
}
