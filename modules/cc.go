package modules

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/chenjiandongx/oscar/fixtures"
)

type cc struct{}

func (mod *cc) Name() string {
	return ModCc
}

func (mod *cc) Display(highCpu bool) {
	if highCpu {
		ctx, cancel := context.WithCancel(context.Background())
		go BusyCPUWorking(ctx)
		defer func() {
			cancel()
			time.Sleep(100 * time.Microsecond)
		}()
	}

	var (
		compilers     = []string{"gcc", "clang"}
		flagsOpt      = []string{"-O0", "-O1", "-O2", "-O3", "-Og", "-Os"}
		flagsWarnBase = []string{"-Wall", "-Wall -Wextra"}
		flagsWarn     = []string{
			"-Wno-unused-variable",
			"-Wno-sign-compare",
			"-Wno-unknown-pragmas",
			"-Wno-parentheses",
			"-Wundef",
			"-Wwrite-strings",
			"-Wold-style-definition",
		}
		flagsF       = []string{"-fsigned-char", "-funroll-loops", "-fgnu89-inline", "-fPIC"}
		flagsArch    = []string{"-march=x86-64", "-mtune=generic", "-pipe"}
		flagsDefBase = []string{"-DDEBUG", "-DNDEBUG"}
		flagsDef     = []string{
			"-D_REENTRANT",
			"-DMATH_LOOP",
			"-D_LIBS_REENTRANT",
			"-DNAMESPACE=lib",
			"-DMODULE_NAME=lib",
			"-DPIC",
			"-DSHARED",
		}
	)

	pkg := fixtures.Packages[GenIntN(1, len(fixtures.Packages))]
	cmp := compilers[rand.Intn(len(compilers))]
	var cfiles []string
	for i := 0; i < GenIntN(300, 1200); i++ {
		cfiles = append(cfiles, fixtures.Cfiles[GenIntN(1, len(fixtures.Cfiles))])
	}
	sort.Strings(cfiles)
	opt := flagsOpt[rand.Intn(len(flagsOpt))]
	warn := flagsWarnBase[rand.Intn(len(flagsWarnBase))] + GenRandStringFromList(flagsWarn, GenIntN(1, len(flagsWarn)))

	fgs := GenRandStringFromList(flagsF, GenIntN(1, len(flagsF)))
	arch := GenRandStringFromList(flagsArch, GenIntN(1, len(flagsArch)))
	linkerFlags := GenLinkerFlags(fixtures.Packages, rand.Intn(10))

	includes := GenIncludes(cfiles, 20)

	defs := flagsDefBase[rand.Intn(len(flagsDefBase))] + GenRandStringFromList(flagsDef, GenIntN(1, len(flagsDef)))

	for _, cf := range cfiles {
		fmt.Printf(
			"%s -c %s %s%s%s %s%s -o %s\n",
			cmp, opt, warn, fgs, arch, includes, defs, strings.ReplaceAll(cf, ".c", ".o"),
		)
		SleepInMills(30, 200)
	}

	var buf bytes.Buffer
	for _, cf := range cfiles {
		buf.WriteString(strings.ReplaceAll(cf, ".c", ".o"))
		buf.WriteString(" ")
	}

	fmt.Printf("%s -o %s %s%s\n", cmp, pkg, buf.String(), linkerFlags)
	SleepInMills(300, 1000)
	PrintNewline()
}

func NewCcModule() Moduler {
	return &cc{}
}
