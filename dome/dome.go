package dome

import (
	"net/http"

	"github.com/benpate/data"
	"github.com/benpate/derp"
	"github.com/cloudflare/ahocorasick"
	"github.com/maypok86/otter"
)

// On advice from Gopher Academy, Silicon Dome uses Aho-Corasick string matching to block user agents.
// https://blog.gopheracademy.com/advent-2014/string-matching/
// https://github.com/cloudflare/ahocorasick

// Dome object contains the matcher that is used to identify blocked user agents.
type Dome struct {
	blockedUserAgents *ahocorasick.Matcher
	blockedPaths      *ahocorasick.Matcher
	blockedIPs        otter.CacheWithVariableTTL[string, IPAddress]
	logDatabase       data.Collection
	blockNotFound     bool
}

// New returns a fully initialized Dome object.
func New(options ...Option) Dome {
	result := Dome{}

	CacheCapacity(1024)(&result)

	result.With(options...)
	return result
}

// With applies the provided options to the Dome object.
func (dome *Dome) With(options ...Option) {
	for _, option := range options {
		option(dome)
	}
}

// VerifyHeader verifies the returns TRUE if the provided user agent is blocked (not allowed).
func (dome *Dome) VerifyHeader(header http.Header) error {

	// Check UserAgents
	userAgent := header.Get("User-Agent")

	if userAgent == "" {
		return derp.NewForbiddenError("dome.VerifyHeader", "User Agent must not be empty")
	}

	if dome.blockedUserAgents != nil {
		if dome.blockedUserAgents.Contains([]byte(userAgent)) {
			return derp.NewForbiddenError("dome.VerifyHeader", "User Agent is blocked", userAgent)
		}
	}

	// Check Paths
	if dome.blockedPaths != nil {
		path := header.Get("Path")
		if dome.blockedPaths.Contains([]byte(path)) {
			return derp.NewForbiddenError("dome.VerifyHeader", "Path is blocked", path)
		}
	}

	return nil
}

func (d *Dome) Close() {
	d.blockedIPs.Close()
}
