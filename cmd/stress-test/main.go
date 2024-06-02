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
		fmt.Println("A URL do serviço é obrigatória")
		flag.Usage()
		return
	}

	stressTest := stress.NewStressTest(cfg)
	report, err := stressTest.Run()
	if err != nil {
		fmt.Println("Erro ao executar stress test:", err)
		return
	}

	fmt.Println(report.Generate())
}
