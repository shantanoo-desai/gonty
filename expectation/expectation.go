package expectation

type Expectation struct {
	Who       string
	Expecting bool
	Weapons   []string
}

var NeverExpected = Expectation{
	Who:       "Spanish Inquisition",
	Expecting: false,
	Weapons: []string{
		"fear",
		"surprise",
		"ruthless efficiency",
		"fanatical devotion to the pope",
		"nice red uniforms",
	},
}
