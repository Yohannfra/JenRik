package testData

import (
	"fmt"
	"github.com/Yohannfra/JenRik/internal/tomlLoader/tomlUtils"
	"github.com/Yohannfra/JenRik/internal/utils"
	"github.com/pelletier/go-toml"
)

type Test struct {
	name       string
	args       []string
	pipeStdout string
	pipeStderr string
	stdout     string
	stderr     string
	stdin      string
	status     int
	timeout    int
	shouldFail bool
	pre        string
	repeat     int
	post       string
	env        map[string]string
}

func (t Test) Print() {
	fmt.Printf("Test: '%s':\n", t.name)
	fmt.Print("\targs: ")
	utils.PrintStrArray(t.args)
	fmt.Printf("\tpipe stdout: '%s'\n", t.pipeStdout)
	fmt.Printf("\tpipe stderr: '%s'\n", t.pipeStderr)
	fmt.Printf("\tstdout: '%s'\n", t.stdout)
	fmt.Printf("\tstderr: '%s'\n", t.stderr)
	fmt.Printf("\tstdin: '%s'\n", t.stdin)
	fmt.Println("\tstatus: ", t.status)
	fmt.Println("\ttimeout: ", t.timeout)
	fmt.Println("\tshouldFail: ", t.shouldFail)
	fmt.Printf("\tpre: '%s'\n", t.pre)
	fmt.Println("\trepeat: ", t.repeat)
	fmt.Printf("\tpost: '%s\n", t.post)
	fmt.Println("\tenv: ", t.env)
}

func NewTest(name string, tomlContent *toml.Tree) *Test {
	t := new(Test)

	t.name = name

	t.status = int(tomlContent.Get("status").(int64))
	t.args = tomlUtils.ToStrArr(tomlContent.Get("args"))

	if tomlContent.Has("repeat") {
		t.repeat = int(tomlContent.Get("repeat").(int64))
	}
	if tomlContent.Has("timeout") {
		t.timeout = int(tomlContent.Get("timeout").(int64))
	}
	if tomlContent.Has("shouldFail") {
		t.shouldFail = tomlContent.Get("shouldFail").(bool)
	}
	if tomlContent.Has("stderr") {
		t.stderr = utils.ArrOrStrToStr(tomlContent.Get("stderr"))
	}
	if tomlContent.Has("stdout") {
		t.stdout = utils.ArrOrStrToStr(tomlContent.Get("stdout"))
	}
	if tomlContent.Has("pre") {
		t.pre = utils.ArrOrStrToStr(tomlContent.Get("pre"))
	}
	if tomlContent.Has("post") {
		t.post = utils.ArrOrStrToStr(tomlContent.Get("post"))
	}
	if tomlContent.Has("stdin") {
		t.stdin = utils.ArrOrStrToStr(tomlContent.Get("stdin"))
	}
	if tomlContent.Has("pipeStdout") {
		t.pipeStdout = tomlContent.Get("pipeStdout").(string)
	}
	if tomlContent.Has("pipeStderr") {
		t.pipeStderr = tomlContent.Get("pipeStderr").(string)
	}
	if tomlContent.Has("env") {
		t.env = tomlUtils.ToStrMap(tomlContent.Get("env").(*toml.Tree))
	}
	return t
}
