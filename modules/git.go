package modules

import (
	"fmt"
	"github.com/chenjiandongx/oscar/fixtures"
)

type git struct{}

func (mod *git) Name() string {
	return ModGit
}

func (mod *git) Display() {

	for i := 0; i < len(fixtures.GitConsole); i++ {
		SleepInMills(500, 1000)
		fmt.Println(fixtures.GitConsole[i])
	}

}

func NewGitModule() Moduler {
	return &git{}
}
