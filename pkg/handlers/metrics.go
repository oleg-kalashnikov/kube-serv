package handlers

import (
	"net/http"
	"time"
)

//CaptureMetrics handler implements middleware logic
func (h *Handler) CaptureMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timer := time.Now()
		next.ServeHTTP(w, r)
		h.countDuration(timer)
		// h.collectCodes(r)

	})
}

func (h *Handler) countDuration(timer time.Time) {
	if !timer.IsZero() {
		h.stats.requestsCount++
		took := time.Now()
		duration := took.Sub(timer)
		h.stats.totalDuration += duration
		if duration > h.stats.maxDuration {
			h.stats.maxDuration = duration
		}
		h.stats.averageDuration = h.stats.totalDuration / h.stats.requestsCount
		h.stats.requests.Duration.Max = h.stats.maxDuration.String()
		h.stats.requests.Duration.Average = h.stats.averageDuration.String()
	}
}

// func (h *Handler) collectCodes(r *http.Request) {
// 	if r.Code >= 500 {
// 		h.stats.requests.Codes.C5xx++
// 	} else {
// 		if c.GetCode() >= 400 {
// 			h.stats.requests.Codes.C4xx++
// 		} else {
// 			if c.GetCode() >= 200 && c.GetCode() < 300 {
// 				h.stats.requests.Codes.C2xx++
// 			}
// 		}
// 	}
// }
// github.com/felixge/httpsnoop
