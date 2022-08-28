package questions

var (
	T_questionOneOne = Test{
		Name:       "Hello World",
		Executable: "../build/questionOneOne",
		TestSet: []testSet{
			{
				ItName: "Hello world executable",
				Result: "Hello World !!",
			},
		},
	}

	T_questionTwoTwo = Test{
		Name:       "Max3 Argv",
		Executable: "../build/questionTwoTwo",
		TestSet: []testSet{
			{
				ItName: "Max3 argv normal values",
				Args:   []string{"1", "10", "11"},
				Result: "11",
			},
			{
				ItName: "Max3 argv equal values",
				Args:   []string{"10", "10", "10"},
				Result: "10",
			},
			{
				ItName: "Max3 argv negative values",
				Args:   []string{"-1", "10", "2"},
				Result: "10",
			},
			{
				ItName: "Max3 argv normal values",
				Args:   []string{"10", "100", "30"},
				Result: "100",
			},
			{
				ItName: "Max3 argv negative values",
				Args:   []string{"-1", "-30", "-2"},
				Result: "-1",
			},
			{
				ItName: "Max3 argv only two values",
				Args:   []string{"-1", "-30"},
				Result: "",
			},
			{
				ItName: "Max3 argv four values",
				Args:   []string{"-1", "-30", "-2", "4"},
				Result: "",
			},
		},
	}

	T_questionTwoThree = Test{
		Name:       "Max3 stdin",
		Executable: "../build/questionTwoThree",
		TestSet: []testSet{
			{
				ItName:    "Max3 stdin normal values",
				StdinArgs: []string{"10", "12", "15"},
				Result:    "15",
			},
			{
				ItName:    "Max3 stdin negative values",
				StdinArgs: []string{"-1", "12", "0"},
				Result:    "12",
			},
			{
				ItName:    "Max3 stdin negative values",
				StdinArgs: []string{"-10", "-1", "0"},
				Result:    "0",
			},
			{
				ItName:    "Max3 stdin only zero",
				StdinArgs: []string{"0", "0", "0"},
				Result:    "0",
			},
			{
				ItName:    "Max3 stdin negative values",
				StdinArgs: []string{"-90", "-100", "-90"},
				Result:    "-90",
			},
		},
	}

	T_questionThreeTwo = Test{
		Name:       "Dayofbirth argv",
		Executable: "../build/questionThreeTwo",
		TestSet: []testSet{
			{
				ItName: "Dayofbirth argv normal inputs",
				Args:   []string{"Thibault", "C.", "23/03/2001"},
				Result: "Thibault C. was born on a Friday.",
			},
			{
				ItName: "Dayofbirth argv normal inputs",
				Args:   []string{"Gerald", "O.", "19/03/1978"},
				Result: "Gerald O. was born on a Sunday.",
			},
			{
				ItName: "Dayofbirth argv normal inputs",
				Args:   []string{"John", "D.", "23/12/1970"},
				Result: "John D. was born on a Wednesday.",
			},
			{
				ItName: "Dayofbirth argv unexisting date",
				Args:   []string{"John", "C.", "30/14/1978"},
				Result: "The date 30/14/1978 don't exist.",
			},
			{
				ItName: "Dayofbirth argv not enough arguments",
				Args:   []string{"Thibault", "C."},
				Result: "",
			},
		},
	}

	T_questionThreeThree = Test{
		Name:       "Dayofbirth stdin",
		Executable: "../build/questionThreeThree",
		TestSet: []testSet{
			{
				ItName:    "Dayofbirth argv normal inputs",
				StdinArgs: []string{"Thibault C.", "23/03/2001"},
				Result:    "Thibault C. was born on a Friday.",
			},
			{
				ItName:    "Dayofbirth argv normal inputs",
				StdinArgs: []string{"Gerald O.", "19/03/1978"},
				Result:    "Gerald O. was born on a Sunday.",
			},
			{
				ItName:    "Dayofbirth argv normal inputs",
				StdinArgs: []string{"John D.", "23/12/1970"},
				Result:    "John D. was born on a Wednesday.",
			},
			{
				ItName:    "Dayofbirth argv unexisting date",
				StdinArgs: []string{"John C.", "30/14/1978"},
				Result:    "The date 30/14/1978 don't exist.",
			},
		},
	}
)
