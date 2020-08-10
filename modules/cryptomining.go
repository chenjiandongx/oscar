package modules

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

type cryptomining struct{}

func (mod *cryptomining) Name() string {
	return ModCryptomining
}

func (mod *cryptomining) Display() {
	newJobsEveryNLines := GenIntN(20, 50)
	remainUntilNewJobs := newJobsEveryNLines

	approximateMhsPerSec := GenIntN(1, 99)
	gpuNum := GenIntN(1, 8)

	solutionFoundEveryNLine := GenIntN(80, 200)
	remainNextSolution := solutionFoundEveryNLine

	solutionFoundNum := 0
	now := time.Now()
	for i := 0; i < GenIntN(300, 1000); i++ {
		t := MagentaString(time.Now().Format("15:04:05"))
		switch {
		case remainUntilNewJobs == 0:
			remainUntilNewJobs = newJobsEveryNLines

			info := CyanBoldString("ℹ")
			fmt.Printf("%-3s  %s%s%-20s Received new job #%s  seed: #%s  target: #%s\n",
				info, t, BlackString("|"), BlueString("stratum"), GenHashHex(8), GenHashHex(32), GenHashHex(24),
			)

		case remainNextSolution == 0:
			remainNextSolution = solutionFoundEveryNLine
			solutionFoundNum++

			info := CyanBoldString("ℹ")
			fmt.Printf("%-3s  %s%s%-20s Solution found; Submitted to stratum.buttcoin.org\n",
				info, t, BlueString("|"), BlueString("CUDA0"),
			)
			fmt.Printf("%-3s  %s%s%-20s Nonce: 0x%s\n",
				info, t, BlueString("|"), BlueString("CUDA0"), GenHashHex(16),
			)
			fmt.Printf("%-3s  %s%s%-20s %s\n",
				info, t, BlueString("|"), BlueString("stratum"), GreenString("Accepted."),
			)

		default:
			remainUntilNewJobs--
			remainNextSolution--

			info := CyanBoldString("m")
			var totalMhs float64
			var buf bytes.Buffer
			for j := 0; j < gpuNum; j++ {
				actualMhsPerGpu := float64(approximateMhsPerSec) + rand.Float64()
				buf.WriteString(fmt.Sprintf("gpu/%d %s ", j, CyanString(fmt.Sprintf("%.2f", actualMhsPerGpu))))
				totalMhs += actualMhsPerGpu
			}

			speed := fmt.Sprintf("Speed %s Mhs", CyanBoldString(fmt.Sprintf("%-6.2f", totalMhs)))
			duration := time.Since(now)

			hour, minu := int(duration.Hours()), int(duration.Minutes())
			fmt.Printf("%-3s  %s%s%-18s %s    %s  [A%d+0:R0+0:F0] Time: %02d:%02d\n",
				info, t, BlueString("|"), BlueString("cryptominer"), speed, buf.String(), solutionFoundNum, hour, minu,
			)
		}

		time.Sleep(300 * time.Millisecond)
	}
}

func NewCryptominingModule() Moduler {
	return &cryptomining{}
}
