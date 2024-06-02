package domain

import (
	"fmt"
)

type RequestResult struct {
	StatusCode int
}

type Report struct {
	TotalRequests      int
	Success            int
	Failures           map[int]int
	Duration           float64
	ConfiguredRequests int
}

func (r *Report) PrintResults() {
	requestsCompleted := r.getTotalRequests()
	fmt.Println()
	fmt.Printf("Total time: \t\t\t\t%s\n", r.getDuration())
	fmt.Printf("Configured requests: \t\t\t%d\n", r.ConfiguredRequests)
	fmt.Printf("Completed requests: \t\t\t%d\n", requestsCompleted)
	fmt.Printf("Successful requests (HTTP 200): \t%d\n", r.Success)
	fmt.Printf("\n HTTP STATUS\tTOTAL\n")
	for key, value := range r.Failures {
		fmt.Printf("%d\t\t%d\n", key, value)
	}
	fmt.Println()
}

func (r *Report) getTotalRequests() int {
	t := 0
	for _, v := range r.Failures {
		t += v
	}
	return r.Success + t
}

func (r *Report) getDuration() string {
	duration := r.Duration
	if duration < 60 {
		return fmt.Sprintf("%.2f seconds", duration)
	} else if duration < 3600 {
		return fmt.Sprintf("%.2f minutes", duration/60)
	} else {
		return fmt.Sprintf("%.2f hours", duration/3600)
	}
}
