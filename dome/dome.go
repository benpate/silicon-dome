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
	blockedIPs        otter.CacheWithVariableTTL[string, int]
	logDatabase       data.Collection
	logStatusCodes    []int
	blockStatusCodes  []int
}

// New returns a fully initialized Dome object.
func New(options ...Option) Dome {

	result := Dome{
		blockedIPs: createCache(1024),
	}

	// Default settings...
	result.With(
		BlockKnownBadBots(),
		BlockKnownPaths(),
		BlockStatusCodes(http.StatusForbidden),
		LogStatusCodes(http.StatusBadRequest, http.StatusNotFound),
	)

	// Custom settings...
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

	// If this IP address has caused more than 5 qualifying errors (since the TTL) then block this request.
	if count, _ := dome.blockedIPs.Get(realIPAddress(request)); count > 5 {
		return derp.NewForbiddenError("dome.VerifyHeader", "Blocked due to previous scanning activity.  Try again later.", request.RemoteAddr)
	}

	// Try to block request based on the User-Agent
	userAgent := request.Header.Get("User-Agent")

	if userAgent == "" {
		return derp.NewForbiddenError("dome.VerifyHeader", "User Agent must not be empty")
	}

	if dome.blockedUserAgents != nil {
		if dome.blockedUserAgents.Contains([]byte(userAgent)) {
			return derp.NewForbiddenError("dome.VerifyHeader", "User Agent is blocked", userAgent)
		}
	}

	// Try to block request based on the URL/Path
	if dome.blockedPaths != nil {
		path := request.URL.Path
		if dome.blockedPaths.Contains([]byte(path)) {
			return derp.NewForbiddenError("dome.VerifyHeader", "Path is blocked", path)
		}
	}

	// This request is ALLOWED.
	return nil
}

// HandleError is called by the HTTP middleware to report an error back into the Dome.
// Based on configureation settings, this will log the error and/or block the IP address.
func (d *Dome) HandleError(request *http.Request, err error) {

	// If no error, then no error
	if err == nil {
		return
	}

	statusCode := derp.ErrorCode(err)

	// Try to add this error to the database log.
	if d.logDatabase != nil {

		// If this is a status code that we want to log, then log it.
		if sliceContains(d.logStatusCodes, statusCode) {

			record := Request{
				UserAgent:  request.Header.Get("User-Agent"),
				IPAddress:  realIPAddress(request),
				URL:        request.Host + request.URL.RequestURI(),
				Method:     request.Method,
				StatusCode: statusCode,
				StatusText: http.StatusText(statusCode),
				Error:      err,
			}

			if err := d.logDatabase.Save(&record, ""); err != nil {
				derp.Report(derp.Wrap(err, "dome.HandleError", "Error saving log record"))
			}
		}
	}

	// Try to block this IP address based on the statusCode
	if sliceContains(d.blockStatusCodes, statusCode) {
		remoteAddress := realIPAddress(request)          // get the real IP address (not some shifty, fake one)
		errorCount, _ := d.blockedIPs.Get(remoteAddress) // get the existing error count
		errorCount = errorCount + 1                      // increment
		ttl := getTTL(errorCount)                        // calculate the TTL based on the number of errors in the queue
		d.blockedIPs.Set(remoteAddress, errorCount, ttl) // save the new error count.
	}
}

func (d *Dome) Close() {
	d.blockedIPs.Close()
}
