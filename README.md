# Digital Dome ✨

<img alt="AI-generated image of flying robotic vehicles exploding in a war" src="https://github.com/benpate/digital-dome/raw/main/meta/banner.webp" style="width:100%; display:block; margin-bottom:20px;">


[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://pkg.go.dev/github.com/benpate/digital-dome)
[![Version](https://img.shields.io/github/v/release/benpate/digital-dome?include_prereleases&style=flat-square&color=brightgreen)](https://github.com/benpate/digital-dome/releases)
[![Build Status](https://img.shields.io/github/actions/workflow/status/benpate/digital-dome/go.yml?style=flat-square)](https://github.com/benpate/digital-dome/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/benpate/digital-dome?style=flat-square)](https://goreportcard.com/report/github.com/benpate/digital-dome)
[![Codecov](https://img.shields.io/codecov/c/github/benpate/digital-dome.svg?style=flat-square)](https://codecov.io/gh/benpate/digital-dome)

## Opinionated protection from AIs and scanner bots.

Digital dome is a fast, minimal web application firewall that uses request information to protect a site against AI scanners.  It hosts several configurable rules, along with sensible (if aggressive) defaults.

## Options

### Block User Agents

| Option | Description |
|--------|-------------|
| `BlockUserAgents(strings...)` | Digital Dome can block requests based on any number of provided `User-Agent` strings.  It uses an efficient [Aho-Corasick](https://github.com/cloudflare/ahocorasick) string matching algorithm from CloudFlare to perform this operation quickly. |
| `BlockKnownBadBots()` | (DEFAULT) Digital Dome maintains a [list of known bad actors](https://github.com/benpate/digital-dome/blob/main/dome/constant_userAgents.go#L11) that it can compare against each request's `User-Agent` |
| `BlockKnownAIBots()` | Digital Dome maintains a [list of known AI bots](https://github.com/benpate/digital-dome/blob/main/dome/constant_userAgents.go#L59) that it can compare against each request's `User-Agent` |

### Block Paths

| Option | Description |
|--------|-------------|
| `BlockPaths(strings...)` | Digital Dome can block requests based on any number of provided path names.  As with `User-Agent` blocking, it uses an efficient Aho-Corasick string matching algorithm from CloudFlare to perform this operation quickly. |
| `BlockKnownPaths()` | (DEFAULT) Digital Dome maintains a [list of known paths]() that are frequently scanned by scammers and are blocked by default. |

### Log Errors

| Option | Description |
|--------|-------------|
| `LogDatabase(data.Collection)` | Digital Dome can log failed requests to a database if a data.Collection is provided to this Option (see Storage below) |
| `LogStatusCodes(ints...)` | Customize the status codes are logged using this Option.  Default values are: `StatusBadRequest`, `StatusNotFound`, and `StatusForbidden` |

### Block Malicious Requests

`BlockStatusCodes(ints...)`

`BlockCache(capacity)`

## Routers

Digital Dome is build to work with any Go HTTP Router library or framework.  These are very easy to make, so if your router is not listed below, please [file an issue](https://github.com/benpate/digital-dome/issues) to get one made.

### Labstack Echo

Digital Dome currently ships with adapters for the [labstack echo](https://echo.labstack.com) router. 

```golang
import github.com/benpate/digital-dome/dome
import github.com/benpate/digital-dome/dome4echo

// WAF Middleware

domeConfig := dome.New()                // Additional options go here.
middleware := dome4echo.New(domeConfig) // Populate the dome4echo middleware
e.Pre(middleware)                       // Install the middleware
```

## Storage

Digital Dome uses a [storage adapter](github.com/benpate/data) to write log files to a database.  Currently, I have only written an adapter for MongoDB, but adapters can be written for any database, so please [file an issue](https://github.com/benpate/digital-dome/issues) to make one for your chosen database.

To begin writing log files, simply pass a `data.Collection` into the `LogDatabase` option, and Digital Dome will use it whenever it needs to log a request.

```golang
db, err := mongodb.Connect(...)          // Connect to mongodb
collection := mongodb.                   // Wrap mongo with a data.Collection adapter
    NewSession(db).
    Collection("SiliconDome_LogFiles") 

domeConfig := dome.New(                  // Create the digital dome shield
    dome.LogDatabase(collection),        // Use this database collection to log errors
    dome.LogStatusCodes(404)             // Choose status codes to log
)  
```

## Image Used Without Permission

**Image Credits.  None.**  The banner image was an AI-generated using https://www.freepik.com.  AI images cannot be copyrighted or owned by anyone.  Savor the delicious irony.

## Pull Requests Welcome

While many parts of this module have been used for years in production environments, it is still a work in progress and will benefit from your experience reports, use cases, and contributions.  If you have an idea for making Rosetta better, send in a pull request.  We're all in this together! ✨