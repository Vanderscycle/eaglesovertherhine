package pops

type Role int

const (
	Soldier Role = iota
	Nco
	Officier
	Logistics
	MusicColours
	Civilian
)

var AssignedRole = map[Role]string{
	Soldier:      "Soldier",
	Nco:          "Non-Commisioned Officer",
	Officier:     "Officier",
	Logistics:    "Logistics",
	MusicColours: "Musician and Colours",
	Civilian:     "Civilian",
}

func (r Role) String() string {
	return AssignedRole[r]
}

type Sex int

const (
	Male Sex = iota
	Female
)

var PopsSex = map[Sex]string{
	Male:   "Male",
	Female: "Female",
}

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
	ChefDeBattalion
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
	ChefDeBattalion:      "Chef de Battaillon",
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
