// Package tollbooth_iris provides rate-limiting logic to Iris version 6 request handler.
package tollbooth_iris

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"gopkg.in/kataras/iris.v6"
)

// LimitHandler is a middleware that performs
// rate-limiting given a "limiter" configuration.
func LimitHandler(lmt *limiter.Limiter) iris.HandlerFunc {
	return func(ctx *iris.Context) {
		httpError := tollbooth.LimitByRequest(lmt, ctx.Response, ctx.Request)
		if httpError != nil {
			ctx.SetStatusCode(httpError.StatusCode)
			ctx.WriteString(httpError.Message)
			ctx.StopExecution()
			return
		}

		ctx.Next()
	}
}
