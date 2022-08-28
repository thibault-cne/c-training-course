package questions

type Test struct {
	Name       string
	Executable string
	TestSet    []testSet
}

type testSet struct {
	ItName    string
	Args      []string
	Result    string
	StdinArgs []string
}
