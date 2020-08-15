package modules

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/chenjiandongx/oscar/fixtures"
)

type kernelCompile struct{}

func (mod *kernelCompile) Name() string {
	return ModKernelCompile
}

func (mod *kernelCompile) Display(highCpu bool) {
	if highCpu {
		ctx, cancel := context.WithCancel(context.Background())
		go BusyCPUWorking(ctx)
		defer func() {
			cancel()
			time.Sleep(100 * time.Microsecond)
		}()
	}

	arch := fixtures.Archs[rand.Intn(len(fixtures.Archs))]
	for i := 0; i < GenIntN(100, 600); i++ {
		line := GenKernelCompileLine(arch)
		fmt.Println(line)
		SleepInMills(20, 1000)
	}

	fmt.Printf("BUILD   arch/%s/boot/bzImage\n\n", arch)

	bs := GenIntN(9000, 1000000)
	fmt.Printf("Setup is %d bytes (padded to %d bytes).\n", bs, GenIntN(bs, 1000000))
	fmt.Printf("System is %d kB\n", GenIntN(300, 3000))
	fmt.Printf("CRC %x\n", GenIntN(0x10000000, 0xffffffff))
	fmt.Printf("Kernel: arch/%s/boot/bzImage is ready (#1)\n", arch)
}

func NewKernelCompileModule() Moduler {
	return &kernelCompile{}
}
