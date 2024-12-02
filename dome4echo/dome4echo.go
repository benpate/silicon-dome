package dome4echo

import (
	"net/http"

	"github.com/benpate/derp"
	"github.com/benpate/silicon-dome/dome"
	"github.com/labstack/echo/v4"
)

func New(options ...dome.Option) echo.MiddlewareFunc {

	dome := dome.New(options...)

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(ctx echo.Context) error {

			// If this request is blocked, then halt here.
			if err := dome.VerifyRequest(ctx.Request()); err != nil {
				dome.HandleError(ctx.Request(), err)
				ctx.Response().Header().Set("X-Dome-Blocked", derp.Message(err))
				return ctx.String(http.StatusForbidden, "Forbidden")
			}

			// Try to execute the request
			if err := next(ctx); err != nil {
				dome.HandleError(ctx.Request(), err)
				return err
			}

			// Done.
			return nil
		}
	}
}
