package modules

import (
	"context"
	"fmt"
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
		SleepInMills(300, 800)
		fmt.Println(fixtures.GitConsole[i])
	}

}

func NewGitModule() Moduler {
	return &git{}
}
