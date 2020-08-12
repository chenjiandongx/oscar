package modules

import (
	"fmt"
	"reflect"
)

type git struct{}

type gitConsoleLine struct {
	c0 string
	c1 string
	c2 string
	c3 string
	c4 string
	c5 string
	c6 string
	c7 string
	c8 string
}

func (mod *git) Name() string {
	return ModGit
}

func (mod *git) Display() {

	gcl := gitConsoleLine{
		"cd ~/git",
		"git clone https://github.com/chenjiandongx/oscar.git",
		"Cloning into 'oscar'...",
		"remote: Enumerating objects: 101, done.",
		"remote: Counting objects: 100% (101/101), done.",
		"remote: Compressing objects: 100% (71/71), done.",
		"remote: Total 101 (delta 53), reused 77 (delta 29), pack-reused 0",
		"Receiving objects: 100% (101/101), 244.37 KiB | 23.00 KiB/s, done.",
		"Resolving deltas: 100% (53/53), done.",
	}

	v := reflect.ValueOf(gcl)
	for i := 0; i < 9; i++ {
		SleepInMills(500, 1000)
		fmt.Println(v.Field(i))
	}
}

func NewGitModule() Moduler {
	return &git{}
}
