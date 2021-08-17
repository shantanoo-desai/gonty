package joke

// Joke Struct that givess vital details
type Joke struct {
	Delivery string
	Deadly   bool
	Language string
}

// Jokes :Two of the Funniest (and one Deadly) Jokes in the World
var Jokes = []Joke{
	{
		Delivery: "Wenn ist das Nunstück git und Slotermeyer? Ja! Beiherhund das Oder die Flipperwaldt gersput!",
		Deadly:   true,
		Language: "German",
	},
	{
		Delivery: "There were zwei [two] peanuts walking down der strasse [street]. Und one was assaulted... peanut!",
		Deadly:   false,
		Language: "English",
	},
}
