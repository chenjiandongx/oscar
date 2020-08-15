package modules

import (
	"context"
	"fmt"
	"time"

	"github.com/chenjiandongx/oscar/fixtures"
)

type gomod struct{}

func (mod *gomod) Name() string {
	return ModGomod
}

func (mod *gomod) Display(highCpu bool) {
	if highCpu {
		ctx, cancel := context.WithCancel(context.Background())
		go BusyCPUWorking(ctx)
		defer func() {
			cancel()
			time.Sleep(100 * time.Microsecond)
		}()
	}

	stages := []string{"finding", "downloading", "extracting"}
	output := make([]string, 0)

	for _, idx := range GenRandomIndex(50, 200, len(fixtures.Gomod)) {
		output = append(output, fixtures.Gomod[idx])
	}

	for _, stage := range stages {
		for _, o := range output {
			fmt.Printf("go: %s %s\n", stage, o)
			SleepInMills(200, 600)
		}
	}
}

func NewGomodModule() Moduler {
	return &gomod{}
}
