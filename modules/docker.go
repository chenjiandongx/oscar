package modules

import (
	"context"
	"fmt"
	"time"

	"github.com/chenjiandongx/oscar/fixtures"
)

type docker struct{}

func (mod *docker) Name() string {
	return ModDocker
}

func (mod *docker) Display(highCpu bool) {
	if highCpu {
		ctx, cancel := context.WithCancel(context.Background())
		go BusyCPUWorking(ctx)
		defer func() {
			cancel()
			time.Sleep(100 * time.Microsecond)
		}()
	}

	for _, idx := range GenRandomIndex(20, 100, len(fixtures.DockerPackages)) {
		pack, version := fixtures.DockerPackages[idx], GenPackageTag()
		fmt.Printf("Untagged: %s:%s\n", pack, version)
		fmt.Printf("Untagged: %s:%s@%s\n", pack, version, GenHashHex(32))

		for i := 0; i < GenIntN(10, 30); i++ {
			fmt.Printf("Deleted: sha256:%s\n", GenHashHex(32))
		}
		SleepInMills(500, 5000)
	}
}

func NewDockerModule() Moduler {
	return &docker{}
}
