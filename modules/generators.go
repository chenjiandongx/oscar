package modules

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"math/rand"
	"path"
	"time"

	"github.com/chenjiandongx/react/fixtures"
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
	idx := rand.Intn(len(fixtures.DockerTags))
	return fmt.Sprintf("%s", fixtures.DockerTags[idx])
}

func GenHashHex() string {
	key := [256]byte{}
	rand.Read(key[:])
	return fmt.Sprintf("%x", md5.Sum(key[:]))
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
