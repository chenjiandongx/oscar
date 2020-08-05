package modules

import "time"

type kernelCompile struct{}

func (mod *kernelCompile) Name() string {
	return ModKernelCompile
}

func (mod *kernelCompile) Display(forever bool) {
	mod.display()
	if forever {
		for {
			time.Sleep(time.Second)
			mod.display()
		}
	}
}

func (mod *kernelCompile) display() {}

func NewKernelCompileModule() Moduler {
	return &kernelCompile{}
}
