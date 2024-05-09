package main

import (
	"net/http"
	"sync"
	"time"
)

var (
	maxRequestsPerSecond = 10
	ipRequestCount       = make(map[string]int)
	ipRequestCountMutex  sync.Mutex
)

func protecaoDDoS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := r.RemoteAddr

		ipRequestCountMutex.Lock()
		defer ipRequestCountMutex.Unlock()
		count, exists := ipRequestCount[clientIP]

		if !exists {
			ipRequestCount[clientIP] = 1
		} else {
			if count >= maxRequestsPerSecond {
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}
			ipRequestCount[clientIP]++
		}

		go func() {
			<-time.After(time.Second)
			ipRequestCountMutex.Lock()
			defer ipRequestCountMutex.Unlock()
			delete(ipRequestCount, clientIP)
		}()

		next.ServeHTTP(w, r)
	})
}
