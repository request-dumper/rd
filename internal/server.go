package internal

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func StartServer(port int) {
	appName := "Request-Dumper"

	app := fiber.New(fiber.Config{
		AppName:               appName,
		Prefork:               false,
		DisableKeepalive:      true,
		DisableStartupMessage: true,
		Network:               fiber.NetworkTCP,
	})

	app.Use(logRequests)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}

func logRequests(c *fiber.Ctx) error {
	url, err := url.Parse(string(c.Request().RequestURI()[:]))

	if err != nil {
		log.Fatal(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	header := parseHeader(c)

	fields := log.Fields{
		"method":          c.Method(),
		"baseUrl":         c.BaseURL(),
		"path":            url.Path,
		"queryParameters": url.Query(),
		"header":          header,
		"clientIp":        c.IP(),
	}

	if c.Body() != nil {
		fields["body"] = string(c.Body()[:])
	}

	log.WithFields(fields).Info(fmt.Sprintf("%s %s%s", c.Method(), c.BaseURL(), c.OriginalURL()))

	return c.SendStatus(fiber.StatusOK)
}

// Parse the header into a map. This function also handels multiple headers with the same name
// and also multiple values which are separated via a colon.
func parseHeader(c *fiber.Ctx) map[string][]string {
	header := make(map[string][]string)

	c.Context().Request.Header.VisitAll(func(key []byte, value []byte) {
		headerName := string(key[:])
		headerValue := string(value[:])

		for _, v := range strings.Split(headerValue, ", ") {
			for _, v2 := range strings.Split(v, ",") {
				header[headerName] = append(header[headerName], v2)
			}
		}
	})

	return header
}
