package logging

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
	"net/url"
	"sync"
	"time"
)

// ZerologFiberMiddleware uses zerolog as a logging middleware for fiber requests
func ZerologFiberMiddleware() fiber.Handler {
	var (
		errHandler fiber.ErrorHandler
		once       sync.Once
	)

	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Set error handler once
		once.Do(func() {
			errHandler = c.App().ErrorHandler
		})

		// Handle request, store err for logging
		chainErr := c.Next()
		if chainErr != nil {
			// Sent status code 500 when handling of an error fails
			if err := errHandler(c, chainErr); err != nil {
				_ = c.SendStatus(fiber.StatusInternalServerError)
			}
		}

		msg := "Request"
		code := c.Response().StatusCode()
		parsedUrl, err := url.Parse(string(c.Request().RequestURI()[:]))
		if err != nil {
			log.Fatal().
				Err(err).
				Str("request-uri", string(c.Request().RequestURI())).
				Msg("Request uri could not be parsed")
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		headersDict := AssociativeStringArray{value: parseHeader(&c.Context().Request.Header)}
		queryParametersDict := AssociativeStringArray{value: parsedUrl.Query()}

		dumplogger := log.With().
			Int("status", code).
			Str("method", c.Method()).
			Str("path", c.Path()).
			Str("baseUrl", c.BaseURL()).
			Object("queryParameters", &queryParametersDict).
			Object("headers", &headersDict).
			Str("clientIp", c.IP()).
			Str("latency", time.Since(start).String())

		if c.Body() != nil {
			dumplogger.Str("body", string(c.Body()[:]))
		}

		logger := dumplogger.Logger()

		switch {
		case code >= fiber.StatusBadRequest && code < fiber.StatusInternalServerError:
			logger.Warn().Msg(msg)
		case code >= fiber.StatusInternalServerError:
			logger.Error().Err(chainErr).Msg(msg)
		default:
			logger.Info().Msg(msg)
		}

		return nil
	}
}

// Parse the header into a map. This function also handles multiple headers with the same name
// and also multiple values which are separated via a colon.
func parseHeader(requestHeader *fasthttp.RequestHeader) map[string][]string {
	headers := make(map[string][]string)

	requestHeader.VisitAll(func(key []byte, value []byte) {
		headerName := string(key[:])
		headerValue := string(value[:])

		headers[headerName] = append(headers[headerName], headerValue)
	})

	return headers
}
