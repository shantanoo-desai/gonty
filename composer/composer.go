package composer

import "strings"

var composer_surname = "de von Ausfern-schplenden-schlitter-crasscrenbon-fried-digger-dingle-dangle-dongle-dungle-burstein-von-knacker-thrasher-apple-banger-horowitz-ticolensic-grander-knotty-spelltinkle-grandlich-grumblemeyer-spelterwasser-kurstlich-himbleeisen-bahnwagen-gutenabend-bitte-ein-nürnburger-bratwustle-gerspurten-mitzweimache-luber-hundsfut-gumberaber-shönendanker-kalbsfleisch-mittler-aucher von Hautkopft of Ulm"

type Composer struct {
	Name        string
	FamilyName  []string
	Genre       string
	Nationality string
}

var GreatestBaroqueComposer = Composer{
	Name:        "Johann Gambolputty",
	FamilyName:  strings.Split(composer_surname, "-"),
	Genre:       "Baroque",
	Nationality: "German",
}
