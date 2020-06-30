package logging

import (
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

var CorrelationIdHeader = "X-Correlation-Id"
var UserCorrelationCookie = ""

// EnsureCorrelationId returns the correlation from of the request.
// If the request does not have a correlation id, one will be generated and set to the request.
func EnsureCorrelationId(r *http.Request) string {
	id := r.Header.Get(CorrelationIdHeader)
	if id == "" {
		id = randStringBytes(10)
		r.Header.Set(CorrelationIdHeader, id)
	}
	return id
}

// GetCorrelationId returns the correlation from of the request.
func GetCorrelationId(h http.Header) string {
	return h.Get(CorrelationIdHeader)
}

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// GetCorrelationId returns the users correlation id of the headers.
func GetUserCorrelationId(h http.Header) string {
	if UserCorrelationCookie != "" {
		if value, found := ReadCookieValue(h, UserCorrelationCookie); found {
			return value
		}
	}
	return ""
}
func ReadCookieValue(h http.Header, cookieName string) (string, bool) {
	lines, ok := h["Cookie"]
	if !ok {
		return "", false
	}

	for _, line := range lines {
		parts := strings.Split(strings.TrimSpace(line), ";")
		if len(parts) == 1 && parts[0] == "" {
			continue
		}
		for i := 0; i < len(parts); i++ {
			parts[i] = strings.TrimSpace(parts[i])
			if len(parts[i]) == 0 {
				continue
			}
			name, val := parts[i], ""
			if j := strings.Index(name, "="); j >= 0 {
				name, val = name[:j], name[j+1:]
			}
			if cookieName == name {
				if len(val) > 1 && val[0] == '"' && val[len(val)-1] == '"' {
					val = val[1 : len(val)-1]
				}
				return val, true
			}
		}
	}
	return "", false
}
