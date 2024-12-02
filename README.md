# Silicon Dome

Silicon Dome is an opinionated shield against AI and bot scanners.

## Routers

Silicon Dome currently ships with adapters for [labstack echo](https://echo.labstack.com).  Other router configurations
are easy to add

```
import github.com/benpate/silicon-dome/dome
import github.com/benpate/silicon-dome/dome4echo

// WAF Middleware
siliconDome := dome4echo.New()
e.Pre(siliconDome)
```


## other Routers
