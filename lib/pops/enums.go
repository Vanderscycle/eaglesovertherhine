package pops

type Rank int

const (
	Soldat Rank = iota
	Caporal
	CaporalFourrier
	Sergent
	SergentMajor
	AdjudantSousOfficier
	AdjudantMajor
	SousLieutenant
	Lieutenant
	Capitaine
	ChefDeBattaillon
	MajorEnSecond
	Major
)

var AssignedRank = map[Rank]string{
	Soldat:               "Soldat",
	Caporal:              "Caporal",
	CaporalFourrier:      "Caporal Fourrier",
	Sergent:              "Sergent",
	SergentMajor:         "Sergent Major",
	AdjudantSousOfficier: "Adjudant Sous-Officier",
	AdjudantMajor:        "Adjudant Major",
	SousLieutenant:       "Sous Lieutenant",
	Lieutenant:           "Lieutenant",
	Capitaine:            "Capitaine",
	ChefDeBattaillon:     "Chef de Battaillon",
	MajorEnSecond:        "Major en Second",
	Major:                "Major",
}

func (r Rank) String() string {
    return AssignedRank[r]
}

type Nation int

const (
	France Nation = iota
	Austria
	Brittain
)

var AssignedNation = map[Nation]string{
	France:   "France",
	Austria:  "Austria",
	Brittain: "Brittain",
}
func (n Nation) String() string {
    return AssignedNation[n]
}
