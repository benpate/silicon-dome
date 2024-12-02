package dome

import (
	"github.com/benpate/data/journal"
)

type Request struct {
	UserAgent  string `bson:"userAgent"`
	IPAddress  string `bson:"ipAddress"`
	Path       string `bson:"path"`
	StatusCode int    `bson:"statusCode"`
	StatusText string `bson:"statusText"`
	journal.Journal
}

func (request Request) ID() string {
	return ""
}
