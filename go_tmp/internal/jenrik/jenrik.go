package jenrik

import (
	"fmt"
	"log"
	"github.com/pelletier/go-toml"
	"github.com/Yohannfra/JenRik/internal/utils"
	"github.com/Yohannfra/JenRik/internal/parser"
)

func run_build_command(command string) {
	fmt.Println(command)
}


func start_jenrik(fp string) {
	fc := get_file_content(fp)
	toml_content, err := toml.Load(fc)


	binary_path := ""
	// var test_dict map[string]interface{} // string -> *tomlValue, *Tree, []*Tree

	if err != nil {
		log.Fatal(err)
	}
	for _, key := range toml_content.Keys() {
		if key == "binary_path" {
			binary_path = toml_content.Get(key).(string)
		} else if key == "build_command" {
			run_build_command(toml_content.Get(key).(string))
		} else {
			check_tests_validity(key, toml_content.Get(key).(*toml.Tree))
			// test_dict[key] = toml_content.Get(key).(string)
			fmt.Println(toml_content.Get(key).(*toml.Tree))
		}
	}
	if binary_path == "" {
		log.Fatal("Could not find binary_path key in", fp)
	}

}

