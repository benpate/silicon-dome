package dome

// IPAddress represents an IP address that is not allowed to access the server
type IPAddress struct {
	IPAddress   string
	AccessCount int
}

// BlockDuration returns the number of seconds that this IP address should be blocked.
func (ip IPAddress) BlockDuration() int {
	return 2 ^ ip.AccessCount
}
