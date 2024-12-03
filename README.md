# Digital Dome

<img alt="AI-generated image of flying robotic vehicles exploding in a war" src="https://github.com/benpate/digital-dome/raw/main/meta/banner.webp" style="width:100%; display:block; margin-bottom:20px;">


[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://pkg.go.dev/github.com/benpate/digital-dome)
[![Version](https://img.shields.io/github/v/release/benpate/digital-dome?include_prereleases&style=flat-square&color=brightgreen)](https://github.com/benpate/digital-dome/releases)
[![Build Status](https://img.shields.io/github/actions/workflow/status/benpate/digital-dome/go.yml?style=flat-square)](https://github.com/benpate/digital-dome/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/benpate/digital-dome?style=flat-square)](https://goreportcard.com/report/github.com/benpate/digital-dome)
[![Codecov](https://img.shields.io/codecov/c/github/benpate/digital-dome.svg?style=flat-square)](https://codecov.io/gh/benpate/digital-dome)

## Opinionated protection from AIs and scanner bots.

Digital dome is a fast, minimal web application firewall that uses request information to protect a site against AI scanners.  It hosts several configurable rules, along with sensible (if aggressive) defaults.

## Routers: Echo

Silicon Dome currently ships with adapters for [labstack echo](https://echo.labstack.com).  Other router configurations
are easy to add

```
import github.com/benpate/digital-dome/dome
import github.com/benpate/digital-dome/dome4echo

// WAF Middleware
myDome := dome4echo.New()
e.Pre(myDome)
```

## Storage MongoDB
...

## Image Used Without Permission

**Image Credits.  None.**  The banner image is an AI-generated image that cannot be copyrighted or owned by anyone.  Savor the delicious irony.

## Pull Requests Welcome

While many parts of this module have been used for years in production environments, it is still a work in progress and will benefit from your experience reports, use cases, and contributions.  If you have an idea for making Rosetta better, send in a pull request.  We're all in this together! 🌹




