package middleware

import (
	"net"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
)

func Logger() echo.MiddlewareFunc {
	return LoggerWithConfig()
}

func LoggerWithConfig() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			res := c.Response()
			start := time.Now()
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()

			ra := req.RemoteAddress()
			if ip := req.Header().Get(echo.HeaderXRealIP); ip != "" {
				ra = ip
			} else if ip = req.Header().Get(echo.HeaderXForwardedFor); ip != "" {
				ra = ip
			} else {
				ra, _, _ = net.SplitHostPort(ra)
			}

			status := res.Status()

			rx_bytes := req.Header().Get(echo.HeaderContentLength)

			if rx_bytes == "" {
				rx_bytes = "0"
			}

			tx_bytes := strconv.FormatInt(res.Size(), 10)

			l := stop.Sub(start).Nanoseconds() / 1000
			latency := strconv.FormatInt(l, 10)
			humanLatency := stop.Sub(start).String()

			fields := log.Fields{
				"remote_ip":    ra,
				"method":       req.Method(),
				"uri":          req.URI(),
				"status":       status,
				"rx_bytes":     rx_bytes,
				"tx_bytes":     tx_bytes,
				"latency":      latency,
				"humanLatency": human_latency,
			}

			//		log.SetFormatter(&log.JSONFormatter{})

			switch {
			case status >= 500:
				log.WithFields(fields).Error("Complete hanlding request with error")
			case status >= 300:
				log.WithFields(fields).Warn("Complete handling request with warning")
			default:
				log.WithFields(fields).Info("Complete handling request")
			}
			return
		}
	}
}
