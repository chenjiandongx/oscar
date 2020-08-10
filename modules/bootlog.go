package modules

import (
	"math/rand"
	"strings"
	"time"

	"github.com/chenjiandongx/oscar/fixtures"
)

type bootlog struct{}

func (mod *bootlog) Name() string {
	return ModBootlog
}

func (mod *bootlog) Display() {
	lines := GenIntN(100, 300)
	burstMode := false
	burstLineCount := 0

	for i := 0; i < lines; i++ {
		choice := fixtures.BootLog[rand.Intn(len(fixtures.BootLog))]
		lineSleep := GenIntN(10, 1000)
		charSleep := 5
		burstLine := GenIntN(10, 50)

		if burstMode && burstLineCount < burstLine {
			lineSleep = 30
			charSleep = 0
		} else if burstLineCount == burstLine {
			burstMode = false
			burstLineCount = 0
		} else if !burstMode {
			burstMode = GenBool(0.05)
		}

		wrong := GenBool(0.01)
		if wrong {
			PrintlnWithDelay(RedString("ERROR: %s", choice), 10)
		} else {
			var hasBoldWord = GenBool(0.1)
			if hasBoldWord {
				words := strings.Split(choice, " ")
				PrintWithDelay(BoldString("%s", words[0]), charSleep)
				PrintlnWithDelay(choice[len(words[0]):], charSleep)
			} else {
				PrintlnWithDelay(choice, charSleep)
			}
		}

		if burstMode {
			burstLineCount += 1
		}

		time.Sleep(time.Duration(lineSleep) * time.Millisecond)
	}
	PrintNewline()
}

func NewBootlogModule() Moduler {
	return &bootlog{}
}
