package data

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Result is a request's result
type Result struct {
	RPS      float64       // RPS is the request per second
	Requests int           // Requests is the number of request
	Errors   int           // Number of errors
	Bytes    int64         // Number of bytes downloaded
	Duration time.Duration // Duration of requests
	Fastest  time.Duration // Fastest request results
	Slowest  time.Duration // Slowest request result
	Status   int           // Status is a request's HTTP status
	Error    error         // Error is not nil
}

func (r *Result) String() string {
	var s strings.Builder
	r.Fprint(&s)
	return s.String()
}

// Merge
func (r *Result) Merge(b *Result) {
	r.Requests++
	r.Bytes += b.Bytes

	if r.Fastest == 0 || b.Duration < r.Fastest {
		r.Fastest = b.Duration
	}
	if b.Duration > r.Slowest {
		r.Slowest = b.Duration
	}

	switch {
	case b.Error != nil:
		fallthrough
	case b.Status >= http.StatusBadRequest:
		r.Errors++
	}
}

// Finalize
func (r *Result) Finalize(td time.Duration) *Result {
	r.Duration = td
	r.RPS = float64(r.Requests) / td.Seconds()
	return r
}

// Fprint
func (r *Result) Fprint(out io.Writer) {
	p := func(format string, args ...any) {
		fmt.Fprintf(out, format, args...)
	}
	p("\nSummary    :\n")
	p("\tSuccess    : %.0f%%\n", r.success())
	p("\tRPS        : %.1f\n", r.RPS)
	p("\tRequests   : %d\n", r.Requests)
	p("\tErrors     : %d\n", r.Errors)
	p("\tBytes      : %d\n", r.Bytes)
	p("\tDuration   : %s\n", round(r.Duration))
	if r.Requests > 1 {
		p("\tFastest    : %s\n", round(r.Fastest))
		p("\tSlowest    : %s\n", round(r.Slowest))
	}
}

func (r *Result) success() float64 {
	rr, e := float64(r.Requests), float64(r.Errors)
	return (rr - e) / rr * 100
}

func round(t time.Duration) time.Duration {
	return t.Round(time.Microsecond)
}
