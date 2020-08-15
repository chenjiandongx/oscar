package modules

import (
	"context"
	"fmt"
	"time"

	"github.com/buger/goterm"
)

type botnet struct{}

func (mod *botnet) Name() string {
	return ModBotnet
}

func (mod *botnet) Display(highCpu bool) {
	if highCpu {
		ctx, cancel := context.WithCancel(context.Background())
		go BusyCPUWorking(ctx)
		defer func() {
			cancel()
			time.Sleep(100 * time.Microsecond)
		}()
	}

	var clusters []int
	for i := 0; i < GenIntN(8, 16); i++ {
		clusters = append(clusters, GenIntN(100, 200))
	}

	marker := make([]bool, len(clusters))

	sum := 0
	for idx := range clusters {
		sum += clusters[idx]
	}

	for i := 1; i <= sum; i++ {
		fmt.Printf("\rEstablishing connections: %4d/%4d", i, sum)
		SleepInMills(3, 6)
	}
	PrintNewline()
	SleepInMills(300, 400)

	for idx, cluster := range clusters {
		PrintWithDelay(fmt.Sprintf("  Cluster #%02d (%3d nodes)\n", idx, cluster), 20)
	}

	goterm.MoveCursorUp(len(clusters))
	goterm.Flush()

	for {
		goterm.MoveCursorUp(len(clusters))

		for idx, cluster := range clusters {
			status := YellowBoldString("%s", "booting")
			if marker[idx] {
				status = GreenBoldString("%s", "online ")
			} else {
				if GenBool(0.04) {
					marker[idx] = true
					status = GreenBoldString("%s", "online ")
				}
			}

			fmt.Printf("  Cluster #%02d (%3d nodes) [%s]\n", idx, cluster, status)
		}

		all := true
		for _, mark := range marker {
			if !mark {
				all = false
				break
			}
		}

		if all {
			break
		}

		goterm.Flush()
		SleepInMills(100, 200)
	}

	tasks := []string{"Synchronizing clocks...", "Sending login information...", "Sending command..."}

	for _, task := range tasks {
		SleepInMills(300, 350)
		PrintWithDelay(fmt.Sprintf("+ %s", task), 10)
		SleepInMills(600, 650)
		PrintWithDelay("[done]", 10)
		fmt.Println()
	}

	PrintWithDelay(">> Botnet update complete.\n", 10)
}

func NewBotnetModule() Moduler {
	return &botnet{}
}
