package report

import "Stress-Test-Challenge-GoExpertPostGrad/internal/domain"

func GenerateReport(results []domain.RequestResult, duration float64, totalRequests int) *domain.Report {
	report := &domain.Report{
		Failures:           make(map[int]int),
		Duration:           duration,
		ConfiguredRequests: totalRequests,
	}

	for _, result := range results {
		report.TotalRequests++
		if result.StatusCode == 200 {
			report.Success++
		} else {
			report.Failures[result.StatusCode]++
		}
	}

	return report
}
