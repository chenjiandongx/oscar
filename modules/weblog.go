package modules

import (
	"fmt"
	"time"

	"github.com/Pallinder/go-randomdata"

	"github.com/chenjiandongx/oscar/fixtures"
)

type weblog struct{}

func (mod *weblog) Name() string {
	return ModWeblog
}

func (mod *weblog) Display(forever bool) {
	mod.display()
	if forever {
		for {
			time.Sleep(time.Second)
			mod.display()
		}
	}
}

func (mod *weblog) display() {
	method := "GET"
	burstMode := false
	burstLineCnt := 0

	for i := 0; i < GenIntN(200, 500); i++ {
		ip := randomdata.IpV4Address()
		code := GenHTTPCode()
		userAgent := randomdata.UserAgentString()
		now := time.Now().Format("2006-01-02T15:04:05.000")
		path := GenHTTPPath(fixtures.Packages, fixtures.Extensions)

		lineSleep := GenIntN(10, 1000)
		burstLine := GenIntN(10, 50)

		if burstMode && burstLineCnt < burstLine {
			lineSleep = 30
		} else if burstLineCnt == burstLine {
			burstMode = false
			burstLineCnt = 0
		} else if !burstMode {
			burstMode = GenBool(0.05)
		}

		if burstMode {
			burstLineCnt++
		}

		fmt.Printf(`%s - [%s] "%s %s HTTP/1.0" %d "%s"%s`, ip, now, method, path, code, userAgent, "\n")
		time.Sleep(time.Duration(lineSleep) * time.Millisecond)
	}
}

func NewWeblogModule() Moduler {
	return &weblog{}
}
