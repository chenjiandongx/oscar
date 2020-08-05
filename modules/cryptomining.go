package modules

import "time"

type cryptomining struct{}

func (mod *cryptomining) Name() string {
	return ModCryptomining
}

func (mod *cryptomining) Display(forever bool) {
	mod.display()
	if forever {
		for {
			time.Sleep(time.Second)
			mod.display()
		}
	}
}

func (mod *cryptomining) display() {}

func NewCryptominingModule() Moduler {
	return &cryptomining{}
}
