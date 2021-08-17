package nobleman

type Person struct {
	Name       string
	FamilyName string
	Empire     string
	RelatedTo  []Person
}

var Biggus = Person{
	Name:       "Biggus",
	FamilyName: "Dickus",
	Empire:     "Roman",
	RelatedTo: []Person{
		{
			Name:       "Pontius",
			FamilyName: "Pilate",
			Empire:     "Roman",
			RelatedTo: []Person{
				{
					Name:       "Julius",
					FamilyName: "Caesar",
					Empire:     "Roman",
				},
			},
		},
		{
			Name:       "Incontinetia",
			FamilyName: "Buttocks",
			Empire:     "Roman",
		},
	},
}
