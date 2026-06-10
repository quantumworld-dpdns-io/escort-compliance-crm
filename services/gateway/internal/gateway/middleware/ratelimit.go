package gateway

import (
	"net/http"
	"sync"
	"time"
)

type rateLimiter struct {
	mu       sync.Mutex
	clients  map[string]*clientRate
	limit    int
	window   time.Duration
}

type clientRate struct {
	count    int
	lastSeen time.Time
}

func newRateLimiter(limit int, window time.Duration) *rateLimiter {
	rl := &rateLimiter{
		clients: make(map[string]*clientRate),
		limit:   limit,
		window:  window,
	}
	go rl.cleanup()
	return rl
}

func (rl *rateLimiter) cleanup() {
	ticker := time.NewTicker(rl.window)
	for range ticker.C {
		rl.mu.Lock()
		for ip, client := range rl.clients {
			if time.Since(client.lastSeen) > rl.window {
				delete(rl.clients, ip)
			}
		}
		rl.mu.Unlock()
	}
}

func (rl *rateLimiter) allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	client, exists := rl.clients[ip]
	if !exists {
		rl.clients[ip] = &clientRate{count: 1, lastSeen: time.Now()}
		return true
	}

	if time.Since(client.lastSeen) > rl.window {
		client.count = 1
		client.lastSeen = time.Now()
		return true
	}

	if client.count >= rl.limit {
		return false
	}

	client.count++
	client.lastSeen = time.Now()
	return true
}

func rateLimitMiddleware() func(http.Handler) http.Handler {
	limiter := newRateLimiter(100, time.Minute)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !limiter.allow(r.RemoteAddr) {
				http.Error(w, `{"error":"rate_limit_exceeded"}`, http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
