package testData

type testData struct {
	args       []string
	pipeStdout string
	pipeStderr string
	stdout     string
	stderr     string
	stdin      string
	status     int
	timeout    float32
	shouldFail bool
	pre        string
	repeat     int
	post       string
	env        map[string]string
}

func (*testData) print(dt *testData) {

}
