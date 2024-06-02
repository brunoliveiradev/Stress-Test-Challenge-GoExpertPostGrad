package stress

import (
	"Stress-Test-Challenge-GoExpertPostGrad/internal/configs"
	"Stress-Test-Challenge-GoExpertPostGrad/internal/domain"
	customHttp "Stress-Test-Challenge-GoExpertPostGrad/internal/http"
	"Stress-Test-Challenge-GoExpertPostGrad/internal/report"
	"errors"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

type StressTest struct {
	cfg    *configs.Config
	client customHttp.Client
}

func NewStressTest(cfg *configs.Config) *StressTest {
	return &StressTest{
		cfg:    cfg,
		client: customHttp.NewHttpClient(),
	}
}

func (s *StressTest) Run() (*domain.Report, error) {
	var wg sync.WaitGroup
	ch := make(chan struct{}, s.cfg.Concurrency)
	results := make([]domain.RequestResult, s.cfg.Requests)

	startTime := time.Now()

	for i := 0; i < s.cfg.Requests; i++ {
		wg.Add(1)
		ch <- struct{}{}

		go func(index int) {
			defer wg.Done()
			statusCode, err := s.client.DoRequest(s.cfg.URL)
			if err != nil {
				statusCode = handleRequestError(err)
			}
			results[index] = domain.RequestResult{StatusCode: statusCode}
			<-ch
		}(i)
	}

	wg.Wait()
	duration := time.Since(startTime).Seconds()

	return report.GenerateReport(results, duration, s.cfg.Requests), nil
}

func handleRequestError(err error) int {
	if strings.Contains(err.Error(), "stopped after 10 redirects") {
		// the Go http.Client has a default redirect limit of 10, so we return 429 to indicate too many redirects
		return http.StatusTooManyRequests
	}

	var urlErr *url.Error
	if errors.As(err, &urlErr) {
		// A *url.Error is returned for errors like malformed URLs.
		return http.StatusBadRequest
	}

	var netErr net.Error
	if errors.As(err, &netErr) {
		// A net.Error is returned for network-related errors.
		if netErr.Timeout() {
			return http.StatusRequestTimeout
		}
		return http.StatusInternalServerError
	}

	return http.StatusInternalServerError
}
