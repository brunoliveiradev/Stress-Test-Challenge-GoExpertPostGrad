package stress

import (
	"Stress-Test-Challenge-GoExpertPostGrad/internal/configs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRun_WithValidConfig_ReturnsReport(t *testing.T) {
	cfg := &configs.Config{
		URL:         "http://google.com",
		Requests:    50,
		Concurrency: 5,
	}
	st := NewStressTest(cfg)

	report, err := st.Run()

	assert.NoError(t, err)
	assert.NotNil(t, report)
}
