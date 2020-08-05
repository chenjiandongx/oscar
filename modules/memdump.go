package modules

import (
	"bytes"
	"fmt"
	"time"
)

type memdump struct{}

func (mod *memdump) Name() string {
	return ModMemdump
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
	currentLoc := GenIntN(2<<60, 2<<61)

	var rows []string
	for i := 0; i < GenIntN(600, 1000); i++ {
		var buf bytes.Buffer

		buf.WriteString(fmt.Sprintf("%x ", currentLoc))
		currentLoc += 0x10

		var ls []string
		for j := 0; j < 16; j++ {
			ls = append(ls, fmt.Sprintf("%s", GenHashHexWithLen(2)))
		}

		for idx, l := range ls {
			if idx == 8 {
				buf.WriteString(" ")
			}
			buf.WriteString(fmt.Sprintf("%s ", l))
		}

		buf.WriteString("[")
		for j := 0; j < 16; j++ {
			if GenIntN(0, 2) == 0 {
				buf.WriteString(string(rune(GenIntN(33, 126))))
				continue
			}
			buf.WriteString(".")
		}
		buf.WriteString("]")

		rows = append(rows, buf.String())
	}

	for _, row := range rows {
		fmt.Println(row)
		SleepInMills(60, 200)
	}
}

func NewMemdumpModule() Moduler {
	return &memdump{}
}
