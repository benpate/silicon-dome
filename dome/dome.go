package dome

import (
	"net/http"

	"github.com/benpate/data"
	"github.com/benpate/derp"
	"github.com/cloudflare/ahocorasick"
)

// On advice from Gopher Academy, Silicon Dome uses Aho-Corasick string matching to block user agents.
// https://blog.gopheracademy.com/advent-2014/string-matching/
// https://github.com/cloudflare/ahocorasick

// Dome object contains the matcher that is used to identify blocked user agents.
type Dome struct {
	blockedUserAgents *ahocorasick.Matcher
	blockedPaths      *ahocorasick.Matcher
	// blockedIPs        otter.CacheWithVariableTTL[string, int]
	// blockOnError bool
	logDatabase data.Collection
}

// New returns a fully initialized Dome object.
func New(options ...Option) Dome {

	/*
		cache, err := createCache(1024)

		if err != nil {
			panic(err)
		}
	*/

	result := Dome{
		// blockedIPs: cache,
	}

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
func (dome *Dome) VerifyRequest(request *http.Request) error {

	/*/ Check blocked IP addresses
	spew.Dump("Blocked IP addresses::")
	dome.blockedIPs.Range(func(key string, value int) bool {
		spew.Dump(key, value)
		return true
	})

	if count, _ := dome.blockedIPs.Get(request.RemoteAddr); count > 3 {
		return derp.NewForbiddenError("dome.VerifyHeader", "Blocked due to previous scanning activity.  Try again later.", request.RemoteAddr)
	}*/

	// Check UserAgents
	userAgent := request.Header.Get("User-Agent")

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
		path := request.URL.Path
		if dome.blockedPaths.Contains([]byte(path)) {
			return derp.NewForbiddenError("dome.VerifyHeader", "Path is blocked", path)
		}
	}

	return nil
}

func (d *Dome) HandleError(request *http.Request, err error) {

	// If no error, then no error
	if err == nil {
		return
	}

	// Try to add this error to the database log.
	if d.logDatabase != nil {

		record := Request{
			UserAgent:  request.Header.Get("User-Agent"),
			IPAddress:  request.RemoteAddr,
			Path:       request.URL.Path,
			StatusCode: derp.ErrorCode(err),
		}

		if err := d.logDatabase.Save(&record, ""); err != nil {
			derp.Report(derp.Wrap(err, "dome.HandleError", "Error saving log record"))
		}
	}

	/*/ If we're blocking certain errors, then try to check for that now
	if d.blockOnError {
		switch derp.ErrorCode(err) {
		case http.StatusNotFound, http.StatusForbidden:
			errorCount, _ := d.blockedIPs.Get(request.RemoteAddr)
			ttl := time.Duration(2^errorCount) * time.Minute
			d.blockedIPs.Set(request.RemoteAddr, errorCount+1, ttl)
		}
	}
	*/
}

func (d *Dome) Close() {
	// d.blockedIPs.Close()
}

/*
func createCache(capacity int) (otter.CacheWithVariableTTL[string, int], error) {

	// Don't allow negative cache sizes
	if capacity < 0 {
		capacity = 0
	}

	// Create a new cache with the correct capacity
	return otter.MustBuilder[string, int](capacity).
		WithVariableTTL().
		Build()
}
*/
