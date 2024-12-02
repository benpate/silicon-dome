# Digital Dome

Digital Dome is an opinionated shield against AI and bot scanners.

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
