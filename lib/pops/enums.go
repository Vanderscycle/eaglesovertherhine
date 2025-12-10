package pops

import "encoding/json"

type CivilianRole int

const (
	Farmer CivilianRole = iota
	Miner
	Physician
	Researcher
	FactoryWorker
	HouseKeeper
	Sailor
	Clergyman
	Baker
	MilitarySupport
)

var AssignedCivilianRole = map[CivilianRole]string{
	Farmer:          "Farmer",
	Miner:           "Miner",
	Physician:       "Physician",
	Researcher:      "Researcher",
	FactoryWorker:   "Factory Worker",
	HouseKeeper:     "Housekeeper",
	Sailor:          "Sailor",
	Clergyman:       "Clergyman",
	Baker:           "Baker",
	MilitarySupport: "Military Support",
}

func (c CivilianRole) String() string {
	return AssignedCivilianRole[c]
}

func (c CivilianRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

type MilitaryRole int

const (
	Soldier MilitaryRole = iota
	Nco
	Officier
	Logistics
	MusicColours
	Civilian
)

var AssignedRole = map[MilitaryRole]string{
	Soldier:      "Soldier",
	Nco:          "Non-Commisioned Officer",
	Officier:     "Officier",
	Logistics:    "Logistics",
	MusicColours: "Musician and Colours",
	Civilian:     "Civilian",
}

func (m MilitaryRole) String() string {
	return AssignedRole[m]
}

func (m MilitaryRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
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

func (r Rank) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
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

func (n Nation) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.String())
}
