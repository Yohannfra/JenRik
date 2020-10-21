package jenrik

import (
	"fmt"
	"github.com/Yohannfra/JenRik/internal/parser"
	"github.com/Yohannfra/JenRik/internal/utils"
	"github.com/pelletier/go-toml"
	"log"
)

func runBuildCommand(command string) {
	fmt.Println(command)
}

func Start(fp string) {
	fc := utils.GetFileContent(fp)
	tomlContent, err := toml.Load(fc)

	binaryPath := ""
	// var test_dict map[string]interface{} // string -> *tomlValue, *Tree, []*Tree

	if err != nil {
		log.Fatal(err)
	}
	for _, key := range tomlContent.Keys() {
		if key == "binary_path" {
			binaryPath = tomlContent.Get(key).(string)
		} else if key == "build_command" {
			runBuildCommand(tomlContent.Get(key).(string))
		} else {
			parser.CheckTestsValidity(key, tomlContent.Get(key).(*toml.Tree))
			// test_dict[key] = toml_content.Get(key).(string)
			fmt.Println(tomlContent.Get(key).(*toml.Tree))
		}
	}
	if binaryPath == "" {
		log.Fatal("Could not find binary_path key in", fp)
	}

}
