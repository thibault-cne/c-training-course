package main

type Test struct {
	Name    string
	TestSet []TestSet
}

type TestSet struct {
	ItName    string
	Args      []string
	Result    string
	StdinArgs []string
}

var (
	endTests = []Test{
		{
			Name: "Testing EXIT functions",
			TestSet: []TestSet{
				{
					ItName:    "Testing normal EXIT input",
					StdinArgs: []string{"EXIT"},
					Result:    "exit status 1",
				},
				{
					ItName:    "Testing EXIT after ADD",
					StdinArgs: []string{"ADD", "Thibault", "Cheneviere", "0782423595", "Rien", "Thib", "EXIT"},
					Result:    "exit status 1",
				},
				{
					ItName: "Testing EXIT after ADD and SEARCH",
					StdinArgs: []string{"ADD", "Thibault", "Cheneviere", "0782423595", "Rien", "Thib",
						"SEARCH", "0", "EXIT"},
					Result: "exit status 1",
				},
				{
					ItName:    "Testing EXIT after SEARCH",
					StdinArgs: []string{"SEARCH", "0", "EXIT"},
					Result:    "exit status 1",
				},
			},
		},
	}
)
