package modules

import (
	"fmt"
	"time"

	"github.com/chenjiandongx/oscar/fixtures"
)

type gomod struct{}

func (mod *gomod) Name() string {
	return ModGomod
}

func (mod *gomod) Display(forever bool) {
	mod.display()
	if forever {
		for {
			time.Sleep(time.Second)
			mod.display()
		}
	}
}

func (mod *gomod) display() {
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
