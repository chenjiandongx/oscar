package modules

import (
	"fmt"
	"time"

	"github.com/chenjiandongx/react/fixtures"
)

type docker struct{}

func (mod *docker) Name() string {
	return ModDocker
}

func (mod *docker) Display(forever bool) {
	mod.display()
	if forever {
		for {
			time.Sleep(time.Second)
			mod.display()
		}
	}
}

func (mod *docker) display() {
	for _, idx := range GenRandomIndex(20, 100, len(fixtures.DockerPackages)) {
		pack, version := fixtures.DockerPackages[idx], GenPackageTag()
		fmt.Printf("Untagged: %s:%s\n", pack, version)
		fmt.Printf("Untagged: %s:%s@%s\n", pack, version, GenHashHex())

		for i := 0; i < GenIntN(10, 30); i++ {
			fmt.Printf("Deleted: sha256:%s\n", GenHashHex())
		}
		SleepInMills(500, 5000)
	}
}

func NewDockerModule() Moduler {
	return &docker{}
}
