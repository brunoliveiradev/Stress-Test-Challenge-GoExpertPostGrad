package domain

import "fmt"

type RequestResult struct {
	StatusCode int
}

type Report struct {
	TotalRequests int
	Success       int
	Failures      map[int]int
	Duration      float64
}

func (r *Report) Generate() string {
	result := fmt.Sprintf("Total time duration: %.2f seconds\n", r.Duration)
	result += fmt.Sprintf("Total requests: %d\n", r.TotalRequests)
	result += fmt.Sprintf("Successful requests (HTTP 200): %d\n", r.Success)
	result += "Status code distribution:\n"
	for status, count := range r.Failures {
		result += fmt.Sprintf("  Status %d: %d\n", status, count)
	}
	return result
}
