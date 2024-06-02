package stress

import (
	"Stress-Test-Challenge-GoExpertPostGrad/internal/configs"
	"Stress-Test-Challenge-GoExpertPostGrad/internal/domain"
	"Stress-Test-Challenge-GoExpertPostGrad/internal/http"
	"Stress-Test-Challenge-GoExpertPostGrad/internal/report"
	"fmt"
	"sync"
	"time"
)

type StressTest struct {
	cfg    *configs.Config
	client http.Client
}

func NewStressTest(cfg *configs.Config) *StressTest {
	return &StressTest{
		cfg:    cfg,
		client: http.NewHttpClient(),
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
				fmt.Println("Error in the request:", err)
				statusCode = 0
			}
			results[index] = domain.RequestResult{StatusCode: statusCode}
			<-ch
		}(i)
	}

	wg.Wait()
	duration := time.Since(startTime).Seconds()

	return report.GenerateReport(results, duration), nil
}
