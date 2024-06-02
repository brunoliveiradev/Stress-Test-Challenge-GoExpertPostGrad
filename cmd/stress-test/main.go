package main

import (
	"Stress-Test-Challenge-GoExpertPostGrad/internal/configs"
	"Stress-Test-Challenge-GoExpertPostGrad/internal/stress"
	"flag"
	"fmt"
)

func main() {
	cfg := configs.LoadConfig()

	if cfg.URL == "" {
		fmt.Println("The URL flag is required.")
		flag.Usage()
		return
	}

	stressTest := stress.NewStressTest(cfg)
	report, err := stressTest.Run()
	if err != nil {
		fmt.Println("Error to execute the stress test:", err)
		return
	}

	report.PrintResults()
}
