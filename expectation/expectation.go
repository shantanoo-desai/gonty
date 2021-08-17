package expectation

type expectation struct {
	Who       string
	Expecting bool
	Weapons   []string
}

// You can Never Expect them! Never!
// *Intense Music*
var NeverExpected = expectation{
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
