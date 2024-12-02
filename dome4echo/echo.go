package dome4echo

import (
	"net/http"

	"github.com/benpate/silicon-dome/dome"
	"github.com/labstack/echo/v4"
)

func New(options ...dome.Option) echo.MiddlewareFunc {

	dome := dome.New(options...)

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(ctx echo.Context) error {

			// If this request is blocked, then halt here.
			if err := dome.VerifyHeader(ctx.Request().Header); err != nil {
				ctx.Response().Header().Set("X-Dome-Blocked", err.Error())
				return ctx.String(http.StatusForbidden, "Forbidden")
			}

			// Try to execute the request
			if err := next(ctx); err != nil {
				return err
			}

			// If we received a 404, then add IP to the exponential block list
			// if ctx.Response().Status == http.StatusNotFound {
			// Add IP to block list/cache
			// }

			// Done.
			return nil
		}
	}
}
