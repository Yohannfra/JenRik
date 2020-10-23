package jenrik

import (
	"fmt"
	"github.com/Yohannfra/JenRik/internal/tester"
	"github.com/Yohannfra/JenRik/internal/tomlLoader"
	"github.com/Yohannfra/JenRik/internal/utils"
	"log"
	"os"
)

// create a default test file
func Init(fp string) {
	testFileName := "test_" + fp + ".toml"
	defaultFileContent :=
		"binary_path = \"{%s}\"\n\n" +
			"# A sample test\n" +
			"[test1]\n" +
			"args = [\"-h\"]\n" +
			"status = 0\n" +
			"stdout=\"\"\n" +
			"stderr=\"\"\n"

	if utils.FileExists(testFileName) {
		log.Fatalf("%s: File already exists\n", testFileName)
	}

	f, err := os.Create(testFileName)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.WriteString(defaultFileContent)
	if err != nil {
		f.Close()
		log.Fatal(err)
	}
	fmt.Printf("Initialized %s with success\n", testFileName)
}

func Start(fp string) {
	tomlContent := tomlLoader.LoadTestFile(fp)
	tester.Run(&tomlContent)
}
