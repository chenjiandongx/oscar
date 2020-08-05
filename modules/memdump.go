package modules

import (
	"time"
)

type memdump struct{}

func (mod *memdump) Name() string {
	return ModGomod
}

func (mod *memdump) Display(forever bool) {
	mod.display()
	if forever {
		for {
			time.Sleep(time.Second)
			mod.display()
		}
	}
}

func (mod *memdump) display() {
	for i := 0; i < GenIntN(50, 200); i++ {

	}
}

func NewMemdumpModule() Moduler {
	return &memdump{}
}
