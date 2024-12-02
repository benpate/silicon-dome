package dome

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserAgents(t *testing.T) {

	dome := New(BlockKnownBadBots())

	verify := func(userAgent string, allowed bool) {
		header := http.Header{
			"User-Agent": []string{userAgent},
		}

		if allowed {
			require.Nil(t, dome.VerifyHeader(header))
		} else {
			require.NotNil(t, dome.VerifyHeader(header))
		}
	}

	verify("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0.1 Safari/605.1.15", true)
	verify("Mozilla 5.0 / Whatever", true)
	verify("Applebot-Extended", false)
	verify("ClaudeBot", false)
}