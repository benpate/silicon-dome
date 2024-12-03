# Digital Dome ✨

<img alt="AI-generated image of flying robotic vehicles exploding in a war" src="https://github.com/benpate/digital-dome/raw/main/meta/banner.webp" style="width:100%; display:block; margin-bottom:20px;">


[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://pkg.go.dev/github.com/benpate/digital-dome)
[![Version](https://img.shields.io/github/v/release/benpate/digital-dome?include_prereleases&style=flat-square&color=brightgreen)](https://github.com/benpate/digital-dome/releases)
[![Build Status](https://img.shields.io/github/actions/workflow/status/benpate/digital-dome/go.yml?style=flat-square)](https://github.com/benpate/digital-dome/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/benpate/digital-dome?style=flat-square)](https://goreportcard.com/report/github.com/benpate/digital-dome)
[![Codecov](https://img.shields.io/codecov/c/github/benpate/digital-dome.svg?style=flat-square)](https://codecov.io/gh/benpate/digital-dome)

## Opinionated protection from AIs and scanner bots.

Digital dome is a fast, minimal web application firewall that uses request information to protect a site against AI scanners.  It hosts several configurable rules, along with sensible (if aggressive) defaults.

## Quickstart

```golang
import github.com/benpate/digital-dome/dome
import github.com/benpate/digital-dome/dome4echo

domeConfig := dome.New()                // Create a new digital dome (using sensible defaults)
middleware := dome4echo.New(domeConfig) // Create echo middleware
e.Pre(middleware)                       // Use the middleware

// easy peasy.
```

## AI Scrapers

I've manually collected this list of AI scrapers that Digital Dome uses to protect your website.

```
Amazonbot
anthropic-ai
AdsBot-Google
Applebot
Applebot-Extended
AwarioRssBot
AwarioSmartBot
Bytespider
CCBot
ChatGPT
ChatGPT-User
Claude
ClaudeBot
Claude-Web
cohere-ai
DataForSeoBot
Diffbot
FacebookBot
FacebookExternalHit
FriendlyCrawler
Google-CloudVertexBot
Google-Extended
GPTBot
ImagesiftBot
magpie-crawler
Meta-ExternalAgent
meta-externalagent
NewsNow
news-please
OAI-SearchBot
omgili
omgilibot
peer39_crawler
PerplexityBot
PetalBot
Quora-Bot
Scrapy
TurnitinBot
Twitterbot
YaK
Yandex
YouBot
```


## Options

| Option | Description |
|--------|-------------|
| **Block User Agents** | |
| `BlockUserAgents(strings...)` | Digital Dome can block requests based on any number of provided `User-Agent` strings.  It uses an efficient [Aho-Corasick](https://github.com/cloudflare/ahocorasick) string matching algorithm from CloudFlare to perform this operation quickly. |
| `BlockKnownAIBots()` | Digital Dome maintains a [list of known AI bots](https://github.com/benpate/digital-dome/blob/main/dome/constant_userAgents.go#L59) that it can compare against each request's `User-Agent` |
| `BlockKnownBadBots()` | (DEFAULT) Digital Dome maintains a [list of known bad actors](https://github.com/benpate/digital-dome/blob/main/dome/constant_userAgents.go#L11) that it can compare against each request's `User-Agent`.  This includes all of the AI bots listed above, plus several hundred more non-search-engine user agents that are used for scraping your website. |
| **Block Paths** | |
| `BlockPaths(strings...)` | Digital Dome can block requests based on any number of provided path names.  As with `User-Agent` blocking, it uses an efficient Aho-Corasick string matching algorithm from CloudFlare to perform this operation quickly. |
| `BlockKnownPaths()` | (DEFAULT) Digital Dome maintains a [list of known paths]() that are frequently scanned by scammers and are blocked by default. |
| **Log Errors** | |
| `LogDatabase(data.Collection)` | Digital Dome can log failed requests to a database if a data.Collection is provided to this Option (see Storage below) |
| `LogStatusCodes(ints...)` | Customize the status codes are logged using this Option.  Default values are: `StatusBadRequest`, `StatusNotFound`, and `StatusForbidden` |
| **Block Malicious Requests** | |
| `BlockStatusCodes(ints...)` | Digital Dome can track when requests trigger specific errors (for example, `StatusForbidden`) and block all requests from that IP address.  See `Blocking` below for details. |
| `BlockCache(capacity)` | This option sets the capacity of the blocked IP address cache.  Default is 1024 IP addresses. |

## Routers

Digital Dome is build to work with any Go HTTP Router library or framework.  There is currently one adapter, made for [labstack echo](https://echo.labstack.com).  Middleware adapters are very easy to make, so if your router is not listed below, please [file an issue](https://github.com/benpate/digital-dome/issues) to get one made.

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

## IP Blocking

Digital Dome tracks errors generated by your application and returned through the middleware.  

If a specific IP address generates too many errors in a designated timespan, future requests from that IP will be blocked before they reach your application.  

The length of time an IP is blocked grows exponentially with the number of bad requests they make.  So, honest mistakes will heal quickly and automatically, and script-kiddie scanners will ban themselves into oblivion.

By default, Digital Dome counts all `StatusForbidden` responses towards the quota for any IP address, and begins blocking all traffic after 5 forbidden requests within one minute.  You can calibrate the kinds of responses that trigger this behavior using the `BlockStatusCodes()` option.

```golang
domeConfig := dome.New(          // Create the digital dome shield
    dome.BlockStatusCodes(404)   // Choose status codes to trigger blocking behavior
```

## Image Used Without Permission

**Image Credits.  None.**  The banner image was an AI-generated using https://www.freepik.com.  AI images cannot be copyrighted or owned by anyone.  Savor the delicious irony.

## Pull Requests Welcome

While many parts of this module have been used for years in production environments, it is still a work in progress and will benefit from your experience reports, use cases, and contributions.  If you have an idea for making Rosetta better, send in a pull request.  We're all in this together! ✨