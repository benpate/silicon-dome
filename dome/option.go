package dome

import (
	"github.com/benpate/data"
	"github.com/benpate/derp"
	"github.com/cloudflare/ahocorasick"
	"github.com/maypok86/otter"
)

// Option is a functional argument that configures a Dome object.
type Option func(*Dome)

/******************************************
 * Blocking Known User Agents
 ******************************************/

// BlockKnownAIBots is a dome.Option that blocks known AI crawlers.
func BlockKnownAIBots() Option {
	return BlockUserAgents(KnownAIBots...)
}

// BlockAllBadBots is a dome.Option that blocks all known bad bots.
func BlockKnownBadBots() Option {
	return BlockUserAgents(AllKnownBadBots...)
}

// BlockUserAgents is a dome.Option that blocks the provided user agents.
func BlockUserAgents(blockedAgents ...string) Option {
	return func(d *Dome) {
		d.blockedUserAgents = ahocorasick.NewStringMatcher(blockedAgents)
	}
}

/******************************************
 * Blocking Known Paths
 ******************************************/

func BlockKnownPaths() Option {
	return BlockPaths(KnownPaths...)
}

// BlockPaths is a dome.Option that blocks the provided paths.
func BlockPaths(blockedPaths ...string) Option {
	return func(d *Dome) {
		d.blockedPaths = ahocorasick.NewStringMatcher(blockedPaths)
	}
}

/******************************************
 * 404/Not Found Handling
 ******************************************/

// BlockNotFound configures Dome to block IP addresses that generate 404/Not Found errors
func BlockNotFound(capacity int) Option {
	return func(d *Dome) {
		d.blockNotFound = true
	}
}

// LogIPAddresses is a dome.Option that configures the collection where failed requests will be logged
func LogIPAddresses(collection data.Collection) Option {
	return func(d *Dome) {
		d.logDatabase = collection
	}
}

// CacheCapacity is a dome.Option that initializes a new IPAddress cache and sets its size
func CacheCapacity(capacity int) Option {
	return func(d *Dome) {

		// Don't allow negative cache sizes
		if capacity < 0 {
			capacity = 0
		}

		// If the capacity has not changed, then do nothing.
		if capacity == d.blockedIPs.Capacity() {
			return
		}

		// Close the previous cache, if it exists
		d.blockedIPs.Close()

		// Create a new cache with the correct capacity
		if cache, err := otter.MustBuilder[string, IPAddress](capacity).WithVariableTTL().Build(); err == nil {
			d.blockedIPs = cache
		} else {
			derp.Report(err)
		}
	}
}
