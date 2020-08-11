package modules

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"math/rand"
	"path"
	"strings"
	"time"

	"github.com/chenjiandongx/oscar/fixtures"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func SleepInMills(min, max int) {
	time.Sleep(time.Duration(GenIntN(min, max)) * time.Millisecond)
}

func GenPackageVersion() string {
	return fmt.Sprintf("%d.%d.%d", rand.Intn(4), rand.Intn(12), rand.Intn(10))
}

func GenRandStringFromList(lst []string, cnt int) string {
	var buf bytes.Buffer
	for _, idx := range GenRandomIndex(0, len(lst), cnt) {
		buf.WriteString(lst[idx])
	}
	return buf.String()
}

func GenHTTPCode() int {
	codes := []int{200, 201, 400, 401, 403, 404, 500, 502, 503}
	if (time.Now().Unix() % 2) == 0 {
		return 200
	}
	return codes[rand.Intn(len(codes))]
}

func GenHTTPPath(files, exts []string) string {
	file := files[rand.Intn(len(files))]
	ext := exts[rand.Intn(len(exts))]

	return fmt.Sprintf("/files/%s.%s", file, ext)
}

func GenLinkerFlags(lst []string, cnt int) string {
	var buf bytes.Buffer
	for _, idx := range GenRandomIndex(0, len(lst), cnt) {
		buf.WriteString("-l")
		buf.WriteString(lst[idx])
		buf.WriteString(" ")
	}
	return buf.String()
}

func GenIncludes(lst []string, cnt int) string {
	filter := make(map[string]struct{})
	num := 0
	for _, l := range lst {
		p := path.Base(l)
		if _, ok := filter[p]; ok {
			continue
		}
		filter[p] = struct{}{}
		num++
		if num >= cnt {
			break
		}
	}

	var buf bytes.Buffer
	for k := range filter {
		buf.WriteString("-I")
		buf.WriteString(k)
		buf.WriteString(" ")
	}

	return buf.String()
}

func GenPackageTag() string {
	return fixtures.DockerTags[rand.Intn(len(fixtures.DockerTags))]
}

func GenHashHex(n int) string {
	key := [64]byte{}
	rand.Read(key[:])
	return fmt.Sprintf("%x", md5.Sum(key[:]))[:n]
}

func GenIntN(min, max int) int {
	return min + rand.Intn(max-min)
}

func GenRandomIndex(min, max, length int) []int {
	ret := make([]int, 0)
	filter := map[int]struct{}{}

	for i := 0; i < GenIntN(min, max); i++ {
		for {
			idx := rand.Intn(length)
			if _, ok := filter[idx]; !ok {
				ret = append(ret, idx)
				break
			}
		}

	}

	return ret
}

func GenBool(p float64) bool {
	return float64(time.Now().UnixNano()/100%100)/100.0 < p
}

func PrintWithDelay(s string, d int) {
	for _, r := range []rune(s) {
		fmt.Print(string(r))
		time.Sleep(time.Duration(d) * time.Millisecond)
	}
}

func PrintlnWithDelay(s string, d int) {
	for _, r := range []rune(s) {
		fmt.Print(string(r))
		time.Sleep(time.Duration(d) * time.Millisecond)
	}
	fmt.Println()
}

func PrintNewline() {
	fmt.Println()
}

func genHeader(arch string) string {
	rareCmds := []string{"SYSTBL ", "SYSHDR "}
	cmds := []string{"WRAP   ", "CHK    ", "UPD    "}

	cmd := cmds[rand.Intn(len(cmds))]
	if GenBool(1.0 / 15.0) {
		cmd = rareCmds[rand.Intn(len(rareCmds))]
	}

	cfiles := fmt.Sprintf("%sh", fixtures.Cfiles[rand.Intn(len(fixtures.Cfiles))])
	if strings.HasPrefix(cfiles, "arch") {
		items := strings.Split(cfiles, "/")
		if len(items) >= 2 {
			items[1] = arch
		}
		cfiles = strings.Join(items, "/")
	}

	return fmt.Sprintf("  %s %s", cmd, cfiles)
}

func genObject(arch string) string {
	rareCmds := []string{"HOSTCC ", "AS     "}
	cmds := []string{"AR     ", "CC     "}

	cmd := cmds[rand.Intn(len(cmds))]
	if GenBool(1.0 / 15.0) {
		cmd = rareCmds[rand.Intn(len(rareCmds))]
	}

	cfiles := fmt.Sprintf("%sh", fixtures.Cfiles[rand.Intn(len(fixtures.Cfiles))])
	if strings.HasPrefix(cfiles, "arch") {
		items := strings.Split(cfiles, "/")
		if len(items) >= 2 {
			items[1] = arch
		}
		cfiles = strings.Join(items, "/")
	}

	return fmt.Sprintf("  %s %s", cmd, cfiles)
}

func genSpecial(arch string) string {
	items := []string{
		"HOSTLD  arch/ARCH/tools/relocs",
		"HOSTLD  scripts/mod/modpost",
		"MKELF   scripts/mod/elfconfig.h",
		"LDS     arch/ARCH/entry/vdso/vdso32/vdso32.lds",
		"LDS     arch/ARCH/kernel/vmlinux.lds",
		"LDS     arch/ARCH/realmode/rm/realmode.lds",
		"LDS     arch/ARCH/boot/compressed/vmlinux.lds",
		"EXPORTS arch/ARCH/lib/lib-ksyms.o",
		"EXPORTS lib/lib-ksyms.o",
		"MODPOST vmlinux.o",
		"SORTEX  vmlinux",
		"SYSMAP  System.map",
		"VOFFSET arch/ARCH/boot/compressed/../voffset.h",
		"OBJCOPY arch/ARCH/entry/vdso/vdso32.so",
		"OBJCOPY arch/ARCH/realmode/rm/realmode.bin",
		"OBJCOPY arch/ARCH/boot/compressed/vmlinux.bin",
		"OBJCOPY arch/ARCH/boot/vmlinux.bin",
		"VDSO2C  arch/ARCH/entry/vdso/vdso-image-32.c",
		"VDSO    arch/ARCH/entry/vdso/vdso32.so.dbg",
		"RELOCS  arch/ARCH/realmode/rm/realmode.relocs",
		"PASYMS  arch/ARCH/realmode/rm/pasyms.h",
		"XZKERN  arch/ARCH/boot/compressed/vmlinux.bin.xz",
		"MKPIGGY arch/ARCH/boot/compressed/piggy.S",
		"DATAREL arch/ARCH/boot/compressed/vmlinux",
		"ZOFFSET arch/ARCH/boot/zoffset.h",
	}

	special := items[rand.Intn(len(items))]
	special = strings.ReplaceAll(special, "ARCH", arch)

	return fmt.Sprintf("  %s", special)
}

func GenKernelCompileLine(arch string) string {
	switch {
	case GenBool(1.0 / 15.9):
		return genSpecial(arch)
	case GenBool(0.1):
		return genHeader(arch)
	default:
		return genObject(arch)
	}
}
