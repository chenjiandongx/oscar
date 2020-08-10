package modules

import (
	"fmt"
	"time"

	"github.com/chenjiandongx/oscar/fixtures"
)

type cargo struct{}

func (mod *cargo) Name() string {
	return ModCargo
}

func (mod *cargo) Display() {
	stages := []string{"Downloading", "Compiling"}
	output := make([]string, 0)

	for _, idx := range GenRandomIndex(10, 100, len(fixtures.Packages)) {
		output = append(output, fmt.Sprintf("%s v%s", fixtures.Packages[idx], GenPackageVersion()))
	}

	start := time.Now()
	for _, stage := range stages {
		for _, o := range output {
			PrintGreenWithBold(fmt.Sprintf("%12s ", stage), o)
			SleepInMills(200, 2000)
		}
	}

	PrintGreenWithBold(
		fmt.Sprintf("%12s ", "Finished"),
		fmt.Sprintf("release [optimized] target(s) in %.2f secs\n", float32(time.Since(start).Milliseconds())/1000.0),
	)
	PrintNewline()
}

func NewCargoModule() Moduler {
	return &cargo{}
}
