package modules

type kernelCompile struct{}

func (mod *kernelCompile) Name() string {
	return ModKernelCompile
}

func (mod *kernelCompile) Display() {}

func NewKernelCompileModule() Moduler {
	return &kernelCompile{}
}
