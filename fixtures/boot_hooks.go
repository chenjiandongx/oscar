package fixtures

import (
	"bufio"
	"bytes"
)

var bootHooks = `
base
udev
usr
resume
autodetect
modconf
block
net
mdadm
mdadm_udev
keyboard
keymap
consolefont
encrypt
lvm2
fsck
filesystems
`

func readLines(in string) []string {
	var buf bytes.Buffer
	buf.WriteString(in)

	lines := make([]string, 0)
	scanner := bufio.NewScanner(&buf)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// the first line is empty
	return lines[1:]
}

var BootHooks = readLines(bootHooks)
