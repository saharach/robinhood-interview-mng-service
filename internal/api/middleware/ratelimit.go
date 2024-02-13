package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// IPRateLimiter represents the rate limiter for each IP address.
type IPRateLimiter struct {
	// mu is used to synchronize access to the requestCounts map.
	mu sync.Mutex
	// requestCounts stores the number of requests made by each IP address.
	requestCounts map[string]int
}

// NewIPRateLimiter creates a new instance of IPRateLimiter.
func NewIPRateLimiter() *IPRateLimiter {
	return &IPRateLimiter{
		requestCounts: make(map[string]int),
	}
}

// RateLimitMiddleware returns a Gin middleware that performs rate limiting per IP address.
func (limiter *IPRateLimiter) RateLimitMiddleware(limit int, interval time.Duration) gin.HandlerFunc {
	// Reset request counts at the specified interval.
	go func() {
		for {
			time.Sleep(interval)
			limiter.mu.Lock()
			limiter.requestCounts = make(map[string]int)
			limiter.mu.Unlock()
		}
	}()

	// Middleware function.
	return func(c *gin.Context) {
		ip := c.ClientIP()

		limiter.mu.Lock()
		defer limiter.mu.Unlock()

		// Increment request count for the IP address.
		limiter.requestCounts[ip]++

		// Check if the request count exceeds the limit.
		if limiter.requestCounts[ip] > limit {
			// Return a 429 Too Many Requests error.
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests. Please try again later.",
			})
			c.Abort()
			return
		}

		// Proceed to the next middleware.
		c.Next()
	}
}
