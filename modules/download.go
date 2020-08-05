package modules

import (
	"time"

	"github.com/schollz/progressbar/v3"
)

type download struct{}

func (mod *download) Name() string {
	return ModDownload
}

func (mod *download) Display(forever bool) {
	mod.display()
	if forever {
		for {
			time.Sleep(time.Second)
			mod.display()
		}
	}
}

func (mod *download) display() {
	count := GenIntN(100, 2000)
	bar := progressbar.Default(int64(count))
	for i := 0; i < count; i++ {
		bar.Add(1)
		SleepInMills(50, 200)
	}
	bar.Clear()
}

func NewDownloadModule() Moduler {
	return &download{}
}
