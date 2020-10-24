package testData

import (
	"fmt"
	"github.com/Yohannfra/JenRik/internal/tomlLoader/tomlUtils"
	"github.com/Yohannfra/JenRik/internal/utils"
	"github.com/pelletier/go-toml"
)

type Test struct {
	Name       string
	Args       []string
	PipeStdout string
	PipeStderr string
	Stdout     string
	Stderr     string
	Stdin      string
	Status     int
	Timeout    int
	ShouldFail bool
	Pre        string
	Repeat     int
	Post       string
	Env        map[string]string
}

func (t Test) Print() {
	fmt.Printf("Test: '%s':\n", t.Name)
	fmt.Print("\targs: ")
	utils.PrintStrArray(t.Args)
	fmt.Printf("\tpipe stdout: '%s'\n", t.PipeStdout)
	fmt.Printf("\tpipe stderr: '%s'\n", t.PipeStderr)
	fmt.Printf("\tstdout: '%s'\n", t.Stdout)
	fmt.Printf("\tstderr: '%s'\n", t.Stderr)
	fmt.Printf("\tstdin: '%s'\n", t.Stdin)
	fmt.Println("\tstatus: ", t.Status)
	fmt.Println("\ttimeout: ", t.Timeout)
	fmt.Println("\tshouldFail: ", t.ShouldFail)
	fmt.Printf("\tpre: '%s'\n", t.Pre)
	fmt.Println("\trepeat: ", t.Repeat)
	fmt.Printf("\tpost: '%s\n", t.Post)
	fmt.Println("\tenv: ", t.Env)
}

func NewTest(name string, tomlContent *toml.Tree) *Test {
	t := new(Test)

	t.Name = name

	t.Status = int(tomlContent.Get("status").(int64))
	t.Args = tomlUtils.ToStrArr(tomlContent.Get("args"))

	if tomlContent.Has("repeat") {
		t.Repeat = int(tomlContent.Get("repeat").(int64))
	}
	if tomlContent.Has("timeout") {
		t.Timeout = int(tomlContent.Get("timeout").(int64))
	}
	if tomlContent.Has("shouldFail") {
		t.ShouldFail = tomlContent.Get("shouldFail").(bool)
	}
	if tomlContent.Has("stderr") {
		t.Stderr = utils.ArrOrStrToStr(tomlContent.Get("stderr"))
	}
	if tomlContent.Has("stdout") {
		t.Stdout = utils.ArrOrStrToStr(tomlContent.Get("stdout"))
	}
	if tomlContent.Has("pre") {
		t.Pre = utils.ArrOrStrToStr(tomlContent.Get("pre"))
	}
	if tomlContent.Has("post") {
		t.Post = utils.ArrOrStrToStr(tomlContent.Get("post"))
	}
	if tomlContent.Has("stdin") {
		t.Stdin = utils.ArrOrStrToStr(tomlContent.Get("stdin"))
	}
	if tomlContent.Has("pipeStdout") {
		t.PipeStdout = tomlContent.Get("pipeStdout").(string)
	}
	if tomlContent.Has("pipeStderr") {
		t.PipeStderr = tomlContent.Get("pipeStderr").(string)
	}
	if tomlContent.Has("env") {
		t.Env = tomlUtils.ToStrMap(tomlContent.Get("env").(*toml.Tree))
	}
	return t
}
