package modules

import (
	"context"
	"time"

	"github.com/schollz/progressbar/v3"
)

type download struct{}

func (mod *download) Name() string {
	return ModDownload
}

func (mod *download) Display(highCpu bool) {
	if highCpu {
		ctx, cancel := context.WithCancel(context.Background())
		go BusyCPUWorking(ctx)
		defer func() {
			cancel()
			time.Sleep(100 * time.Microsecond)
		}()
	}

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
