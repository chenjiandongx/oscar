package modules

import (
	"context"
	"strings"
	"time"

	"github.com/chenjiandongx/oscar/fixtures"
)

type git struct{}

func (mod *git) Name() string {
	return ModGit
}

func (mod *git) Display(highCpu bool) {
	if highCpu {
		ctx, cancel := context.WithCancel(context.Background())
		go BusyCPUWorking(ctx)
		defer func() {
			cancel()
			time.Sleep(100 * time.Microsecond)
		}()
	}

	for i := 0; i < len(fixtures.GitConsole); i++ {
		if strings.HasPrefix(fixtures.GitConsole[i], "#") {
			SleepInMills(500, 600)
			PrintWithDelay(GreenString(fixtures.GitConsole[i]+"\n"), 20)
			continue
		}
		PrintWithDelay(fixtures.GitConsole[i]+"\n", 20)
	}
}

func NewGitModule() Moduler {
	return &git{}
}
